package routes

import (
	"github.com/gofiber/fiber/v2"
)

func (rw *RouteWrapper) GetArtist(c *fiber.Ctx) error {
	artists, err := rw.DB.ListArtists()
	if err != nil {
		return err
	}

	// reply with all artists
	return c.JSON(artists)
}
