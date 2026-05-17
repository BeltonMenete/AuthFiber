package main

import (
	"log"

	// Native v3 imports
	"github.com/gofiber/fiber/v3" 
	scalar "github.com/yokeTH/gofiber-scalar"
)

type User struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

var users = []User{
	{Name: "Belton", Role: "Admin"},
	{Name: "Alice", Role: "Developer"},
}

func main() {
	// Initialize Fiber v3 app instance
	app := fiber.New()

	// Use Scalar middleware 
	app.Use("/docs", scalar.New(scalar.Config{
		Title: "Fiber v3 API Reference",
	}))

	app.Get("/", func(c fiber.Ctx) error { 
		return c.SendString("Welcome to my Fiber v3 API! 'Belton'")
	})

	app.Get("/status", func(c fiber.Ctx) error { 
		return c.JSON(fiber.Map{
			"status": "online",
			"code":   200,
		})
	})

	// Fetch users GET endpoint
	app.Get("/users", func(c fiber.Ctx) error {
		return c.JSON(users)
	})

	app.Post("/user", func(c fiber.Ctx) error { 
		user := new(User)
		// Fiber v3 structural data binding pattern
		if err := c.Bind().JSON(user); err != nil { 
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		
		users = append(users, *user)

		return c.JSON(fiber.Map{
			"message": "User created successfully",
			"user":    user,
		})
	})

	log.Fatal(app.Listen(":3000"))
}
