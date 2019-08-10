package fetch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/erbesharat/goverapi/db"
)

type HeroesResponse struct {
	Next string    `json:"next"`
	Data []db.Hero `json:"data"`
}

func fetchHeros(url string) ([]db.Hero, string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, "", fmt.Errorf("Couldn't get the heroes total number: %s", err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, "", fmt.Errorf("Couldn't read the overwatch API's response: %s", err.Error())
	}

	overwatchResponse := HeroesResponse{}
	if err := json.Unmarshal(body, &overwatchResponse); err != nil {
		return nil, "", fmt.Errorf("Couldn't parse the overwatch API's response: %s", err.Error())
	}

	return overwatchResponse.Data, overwatchResponse.Next, nil
}

func FetchHeros(db *db.DB) error {
	next := "https://overwatch-api.net/api/v1/hero/"
	for next != "" {
		data, nextPage, err := fetchHeros(next)
		if err != nil {
			return err
		}
		next = nextPage
		for _, hero := range data {
			db.InsertHero(&hero)
		}
	}
	return nil
}
