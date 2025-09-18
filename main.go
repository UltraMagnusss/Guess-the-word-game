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
	fmt.Println("Greetings Friend")
	word := "Apple"
	hint := dictionary[word]
	fmt.Println("A hint:", hint)
}
