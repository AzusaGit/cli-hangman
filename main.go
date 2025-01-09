package main

import (
	"fmt"
	"strings"
)

func main() {
	var isPlaying bool = true
	var choice string

	for isPlaying {
		var wordRes, err = getWord()

		if err != nil {
			fmt.Printf("\n%s\n", err)
			fmt.Print("Do you want to try again? (Y/n)\n> ")
			fmt.Scan(&choice)

			choice = strings.ToLower(choice)
			choice = strings.TrimSpace(choice)

			if choice != "y" {
				return
			}

			continue
		}

		startGame(wordRes)

		fmt.Print("\nPlay again? (Y/n)\n> ")
		fmt.Scan(&choice)

        choice = strings.ToLower(choice)
        choice = strings.TrimSpace(choice)

        if choice != "y" {
        	return
        } 
	}
}
