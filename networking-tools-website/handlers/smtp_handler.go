package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gokhangokcen1/subnet-backend/models"
	"github.com/gokhangokcen1/subnet-backend/smtpsender"
)

func SendSMTPMessageHandler(c fiber.Ctx) error {
	request := new(models.SMTPMessageRequest)
	if err := c.Bind().Body(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.SMTPMessageResponse{Message: "İstek verisi okunamadı."})
	}
	if request.Host == "" || request.Port == "" || request.From == "" || request.To == "" || request.Subject == "" || request.Body == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.SMTPMessageResponse{Message: "SMTP sunucusu, port, gönderen, alıcı, konu ve mesaj alanları zorunludur."})
	}
	result, err := smtpsender.Send(request.Host, request.Port, request.From, request.To, request.Subject, request.Body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.SMTPMessageResponse{Message: err.Error(), Trace: result.Trace})
	}
	return c.JSON(models.SMTPMessageResponse{Success: true, Message: "E-posta SMTP sunucusu tarafından teslim için kabul edildi.", Trace: result.Trace})
}
