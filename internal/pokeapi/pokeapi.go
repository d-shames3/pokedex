package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func CallPokeApi(endpoint string) (*http.Response, error) {
	res, err := http.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("error calling %s: %v", endpoint, err)
	}

	return res, nil
}

type LocationResponse struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
	} `json:"results"`
}

type ExploreResponse struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func UnmarshalPokeapiResponse(res *http.Response, responseType string) (any, error) {
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var result any

	switch responseType {
	case "location":
		var response LocationResponse
		err = json.Unmarshal(body, &response)
		if err != nil {
			return nil, fmt.Errorf("error decoding response body: %v\nerror: %v", body, err)
		}
		result = response
	case "explore":
		var response ExploreResponse
		err = json.Unmarshal(body, &response)
		if err != nil {
			return nil, fmt.Errorf("error decoding response body: %v\nerror: %v", body, err)
		}
		result = response
	}

	return result, nil
}
