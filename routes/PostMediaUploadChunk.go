package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lea-video/go-lea-music/model"
)

func (rw *RouteWrapper) AppendMediaTMPFileChunk(c *fiber.Ctx) error {
	// Get the "uid" parameter from the URL
	uid, err := c.ParamsInt("uid")
	if err != nil {
		return err
	}

	tmpFile, err := rw.DB.GetTMPFileByID([]int{uid})
	if err != nil {
		return err
	}
	if len(tmpFile) == 0 {
		return c.SendStatus(fiber.ErrNotFound.Code)
	}

	err = rw.FileDB.AppendFile(
		tmpFile[0].Location,
		c.Body(),
	)
	if err != nil {
		return err
	}

	// TODO: update tmp file max age
	// TODO: if all uploaded mark tmpFile as finished

	return c.JSON(model.ResponseObject{})
}
