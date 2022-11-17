package settings

import (
	"fmt"

	"github.com/TheBoringDude/zeta-library/api/lib"
	"github.com/gofiber/fiber/v2"
)

type SettingsProps struct {
	Key           string `json:"key,omitempty"`
	MoviedbApiKey string `json:"moviedbApiKey"`
}

var UpdateSettings = func(c *fiber.Ctx) error {
	body := &SettingsProps{
		Key: "_default",
	}

	if err := c.BodyParser(body); err != nil {
		return c.Status(500).JSON(lib.APIResponse{
			Error:   true,
			Message: fmt.Sprintf("Error parsing request body. (%v)", err),
		})
	}

	db, err := lib.InitBase(lib.RESERVED_CONFIG_NAME)
	if err != nil {
		return c.Status(500).JSON(lib.APIResponse{
			Error:   true,
			Message: err.Error(),
		})
	}

	if _, err := db.Put(body); err != nil {
		return c.Status(500).JSON(lib.APIResponse{
			Error:   true,
			Message: err.Error(),
		})
	}

	return c.JSON(lib.APIResponse{
		Error:   false,
		Message: "Successfully updated app settings.",
	})

}
