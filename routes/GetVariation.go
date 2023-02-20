package routes

import (
	"github.com/gofiber/fiber/v2"
)

func (rw *RouteWrapper) GetVariation(c *fiber.Ctx) error {
	variations, err := rw.DB.ListSongVariations()
	if err != nil {
		return err
	}

	// reply with variation
	return c.JSON(variations[0])
}
