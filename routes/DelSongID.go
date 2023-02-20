package routes

import (
	"github.com/gofiber/fiber/v2"
)

func (rw *RouteWrapper) DelSongID(c *fiber.Ctx) error {
	songID, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	// Delete - delete product
	err = rw.DB.DeleteSongs([]int{songID})
	if err != nil {
		return err
	}

	return c.JSON(true)
}
