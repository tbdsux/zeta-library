package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
)

var (
	prod = flag.Bool("prod", false, "Enable prefork in Production")
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork: *prod,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":8080")
}
