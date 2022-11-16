package col

import (
	"fmt"

	"github.com/TheBoringDude/zeta-library/api/lib"
	"github.com/deta/deta-go/deta"
	"github.com/deta/deta-go/service/base"
	"github.com/gofiber/fiber/v2"
)

type UpdateColBody struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

var UpdateCollection = func(c *fiber.Ctx) error {
	id := c.Params("id")
	body := &UpdateColBody{}

	if err := c.BodyParser(body); err != nil {
		return c.Status(500).JSON(lib.APIResponse{
			Error:   true,
			Message: fmt.Sprintf("Error parsing request body. (%v)", err),
		})
	}

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

	updates := base.Updates{}
	if body.Name == "" && body.Description == "" {
		return c.Status(204).JSON(lib.APIResponse{
			Error:   true,
			Message: "Nothing to update.",
		})
	}
	if body.Name != "" {
		updates["name"] = body.Name
	}
	if body.Description != "" {
		updates["description"] = body.Description
	}

	if err = db.Update(id, updates); err != nil {
		return c.Status(500).JSON(lib.APIResponse{
			Error:   true,
			Message: err.Error(),
		})
	}

	return c.JSON(lib.APIResponse{
		Error:   false,
		Message: "Successfully updated collection.",
	})
}
