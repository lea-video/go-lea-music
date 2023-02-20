package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lea-video/go-lea-music/model"
)

func (rw *RouteWrapper) UpdateArtist(c *fiber.Ctx) error {
	// read artist from body
	var artist model.Artist
	err := c.BodyParser(artist)
	if err != nil {
		return err
	}

	artists, err := rw.DB.UpdateArtists([]model.Artist{artist})
	if err != nil {
		return err
	}

	// reply with artist (including ID)
	return c.JSON(artists[0])
}
