package routes

import (
	"github.com/gofiber/fiber/v2"
)

func (rw *RouteWrapper) GetVariationID(c *fiber.Ctx) error {
	variationID, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	// Read
	variations, err := rw.DB.FetchSongVariations([]int{variationID})
	if err != nil {
		return err
	}

	// reply with artist
	return c.JSON(variations[0])
}
