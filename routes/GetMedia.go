package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lea-video/go-lea-music/model"
)

func (rw *RouteWrapper) GetMedia(c *fiber.Ctx) error {
	meds, err := rw.DB.GetMedia()
	if err != nil {
		return err
	}

	// build order id list
	keys := make([]int, 0, len(meds))
	for key := range meds {
		keys = append(keys, key)
	}

	resp := model.ResponseObject{
		Order: keys,
		Media: meds,
	}

	resp, err = doExpansions(c, rw.DB, resp)
	if err != nil {
		return err
	}

	return c.JSON(resp)
}
