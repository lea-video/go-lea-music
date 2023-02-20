package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lea-video/go-lea-music/model"
)

func (rw *RouteWrapper) UpdateSong(c *fiber.Ctx) error {
	// read song from body
	song := model.Song{}
	err := c.BodyParser(&song)
	if err != nil {
		return err
	}

	songs, err := rw.DB.UpdateSongs([]model.Song{song})
	if err != nil {
		return err
	}

	// reply with song (including ID)
	return c.JSON(songs[0])
}
