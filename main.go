package main

import (
	"log"

	"github.com/gofiber/fiber/v3" 
)

type User struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

func main() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error { 
		return c.SendString("Welcome to my Fiber v3 API! 'Belton'")
	})

	app.Get("/status", func(c fiber.Ctx) error { 
		return c.JSON(fiber.Map{
			"status": "online",
			"code":   200,
		})
	})

	app.Post("/user", func(c fiber.Ctx) error { 
		user := new(User)
		if err := c.Bind().JSON(user); err != nil { 
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		return c.JSON(fiber.Map{
			"message": "User created successfully",
			"user":    user,
		})
	})

	log.Fatal(app.Listen(":3000"))
}
