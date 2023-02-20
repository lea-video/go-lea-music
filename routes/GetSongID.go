package routes

import (
	"github.com/gofiber/fiber/v2"
)

func (rw *RouteWrapper) GetSongID(c *fiber.Ctx) error {
	songID, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	// Read
	songs, err := rw.DB.FetchSongs([]int{songID})
	if err != nil {
		return err
	}

	// reply with artist
	return c.JSON(songs[0])
}
