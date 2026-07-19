package handlers

import (
	"github.com/gofiber/fiber/v3"

	"github.com/gokhangokcen1/subnet-backend/dnscheck"
	"github.com/gokhangokcen1/subnet-backend/models"
)

// CheckDNSHandler: body parse -> validasyon -> dnscheck.CheckAllRecords cagir -> yanit.
// Onceki versiyondan fark: artik kayit tipi secimi yok, TEK istekte
// TUM kayit tiplerini donduruyor (senin genisletilmis kodunun mantigi).
func CheckDNSHandler(c fiber.Ctx) error {
	req := new(models.DNSCheckRequest)
	if err := c.Bind().Body(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Istek govdesi okunamadi"})
	}
	if req.Domain == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Domain zorunludur"})
	}

	result := dnscheck.CheckAllRecords(req.Domain)
	return c.Status(fiber.StatusOK).JSON(result)
}
