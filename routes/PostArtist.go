package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lea-video/go-lea-music/model"
)

func (rw *RouteWrapper) PostArtist(c *fiber.Ctx) error {
	// read artist from body
	artist, err := model.ReadArtist(c)
	if err != nil {
		return err
	}

	// Create
	artists, err := rw.DB.CreateArtists([]model.Artist{artist})
	if err != nil {
		return err
	}

	// reply with artist (including ID)
	return c.JSON(artists[0])
}
