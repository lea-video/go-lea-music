package routes

import (
	"github.com/gofiber/fiber/v2"
)

func (rw *RouteWrapper) GetSong(c *fiber.Ctx) error {
	songs, err := rw.DB.ListSongs()
	if err != nil {
		return err
	}

	// reply with song
	return c.JSON(songs[0])
}
