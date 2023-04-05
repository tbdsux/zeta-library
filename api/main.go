package main

import (
	"flag"
	"fmt"
	"os"

	col "github.com/TheBoringDude/zeta-library/api/routes/collection"
	"github.com/TheBoringDude/zeta-library/api/routes/items"
	"github.com/TheBoringDude/zeta-library/api/routes/settings"
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
	port := getPort()

	// use cors middleware
	app.Use(cors.New())

	collection := app.Group("/collections")
	_items := app.Group("/items")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	// `/collection` route
	collection.Post("/create", col.CreateCollection)
	collection.Get("/get/:id", col.GetCollection)
	collection.Get("/", col.FetchAllCollections)
	collection.Patch("/:id", col.UpdateCollection)
	collection.Delete("/:id", col.DeleteCollection)

	// `/items` route
	_items.Patch("/", items.UpdateItems)
	_items.Get("/:id", items.FetchItems)
	_items.Delete("/:id/:itemId", items.RemoveItem)

	// `/settings` route
	app.Patch("/settings", settings.UpdateSettings)
	app.Get("/settings", settings.GetSettings)

	app.Listen(fmt.Sprintf(":%s", port))
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8080"
	}

	return port
}
