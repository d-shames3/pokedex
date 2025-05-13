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
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type ExploreResponse struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func UnmarshalPokeapiResponse(res *http.Response, responseType string) (any, error) {
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code: %d and\nbody : %v", res.StatusCode, res.Body)
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
