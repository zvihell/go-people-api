package controllers

import (
	"encoding/json"
	"fmt"
	"go-people-api/internal/models"
	"io"
	"net/http"
)

func getRichorDieTrying(user models.User) models.User {
	resp, err := http.Get(fmt.Sprintf("https://api.agify.io/?name=%s", user.Name))
	if err != nil {
		fmt.Errorf("Error connect to agify", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
	}

	var agifyResponse models.AgifyResponse
	json.Unmarshal(body, &agifyResponse)

	user.Age = agifyResponse.Age

	resp, err = http.Get(fmt.Sprintf("https://api.genderize.io/?name=%s", user.Name))
	if err != nil {
		fmt.Errorf("Error connect to api genderize", err)
	}

	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
	}

	var genderizeResponse models.GenderizeResponse
	json.Unmarshal(body, &genderizeResponse)

	user.Gender = genderizeResponse.Gender

	resp, err = http.Get(fmt.Sprintf("https://api.nationalize.io/?name=%s", user.Name))
	if err != nil {
		fmt.Errorf("Error connect to api nationalize", err)
	}

	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
	}

	var nationalizeResponse models.NationalizeResponse
	json.Unmarshal(body, &nationalizeResponse)

	if len(nationalizeResponse.Country) > 0 {
		user.Nationality = nationalizeResponse.Country[0].CountryID
	}

	return user
}
