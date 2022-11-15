package items

import (
	"github.com/TheBoringDude/zeta-library/api/lib"
	"github.com/deta/deta-go/deta"
	"github.com/deta/deta-go/service/base"
	"github.com/gofiber/fiber/v2"
)

var FetchItems = func(c *fiber.Ctx) error {
	id := c.Params("id")

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

	var items []*ItemsData
	_, err = itemsDB.Fetch(&base.FetchInput{
		Q:    base.Query{},
		Dest: &items,
	})

	if err != nil {
		return c.Status(500).JSON(lib.APIResponse{
			Error:   true,
			Message: err.Error(),
		})
	}

	return c.JSON(lib.APIResponse{
		Error: false,
		Data:  items,
	})
}
