package main

type mapCall struct {
	next string
	prev string
}

type configStruct struct {
	mapCall mapCall
	mapFull mapStruct
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