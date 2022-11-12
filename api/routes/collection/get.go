package col

import (
	"github.com/TheBoringDude/zeta-library/api/lib"
	"github.com/deta/deta-go/deta"
	"github.com/gofiber/fiber/v2"
)

var GetCollection = func(c *fiber.Ctx) error {
	id := c.Params("id")

	db, err := lib.InitBase("_all_collections")
	if err != nil {
		return c.Status(500).JSON(lib.APIResponse{
			Error:   true,
			Message: err.Error(),
		})
	}

	collection := &CollectionProps{}
	err = db.Get(id, collection)

	// if not found
	if err == deta.ErrNotFound {
		return c.Status(404).JSON(lib.APIResponse{
			Error:   true,
			Message: "Collection not found.",
		})
	}

	// other errors
	if err != nil {
		return c.Status(500).JSON(lib.APIResponse{
			Error:   true,
			Message: err.Error(),
		})
	}

	return c.Status(200).JSON(lib.APIResponse{
		Error: false,
		Data:  collection,
	})
}
