package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Config struct {
	prevLocationURL *string
	nextLocationURL *string
}

type PokeLocationResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func GetPokeMap(requestUrl *string) (PokeLocationResponse, error) {
	url := "https://pokeapi.co/api/v2/location-area/"
	if requestUrl != nil {
		url = *requestUrl
	}
	response := PokeLocationResponse{}
	res, err := http.Get(url)
	if err != nil {
		return response, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}
