package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lea-video/go-lea-music/model"
)

func (rw *RouteWrapper) PostSong(c *fiber.Ctx) error {
	// read song from body
	song := model.Song{}
	err := c.BodyParser(&song)
	if err != nil {
		return err
	}

	// Create
	songs, err := rw.DB.CreateSong([]model.Song{song})
	if err != nil {
		return err
	}

	// reply with song (including ID)
	return c.JSON(songs[0])
}
