package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lea-video/go-lea-music/model"
)

func (rw *RouteWrapper) GetArtistGroups(c *fiber.Ctx) error {
	groups, err := rw.DB.GetArtistGroups()
	if err != nil {
		return err
	}

	// build order id list + typecast artist
	keys := make([]int, 0, len(groups))
	typecastValues := make(map[int]*model.OneOfArtist)
	for key, val := range groups {
		keys = append(keys, key)
		typecastValues[key] = val.ToOneOf()
	}

	resp := model.ResponseObject{
		Order:   keys,
		Artists: typecastValues,
	}

	resp, err = doExpansions(c, rw.DB, resp)
	if err != nil {
		return err
	}

	return c.JSON(resp)
}
