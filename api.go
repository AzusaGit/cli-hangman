package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func getWord() string {
	api := "https://random-word-api.herokuapp.com/word"

	res, err := http.Get(api)
	if err != nil {
		log.Fatal("Error while fetching API: ", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error reading response body: ", err)
	}

	var randWord []string

	err = json.Unmarshal(body, &randWord)
	if err != nil {
		log.Fatal("Error unmarshalling JSON:", err)
	}

	return randWord[0]
}