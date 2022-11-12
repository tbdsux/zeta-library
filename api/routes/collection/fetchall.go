package col

import (
	"github.com/TheBoringDude/zeta-library/api/lib"
	"github.com/deta/deta-go/service/base"
	"github.com/gofiber/fiber/v2"
)

var FetchAllCollections = func(c *fiber.Ctx) error {
	db, err := lib.InitBase("_all_collections")
	if err != nil {
		return c.Status(500).JSON(lib.APIResponse{
			Error:   true,
			Message: err.Error(),
		})
	}

	var cols []*CollectionProps
	_, err = db.Fetch(&base.FetchInput{
		Q:    base.Query{},
		Dest: &cols,
	})

	if err != nil {
		return c.Status(500).JSON(lib.APIResponse{
			Error:   true,
			Message: err.Error(),
		})
	}

	return c.JSON(lib.APIResponse{
		Error: false,
		Data:  cols,
	})
}
