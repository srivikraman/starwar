package client

import (
	"github.com/starwar/model"
	"io/ioutil"
	"fmt"
	"net/http"
	"encoding/json"
)

type swClient struct {
}

func NewSwClient() swClient   {
	return swClient{}
}

func (sw swClient) GetPerson(personID string) (*model.StarWar, error)  {
	resp, err := http.Get("https://swapi.dev/api/people/" + personID)
	if err != nil {
		fmt.Println("Error: ", err.Error());
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error while reading: ", err.Error());
		return nil, err
	}

	starWar := model.StarWar{}

	err = json.Unmarshal(body, &starWar)
	if err != nil {
		fmt.Println("Error while unmarshal: ", err.Error());
		return nil, err
	}

	return &starWar, nil
}
