package client

import (
	"github.com/starwar/model"
)

type StarWar interface {
	GetPerson(string) (*model.StarWar, error)
}

func GetStarWarClient(version string) StarWar {
	switch version {
	case "v1":
		return NewSwClient()
	case "v2":
		return NewSwClientV2()
	default:
		return NewSwClient()
	}
}