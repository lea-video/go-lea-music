package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lea-video/go-lea-music/model"
)

func (rw *RouteWrapper) PostVariation(c *fiber.Ctx) error {
	// read variation from body
	variation := model.SongVariation{}
	err := c.BodyParser(variation)
	if err != nil {
		return err
	}

	// Create
	variations, err := rw.DB.CreateVariations([]model.SongVariation{variation})
	if err != nil {
		return err
	}

	// reply with variation (including ID)
	return c.JSON(variations[0])
}
