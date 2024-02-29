package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Breed struct {
	Display string `json:"display"`
	Key     string `json:"key"`
	Path    string `json:"path"`
}
type DogBreedsResponse struct {
	Status  string              `json:"status"`
	Message map[string][]string `json:"message"`
}

func getBreedsListHandler(c *gin.Context) {
	url := "https://dog.ceo/api/breeds/list/all"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer response.Body.Close()

	var apiResponse DogBreedsResponse
	err = json.NewDecoder(response.Body).Decode(&apiResponse)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	var BreedsList []Breed
	caser := cases.Title(language.English)
	if apiResponse.Status == "success" {
		for mainBreed, subBreeds := range apiResponse.Message {
			if len(subBreeds) == 0 {
				BreedsList = append(BreedsList, Breed{
					Display: caser.String(mainBreed),
					Key:     mainBreed,
					Path:    fmt.Sprintf("/%s", mainBreed),
				})
				continue
			}
			for _, sb := range subBreeds {
				BreedsList = append(BreedsList, Breed{
					Display: caser.String(fmt.Sprintf("%s %s", sb, mainBreed)),
					Key:     fmt.Sprintf("%s-%s", mainBreed, sb),
					Path:    fmt.Sprintf("/%s/%s", mainBreed, sb),
				})
			}
		}
	} else {
		fmt.Println("API request failed:", apiResponse.Status)
	}
	c.JSON(http.StatusOK, gin.H{"breeds": BreedsList})
	getPhotoForBreed(BreedsList[0].Path)
}

type BreedPhotoResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func getPhotoForBreed(breedPath string) (string, error) {
	url := fmt.Sprintf("https://dog.ceo/api/breed%s/images/random", breedPath)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making request:", err)
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("API request failed with status code:", response.Status)
		return "", err
	}

	var breedResponse BreedPhotoResponse
	err = json.NewDecoder(response.Body).Decode(&breedResponse)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return "", err
	}

	return breedResponse.Message, nil
}
