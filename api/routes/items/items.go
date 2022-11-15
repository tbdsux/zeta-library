package items

import (
	"fmt"

	"github.com/TheBoringDude/zeta-library/api/lib"
	"github.com/deta/deta-go/deta"
	"github.com/gofiber/fiber/v2"
)

type ItemsBodyProps struct {
	ID   string      `json:"id"`
	Data []ItemsData `json:"data"`
}

type ItemsData struct {
	Type      string `json:"type"`
	Image     string `json:"image,omitempty"`
	Name      string `json:"name"`
	Url       string `json:"url"`
	ItemID    string `json:"item_id"`
	Key       string `json:"key,omitempty"`
	DateAdded int    `json:"date_added"`
}

var UpdateItems = func(c *fiber.Ctx) error {
	body := &ItemsBodyProps{}

	if err := c.BodyParser(body); err != nil {
		return c.Status(500).JSON(lib.APIResponse{
			Error:   true,
			Message: fmt.Sprintf("Error parsing request body. (%v)", err),
		})
	}

	if len(body.Data) > 25 {
		return c.Status(400).JSON(lib.APIResponse{
			Error:   true,
			Message: "Too many items to put.",
		})
	}

	db, err := lib.InitBase("_all_collections")
	if err != nil {
		return c.Status(500).JSON(lib.APIResponse{
			Error:   true,
			Message: err.Error(),
		})
	}

	_, err = lib.GetCollection(db, body.ID)
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

	itemsDB, err := lib.InitBase(body.ID)
	if err != nil {
		return c.Status(500).JSON(lib.APIResponse{
			Error:   true,
			Message: err.Error(),
		})
	}

	keys, err := itemsDB.PutMany(body.Data)
	if err != nil {
		return c.Status(500).JSON(lib.APIResponse{
			Error:   true,
			Message: err.Error(),
		})
	}

	return c.JSON(lib.APIResponse{
		Error:   false,
		Data:    keys,
		Message: "Successfully added items.",
	})
}
