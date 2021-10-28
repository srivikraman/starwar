package main;

import (
	"fmt"
	"github.com/starwar/client"
	"github.com/starwar/model"
)

func main() {
	starWar := client.GetStarWarClient("v1")
	fmt.Println("StarWar v1 client")
	personDetails, err := starWar.GetPerson("1")
	if err != nil {
		fmt.Println("API call failed")
		return
	}
	fmt.Println("Name: ", personDetails.Name)

	starWarV2 := client.GetStarWarClient("v2")
	fmt.Println("\nStarWar v2 client mock")
	personDetails, _ = starWarV2.GetPerson("1")
	fmt.Println("Name: ", personDetails.Name)

	fmt.Println("\nConcurretly perform three request")
	persons := GetPersonsConcurrently(starWar, []string{"1", "2", "3"})

	for _, person := range persons {
		if person.Error != nil {
			fmt.Println("Error: ", person.Error.Error())
		}
		fmt.Println("Name: ", person.Data.Name)
	}
}

func GetPersonsConcurrently(client client.StarWar,identifiers []string) []model.StarWarResponse {
	starWars := []model.StarWarResponse{}

	channels := []chan model.StarWarResponse{}
	for _, identifier := range identifiers {
		channel := make(chan model.StarWarResponse)
		channels = append(channels, channel)
		go GetPersonResponse(client, identifier, channel)
	}

	for _, channel := range channels {
		starWar := <- channel
		starWars = append(starWars, starWar)
	}

	return starWars
}

func GetPersonResponse(starWarClient client.StarWar, identifier string, responseChan chan model.StarWarResponse)  {
	defer close(responseChan)
	response, err := starWarClient.GetPerson(identifier)
	channelResponse := model.StarWarResponse{Data: response, Error: err}
	responseChan <- channelResponse
}

