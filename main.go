package main

import (
	"fmt"
	"strings"
)

func main() {
	var isPlaying bool = true
	var choice string

	for isPlaying {
		startGame()

		fmt.Print("\nPlay again? (Y/n)\n> ")
		fmt.Scan(&choice)

        	choice = strings.ToLower(choice)
        	choice = strings.TrimSpace(choice)

        	if choice != "y" {
            		return
        	} 
	}
}
