package main

import (
	"fmt"
	"strings"
)

func startGame() {
	var word string = getWord()
	var ans string

	var fails uint8 = 0
	var maxAttempts uint8 = 6

    wrongGuesses := []rune{}
    rightGuesses := []rune{}

	for fails != maxAttempts {
        var matchFound bool = false

        drawHangman(fails)
        printWord(word, &rightGuesses)

        // Check if all letters have been filled
        if checkLettersFilled(word, &rightGuesses) {
            fmt.Println("You win!")
            return
        }

        // Print all used letters
        if len(wrongGuesses) > 0 {
            fmt.Printf("Used letters: ")

            for _, ch := range wrongGuesses {
                fmt.Printf("%c ", ch)
            }

            fmt.Println()
        }

		fmt.Print("> ")
        fmt.Scan(&ans)

        ans = strings.ToLower(ans)
        ans = strings.TrimSpace(ans)

        // If guess matches the word
        if ans == word {
            fmt.Println("You win!")
            return
        }

        // If guess does not matches the word
        if ans != word && len(ans) > 1 {
            fails += 6 - fails
            break
        }

        ansChar := []rune(ans)[0]

        // Check if input already exist in wrongGuesses
        if checkGuesses(ansChar, &wrongGuesses) || checkGuesses(ansChar, &rightGuesses) {
            fmt.Printf("\nYou have already used '%c'!\n", ansChar)
            continue
        }
        
        for _, ch := range word {
            if ch == ansChar {
                rightGuesses = append(rightGuesses, ansChar)
                matchFound = true
            }
        }

        if !matchFound {
            fails += 1
            wrongGuesses = append(wrongGuesses, ansChar)
        }
	}

    drawHangman(fails)
    fmt.Println("You lose!")
    fmt.Printf("Actual word: %s\n", word)
}