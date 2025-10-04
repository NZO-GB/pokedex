package main

var locationURL = "https://pokeapi.co/api/v2/location-area/"


type mapCall struct {
	next *string
	prev *string
}

type mapStruct struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type cliCommand struct {
	name        string
	description string
	callback    func(*configStruct) error
}

