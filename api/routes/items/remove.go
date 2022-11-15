package items

import (
	"github.com/TheBoringDude/zeta-library/api/lib"
	"github.com/deta/deta-go/deta"
	"github.com/gofiber/fiber/v2"
)

var RemoveItem = func(c *fiber.Ctx) error {
	id := c.Params("id")
	itemId := c.Params("itemId")

	db, err := lib.InitBase("_all_collections")
	if err != nil {
		return c.Status(500).JSON(lib.APIResponse{
			Error:   true,
			Message: err.Error(),
		})
	}

	_, err = lib.GetCollection(db, id)
	if err == deta.ErrNotFound {
		return c.Status(404).JSON(lib.APIResponse{
			Error:   true,
			Message: "Collection not found.",
		})
	}
	if err != nil {
		return c.Status(500).JSON(lib.APIResponse{
			Error:   true,
			Message: err.Error(),
		})
	}

	itemsDB, err := lib.InitBase(id)
	if err != nil {
		return c.Status(500).JSON(lib.APIResponse{
			Error:   true,
			Message: err.Error(),
		})
	}

	if err = itemsDB.Delete(itemId); err != nil {
		return c.Status(500).JSON(lib.APIResponse{
			Error:   true,
			Message: err.Error(),
		})
	}

	return c.JSON(lib.APIResponse{
		Error: false,
		Data: fiber.Map{
			"Id": id,
		},
		Message: "Successfully removed item from collection's list.",
	})
}
