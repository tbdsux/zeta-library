package settings

import (
	"github.com/TheBoringDude/zeta-library/api/lib"
	"github.com/deta/deta-go/deta"
	"github.com/gofiber/fiber/v2"
)

var GetSettings = func(c *fiber.Ctx) error {
	db, err := lib.InitBase(lib.RESERVED_CONFIG_NAME)
	if err != nil {
		return c.Status(500).JSON(lib.APIResponse{
			Error:   true,
			Message: err.Error(),
		})
	}

	settings := &SettingsProps{}
	if err := db.Get("_default", settings); err == deta.ErrNotFound {
		// try to create key if it does not exist
		settings.Key = "_default"
		if _, err = db.Put(settings); err != nil {
			return c.Status(500).JSON(lib.APIResponse{
				Error:   true,
				Message: err.Error(),
			})
		}
	} else if err != nil {
		return c.Status(500).JSON(lib.APIResponse{
			Error:   true,
			Message: err.Error(),
		})
	}

	return c.JSON(lib.APIResponse{
		Error: false,
		Data:  settings,
	})
}
