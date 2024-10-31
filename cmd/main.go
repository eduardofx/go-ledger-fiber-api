package main

import (
	"go-api/pkg/domain/accounts"
	"go-api/pkg/shared/config"
	"go-api/pkg/shared/db"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(&c)
	app := fiber.New()

	accounts.RegisterRoutes(app, h)

	app.Listen(c.Port)
}
