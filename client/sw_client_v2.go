package client

import (
	"github.com/starwar/model"
)

type swClientV2 struct {
}

func NewSwClientV2() swClientV2   {
	return swClientV2{}
}

func (sw swClientV2) GetPerson(personID string) (*model.StarWar, error)  {
	//star war version 2 api coming soon, for now mock response
	starWar := model.StarWar{Name: "Vikram"}
	return &starWar, nil
}
