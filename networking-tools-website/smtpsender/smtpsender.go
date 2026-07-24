package smtpsender

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/textproto"
	"os"
	"strings"
	"time"
)

type Result struct {
	Trace []string
}

// Send performs the SMTP dialogue directly over TCP. SMTP commands, CRLF line
// endings, DATA dot-stuffing, and response validation are intentionally handled
// here instead of delegating them to net/smtp.
func Send(host, port, from, to, subject, body string) (Result, error) {
	result := Result{}
	step := func(message string) { result.Trace = append(result.Trace, message) }
	host, from, to = strings.TrimSpace(host), strings.TrimSpace(from), strings.TrimSpace(to)
	if host == "" {
		return result, errors.New("SMTP sunucusu zorunludur")
	}
	step(fmt.Sprintf("[1] Hedef SMTP sunucusu: %s:%s", host, port))
	if port != "25" && port != "587" {
		return result, errors.New("yalnızca 25 veya 587 numaralı port kullanılabilir")
	}
	if net.ParseIP(host) != nil {
		step("[2] DNS: Hedef bir IP adresi olduğu için DNS çözümlemesi gerekmedi.")
	} else if addresses, err := net.LookupHost(host); err != nil {
		step(fmt.Sprintf("[2] DNS başarısız: %v", err))
		return result, fmt.Errorf("SMTP sunucusunun DNS kaydı çözümlenemedi: %w", err)
	} else {
		step(fmt.Sprintf("[2] DNS başarılı: %s → %s", host, strings.Join(addresses, ", ")))
	}
	if parsed, err := mail.ParseAddress(from); err != nil || parsed.Address != from {
		return result, errors.New("geçerli bir gönderen e-posta adresi girin")
	}
	if parsed, err := mail.ParseAddress(to); err != nil || parsed.Address != to {
		return result, errors.New("geçerli bir alıcı e-posta adresi girin")
	}
	if len(subject) > 200 || len(body) > 10000 {
		return result, errors.New("konu veya mesaj izin verilen uzunluğu aşıyor")
	}

	step("[3] TCP bağlantısı kuruluyor...")
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), 20*time.Second)
	if err != nil {
		step(fmt.Sprintf("[3] TCP bağlantısı başarısız: %v", err))
		return result, fmt.Errorf("SMTP relay sunucusuna bağlanılamadı: %w", err)
	}
	defer conn.Close()
	step(fmt.Sprintf("[3] TCP bağlantısı başarılı: yerel %s → uzak %s", conn.LocalAddr(), conn.RemoteAddr()))

	reader := textproto.NewReader(bufio.NewReader(conn))
	writer := bufio.NewWriter(conn)
	readResponse := func(expectedCode int, label string) error {
		code, message, err := reader.ReadResponse(0)
		if err != nil {
			step(fmt.Sprintf("%s Sunucu yanıtı okunamadı: %v", label, err))
			return err
		}
		step(fmt.Sprintf("%s Sunucu yanıtı: %d %s", label, code, message))
		if code != expectedCode {
			return fmt.Errorf("%s beklenmeyen SMTP yanıtı: %d %s", label, code, message)
		}
		return nil
	}
	sendCommand := func(command string, expectedCode int, label string) error {
		step(fmt.Sprintf("%s İstemci komutu: %s", label, command))
		if _, err := writer.WriteString(command + "\r\n"); err != nil {
			return err
		}
		if err := writer.Flush(); err != nil {
			return err
		}
		return readResponse(expectedCode, label)
	}

	step("[4] SMTP sunucusunun 220 karşılama mesajı okunuyor...")
	if err := readResponse(220, "[4]"); err != nil {
		return result, err
	}

	helo := heloName()
	if err := sendCommand("EHLO "+helo, 250, "[5]"); err != nil {
		step("[5] EHLO reddedildi; HELO ile yeniden deneniyor...")
		if err := sendCommand("HELO "+helo, 250, "[5]"); err != nil {
			return result, err
		}
	}
	step("[6] STARTTLS/TLS atlandı; SMTP trafiği şifrelenmeden devam ediyor.")
	if err := sendCommand("MAIL FROM:<"+from+">", 250, "[7]"); err != nil {
		return result, err
	}
	if err := sendCommand("RCPT TO:<"+to+">", 250, "[8]"); err != nil {
		return result, err
	}
	if err := sendCommand("DATA", 354, "[9]"); err != nil {
		return result, err
	}

	message := dotStuff(buildMessage(from, to, subject, body))
	step(fmt.Sprintf("[9] DATA içeriği gönderiliyor: %d bayt; CRLF ve nokta-sonlandırması uygulandı.", len(message)))
	if _, err := writer.Write(message); err != nil {
		return result, err
	}
	if _, err := writer.WriteString(".\r\n"); err != nil {
		return result, err
	}
	if err := writer.Flush(); err != nil {
		return result, err
	}
	if err := readResponse(250, "[9]"); err != nil {
		return result, err
	}
	step("[9] Mesaj veri bölümü kabul edildi; sunucu teslim kuyruğuna aldı.")
	if err := sendCommand("QUIT", 221, "[10]"); err != nil {
		return result, err
	}
	step("[10] SMTP bağlantısı düzgün biçimde kapatıldı.")
	return result, nil
}

func heloName() string {
	name, _ := os.Hostname()
	if name == "" {
		return "localhost"
	}
	return name
}

func buildMessage(from, to, subject, body string) []byte {
	cleanHeader := strings.NewReplacer("\r", "", "\n", "")
	body = strings.ReplaceAll(strings.ReplaceAll(body, "\r\n", "\n"), "\n", "\r\n")
	return []byte(fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/plain; charset=UTF-8\r\nContent-Transfer-Encoding: 8bit\r\n\r\n%s\r\n", cleanHeader.Replace(from), cleanHeader.Replace(to), cleanHeader.Replace(subject), body))
}

// dotStuff escapes a leading dot on every DATA line, as required by SMTP.
func dotStuff(message []byte) []byte {
	lines := strings.Split(string(message), "\r\n")
	for i, line := range lines {
		if strings.HasPrefix(line, ".") {
			lines[i] = "." + line
		}
	}
	return []byte(strings.Join(lines, "\r\n"))
}
