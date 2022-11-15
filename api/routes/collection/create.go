package col

import (
	"fmt"

	"github.com/TheBoringDude/zeta-library/api/lib"
	"github.com/gofiber/fiber/v2"
)

const reserved_base_name = "_all_collections"

var CreateCollection = func(c *fiber.Ctx) error {
	body := &lib.CollectionProps{}

	if err := c.BodyParser(body); err != nil {
		return c.Status(500).JSON(lib.APIResponse{
			Error:   true,
			Message: fmt.Sprintf("Error parsing request body. (%v)", err),
		})
	}

	if body.Name == reserved_base_name {
		return c.Status(400).JSON(lib.APIResponse{
			Error:   true,
			Message: fmt.Sprintf("Cannot use reserved base name `%s`", reserved_base_name),
		})
	}

	if !lib.Includes(body.Type, lib.CollectionTypes) {
		return c.Status(400).JSON(lib.APIResponse{
			Error:   true,
			Message: fmt.Sprintf("Invalid collection type: `%s`", body.Type),
		})
	}

	db, err := lib.InitBase("_all_collections")
	if err != nil {
		return c.Status(500).JSON(lib.APIResponse{
			Error:   true,
			Message: err.Error(),
		})
	}

	// include id and key
	id := lib.GenUniqueID()
	body.ID = id
	body.Key = id

	_, err = db.Put(body)
	if err != nil {
		return c.Status(500).JSON(lib.APIResponse{
			Error:   true,
			Message: err.Error(),
		})
	}

	return c.Status(200).JSON(lib.APIResponse{
		Error:   false,
		Data:    body,
		Message: "Collection has been created!",
	})
}
