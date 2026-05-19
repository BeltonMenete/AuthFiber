package main

import (
	"log"

	"github.com/gofiber/contrib/v3/swaggerui"
	"github.com/gofiber/contrib/v3/swaggo"
	"github.com/gofiber/fiber/v3"
	scalar "github.com/yokeTH/gofiber-scalar/scalar/v3"
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
	app := fiber.New()

	const openApiSpec = `{
	  "openapi": "3.0.0",
	  "info": { "title": "Fiber v3 API Reference", "version": "1.0.0" },
	  "paths": {
	    "/users": {
	      "get": {
	        "summary": "Fetch all users",
	        "responses": { "200": { "description": "Success" } }
	      }
	    }
	  }
	}`

	app.Get("/swagger.json", func(c fiber.Ctx) error {
		c.Set("Content-Type", "application/json")
		return c.SendString(openApiSpec)
	})

	app.Get("/docs/*", scalar.New(scalar.Config{
		Title:             "Fiber v3 API Reference (Scalar)",
		Path:              "docs",
		FileContentString: openApiSpec,
	}))

	app.Use(swaggerui.New(swaggerui.Config{
		Title:    "Fiber v3 API Reference (SwaggerUI)",
		Path:     "swagger-ui",
		URL:      "/swagger.json",
		BasePath: "/",
	}))

	app.Get("/swagger/*", swaggo.HandlerDefault)

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Welcome to my Fiber v3 API! 'Belton'")
	})

	app.Get("/status", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "online",
			"code":   200,
		})
	})

	app.Get("/users", func(c fiber.Ctx) error {
		return c.JSON(users)
	})

	app.Post("/users", func(c fiber.Ctx) error {
		user := new(User)
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