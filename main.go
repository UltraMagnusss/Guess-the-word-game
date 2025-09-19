package main

import "fmt"

func main() {
	dictionary := map[string]string{ // dictionary
		"Apple":  "A fruit that keeps doctors away",
		"Peach":  "A fruit that represents the female butt on the internet",
		"Batman": "The Iron man in the DC films",
		"Sun":    "It takes 8min for it's light to travel to Earth from it",
		"Saturn": "It has a beautiful rigns of ice and stones around it",
	}
	fmt.Println("Greetings Friend, Guess the word using the given hint below")

	var answer string

	for word, hint := range dictionary {
		fmt.Println("A hint:", hint)
		fmt.Println("Write down the word: ")
		for {
			fmt.Scan(&answer)

			if answer == word {
				fmt.Println("Congradulations, You guessed the word:", word)
				break
			} else {
				fmt.Println("Wrong, Try again!")
			}
		}
	}
	fmt.Println("You correctly guessed all the word. Game Over!")
}
