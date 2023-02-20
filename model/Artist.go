package model

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

const (
	TypeSolo  = "SOLO"
	TypeGroup = "GROUP"
)

type Artist interface {
	GetCore() CoreArtist
}

func ReadArtist(c *fiber.Ctx) (Artist, error) {
	var core CoreArtist
	err := c.BodyParser(&core)
	if err != nil {
		return nil, err
	}
	if core.Type == TypeSolo {
		var artist SoloArtist
		err := c.BodyParser(&artist)
		return artist, err
	} else if core.Type == TypeGroup {
		var artist GroupArtist
		err := c.BodyParser(&artist)
		return artist, err
	} else {
		return nil, errors.New("invalid artist object")
	}

}

type CoreArtist struct {
	ID   int
	Type string
	Name string
	Tags []int
}

type SoloArtist struct {
	CoreArtist
}

func (sa SoloArtist) GetCore() CoreArtist {
	return sa.CoreArtist
}

type GroupArtist struct {
	CoreArtist
	Members []int
}

func (ga GroupArtist) GetCore() CoreArtist {
	return ga.CoreArtist
}
