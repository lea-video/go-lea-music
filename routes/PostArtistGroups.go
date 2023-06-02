package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lea-video/go-lea-music/model"
)

func (rw *RouteWrapper) CreateArtistGroup(c *fiber.Ctx) error {
	// Parse the request body into a struct
	var newGroup model.ArtistGroup
	err := c.BodyParser(&newGroup)
	if err != nil {
		return err
	}

	groups, err := rw.DB.CreateArtistGroups([]*model.ArtistGroup{
		&newGroup,
	})
	if err != nil {
		return err
	}

	// build order id list + typecast artist
	keys := make([]int, 0, len(groups))
	typecast_values := make(map[int]*model.OneOfArtist)
	for key, val := range groups {
		keys = append(keys, key)
		typecast_values[key] = val.ToOneOf()
	}

	resp := model.ResponseObject{
		Order:   keys,
		Artists: typecast_values,
	}

	return c.JSON(resp)
}
