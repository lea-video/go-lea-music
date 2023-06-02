package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lea-video/go-lea-music/model"
)

func (rw *RouteWrapper) GetArtistSolos(c *fiber.Ctx) error {
	solos, err := rw.DB.GetArtistSolos()
	if err != nil {
		return err
	}

	// build order id list + typecast artist
	keys := make([]int, 0, len(solos))
	typecast_values := make(map[int]*model.OneOfArtist)
	for key, val := range solos {
		keys = append(keys, key)
		typecast_values[key] = val.ToOneOf()
	}

	resp := model.ResponseObject{
		Order:   keys,
		Artists: typecast_values,
	}

	resp, err = doExpansions(c, rw.DB, resp)
	if err != nil {
		return err
	}

	return c.JSON(resp)
}
