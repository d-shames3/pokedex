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

type PokeapiUnnamedResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func UnmarshalPokeapiResponse(res *http.Response) (PokeapiUnnamedResponse, error) {
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return PokeapiUnnamedResponse{}, fmt.Errorf("request failed with status code: %d and\nbody : %v\n", res.StatusCode, res.Body)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return PokeapiUnnamedResponse{}, fmt.Errorf("error reading response body: %v", err)
	}

	response := PokeapiUnnamedResponse{}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return PokeapiUnnamedResponse{}, fmt.Errorf("error decoding response body: %v\nerror: %v", body, err)
	}

	return response, nil

}
