package col

import (
	"fmt"

	"github.com/TheBoringDude/zeta-library/api/lib"
	"github.com/TheBoringDude/zeta-library/api/routes/items"
	"github.com/deta/deta-go/deta"
	"github.com/deta/deta-go/service/base"
	"github.com/gofiber/fiber/v2"
)

var DeleteCollection = func(c *fiber.Ctx) error {
	id := c.Params("id")

	db, err := lib.InitBase("_all_collections")
	if err != nil {
		return c.Status(500).JSON(lib.APIResponse{
			Error:   true,
			Message: err.Error(),
		})
	}

	var collection *lib.CollectionProps
	if err = db.Get(id, collection); err == deta.ErrNotFound {
		return c.Status(404).JSON(lib.APIResponse{
			Error:   true,
			Message: "Collection not found.",
		})
	} else if err != nil {
		return c.Status(500).JSON(lib.APIResponse{
			Error:   true,
			Message: err.Error(),
		})
	}

	if err = db.Delete(id); err != nil {
		return c.Status(500).JSON(lib.APIResponse{
			Error:   true,
			Message: err.Error(),
		})
	}

	// remove items in the collection base
	itemsDB, err := lib.InitBase(id)
	if err != nil {
		return c.Status(500).JSON(lib.APIResponse{
			Error:   true,
			Message: err.Error(),
		})
	}

	var items []*items.ItemsData
	if _, err = itemsDB.Fetch(&base.FetchInput{
		Q:    base.Query{},
		Dest: &items,
	}); err != nil {
		return c.Status(500).JSON(lib.APIResponse{
			Error:   true,
			Message: err.Error(),
		})
	}

	for _, v := range items {
		if err = itemsDB.Delete(v.Key); err != nil {
			return c.Status(500).JSON(lib.APIResponse{
				Error:   true,
				Message: fmt.Sprintf("Error removing item in collection's base: %s", err.Error()),
			})
		}
	}

	return c.JSON(lib.APIResponse{
		Error:   false,
		Message: "Successfully deleted collection.",
	})
}
