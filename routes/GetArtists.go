package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lea-video/go-lea-music/model"
)

func (rw *RouteWrapper) GetArtists(c *fiber.Ctx) error {
	artists, err := rw.DB.GetArtists()
	if err != nil {
		return err
	}

	// build order id list
	keys := make([]int, 0, len(artists))
	for key := range artists {
		keys = append(keys, key)
	}

	resp := model.ResponseObject{
		Order:   keys,
		Artists: artists,
	}

	resp, err = doExpansions(c, rw.DB, resp)
	if err != nil {
		return err
	}

	return c.JSON(resp)
}
