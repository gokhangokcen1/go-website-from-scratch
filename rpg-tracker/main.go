package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gokhangokcen1/rpg-tracker/models"
)

func TumStatlariYazdir(statlar []models.Stat) string {
	sonuc := ""
	for idx, stat := range statlar {
		sonuc += fmt.Sprintf("%d. %s / Seviye: %d / XP: %d\n", idx+1, stat.Ad, stat.Seviye, stat.XP)
	}
	return sonuc
}

func main() {
	zekaStat := models.Stat{Ad: "Zeka", Seviye: 1, XP: 0}

	statlar := []models.Stat{
		zekaStat,
		{Ad: "Dayanıklılık", Seviye: 10, XP: 85},
		{Ad: "Sanat", Seviye: 3, XP: 40},
	}

	t := models.Task{Baslik: "Kitap Okuma", Tamamlandi: false, Stat: &zekaStat, XPDegeri: 10}
	fmt.Println(zekaStat)
	fmt.Println(t)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Merhaba dünya")
	})

	app.Get("/stats", func(c *fiber.Ctx) error {
		return c.SendString(TumStatlariYazdir(statlar))
	})

	app.Get("/tasks", func(c *fiber.Ctx) error {
		return c.SendString(t.Baslik)
	})

	app.Listen(":3000")
}
