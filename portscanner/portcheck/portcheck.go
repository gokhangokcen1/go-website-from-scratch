package portcheck

import (
	"fmt"
	"net"
	"time"
)

func KontrolEt(ip string, port int) (bool, string) {
	adres := fmt.Sprintf("%s:%d", ip, port)
	timeout := 3 * time.Second

	conn, err := net.DialTimeout("tcp", adres, timeout)
	if err != nil {
		return false, err.Error()
	}
	defer conn.Close()

	return true, ""
}
