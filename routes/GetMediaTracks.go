package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lea-video/go-lea-music/model"
)

func (rw *RouteWrapper) GetMediaTracks(c *fiber.Ctx) error {
	// Get the "mid" parameter from the URL
	mid, err := c.ParamsInt("mid")
	if err != nil {
		return err
	}

	tracks, err := rw.DB.GetMediaTracks([]int{mid})
	if err != nil {
		return err
	}

	// build order id list
	keys := make([]int, 0, len(tracks))
	for key := range tracks {
		keys = append(keys, key)
	}

	resp := model.ResponseObject{
		Order:      keys,
		MediaTrack: tracks,
	}

	resp, err = doExpansions(c, rw.DB, resp)
	if err != nil {
		return err
	}

	return c.JSON(resp)
}
