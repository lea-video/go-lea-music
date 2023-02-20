package routes

import (
	"github.com/gofiber/fiber/v2"
)

func (rw *RouteWrapper) GetArtistID(c *fiber.Ctx) error {
	artistID, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	// Read
	artists, err := rw.DB.FetchArtists([]int{artistID})
	if err != nil {
		return err
	}

	// reply with artist
	return c.JSON(artists[0])
}
