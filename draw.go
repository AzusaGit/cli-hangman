package main

import "fmt"

func drawHangman(step uint8) {
	var hangman = [7]string{
		"\n-------\n|     |\n|\n|\n|\n|\n|\n△\n",     					    // Step 1: Gallows
		"\n-------\n|     |\n|     〇\n|\n|\n|\n△\n",     					// Step 2: Head
		"\n-------\n|     |\n|     〇\n|     |\n|     |\n|\n△\n",     		// Step 3: Body
		"\n-------\n|     |\n|     〇\n|    /|\n|     |\n|\n△\n",     		// Step 4: Left Arm
		"\n-------\n|     |\n|     〇\n|    /|\\\n|     |\n|\n△\n",   		// Step 5: Right Arm
		"\n-------\n|     |\n|     〇\n|    /|\\\n|     |\n|    / \n△\n",   // Step 6: Left Leg
		"\n-------\n|     |\n|     〇\n|    /|\\\n|     |\n|    / \\\n△\n", // Step 7: Right Leg
	}

	fmt.Println(hangman[step])
}