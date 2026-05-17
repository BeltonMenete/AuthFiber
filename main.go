package main

import (
	"log"

	"github.com/gofiber/fiber/v3" // Must be v3
)

func main() {
	app := fiber.New()

	// In v3, use 'c fiber.Ctx' without the '*' pointer
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello from Fiber v3!")
	})

	log.Fatal(app.Listen(":3000"))
}
