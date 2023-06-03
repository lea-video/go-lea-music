package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lea-video/go-lea-music/model"
)

func (rw *RouteWrapper) CreateMediaTMPFile(c *fiber.Ctx) error {
	// Get the "mid" parameter from the URL
	mid, err := c.ParamsInt("mid")
	if err != nil {
		return err
	}

	tmpFile, err := model.NewTMPFile(mid)
	if err != nil {
		return err
	}

	tmpFiles, err := rw.DB.CreateTMPFiles([]*model.TMPFile{tmpFile})
	if err != nil {
		return err
	}

	// build order id list
	keys := make([]int, 0, len(tmpFiles))
	for key := range tmpFiles {
		keys = append(keys, key)
	}

	resp := model.ResponseObject{
		Order:   keys,
		TMPFile: tmpFiles,
	}

	return c.JSON(resp)
}
