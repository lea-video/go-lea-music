package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lea-video/go-lea-music/model"
)

func (rw *RouteWrapper) CreateArtistSolo(c *fiber.Ctx) error {
	// Parse the request body into a struct
	var newSolo model.ArtistSolo
	err := c.BodyParser(&newSolo)
	if err != nil {
		return err
	}

	solos, err := rw.DB.CreateArtistSolos([]*model.ArtistSolo{
		&newSolo,
	})
	if err != nil {
		return err
	}

	// build order id list + typecast artist
	keys := make([]int, 0, len(solos))
	typecastValues := make(map[int]*model.OneOfArtist)
	for key, val := range solos {
		keys = append(keys, key)
		typecastValues[key] = val.ToOneOf()
	}

	resp := model.ResponseObject{
		Order:   keys,
		Artists: typecastValues,
	}

	return c.JSON(resp)
}
