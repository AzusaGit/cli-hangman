package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getWord() (string, error) {
	// Might need to change the API
	api := "https://random-word-api.herokuapp.com/word"

	res, err := http.Get(api)
	if err != nil {
		return "", fmt.Errorf("API unavailable: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API error: received status code %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %v", err)
	}

	var randWord []string
	err = json.Unmarshal(body, &randWord)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	if len(randWord) == 0 {
		return "", fmt.Errorf("API returned an empty response")
	}

	return randWord[0], nil
}

