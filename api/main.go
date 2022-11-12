package main

import (
	"flag"

	col "github.com/TheBoringDude/zeta-library/api/routes/collection"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var (
	prod = flag.Bool("prod", false, "Enable prefork in Production")
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork: *prod,
	})

	// use cors middleware
	app.Use(cors.New())

	collection := app.Group("/collections")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	// `/collection` route
	collection.Post("/create", col.CreateCollection)
	collection.Get("/get/:id", col.GetCollection)
	collection.Get("/", col.FetchAllCollections)

	app.Listen(":8080")
}
