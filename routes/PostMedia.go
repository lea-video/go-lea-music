package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lea-video/go-lea-music/model"
)

func (rw *RouteWrapper) CreateMedia(c *fiber.Ctx) error {
	// Parse the request body into a struct
	var newMedia model.Media
	err := c.BodyParser(&newMedia)
	if err != nil {
		return err
	}

	media, err := rw.DB.CreateMedia([]*model.Media{
		&newMedia,
	})
	if err != nil {
		return err
	}

	// build order id list
	keys := make([]int, 0, len(media))
	for key := range media {
		keys = append(keys, key)
	}

	resp := model.ResponseObject{
		Order: keys,
		Media: media,
	}

	return c.JSON(resp)
}
