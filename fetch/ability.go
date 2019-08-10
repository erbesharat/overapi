package fetch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/erbesharat/goverapi/db"
)

type AbilityResponse struct {
	Next string       `json:"next"`
	Data []db.Ability `json:"data"`
}

func fetchAbilities(url string) ([]db.Ability, string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, "", fmt.Errorf("Couldn't get the abilities total number: %s", err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, "", fmt.Errorf("Couldn't read the overwatch API's response: %s", err.Error())
	}

	overwatchResponse := AbilityResponse{}
	if err := json.Unmarshal(body, &overwatchResponse); err != nil {
		return nil, "", fmt.Errorf("Couldn't parse the overwatch API's response: %s", err.Error())
	}

	return overwatchResponse.Data, overwatchResponse.Next, nil
}

func FetchAbilities(db *db.DB) error {
	next := "https://overwatch-api.net/api/v1/ability/"
	for next != "" {
		data, nextPage, err := fetchAbilities(next)
		if err != nil {
			return err
		}
		next = nextPage
		for _, abilityy := range data {
			db.InsertAbility(&abilityy)
		}
	}
	return nil
}
