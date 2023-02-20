package routes

import (
	"github.com/gofiber/fiber/v2"
)

func (rw *RouteWrapper) DelVariationID(c *fiber.Ctx) error {
	variationID, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	// Delete - delete product
	err = rw.DB.DeleteSongVariations([]int{variationID})
	if err != nil {
		return err
	}

	return c.JSON(true)
}
