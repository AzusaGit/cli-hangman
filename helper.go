package main

import "fmt"

func printWord(word string, guessArr *[]rune) {
	guessMap := make(map[rune]bool)
	for _, ch := range *guessArr {
		guessMap[ch] = true
	}

	runes := []rune(word)

	for i := 0; i < len(runes); i++ {
		if !guessMap[runes[i]] {
			runes[i] = '_'
		}
	}

	fmt.Printf("Word: %s\n", string(runes))
}

func checkGuesses(ans rune, arr *[]rune) bool {
	for _, ch := range *arr {
		if ch == ans {
			return true
		}
	}

	return false
}

func checkLettersFilled(word string, guesses *[]rune) bool {
	if len(*guesses) == 0 {
		return false
	}

	letterMap := make(map[rune]bool)
	for _, ch := range word {
		letterMap[ch] = true
	}

	for _, ch := range *guesses {
		if !letterMap[ch] {
			return false
		}
	}

	for ch := range letterMap {
		found := false
		for _, guess := range *guesses {
			if ch == guess {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}