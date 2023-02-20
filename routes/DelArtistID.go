package routes

import (
	"github.com/gofiber/fiber/v2"
)

func (rw *RouteWrapper) DelArtistID(c *fiber.Ctx) error {
	artistID, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	err = rw.DB.DeleteArtists([]int{artistID})
	if err != nil {
		return err
	}

	return c.JSON(true)
}
