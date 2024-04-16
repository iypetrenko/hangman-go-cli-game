package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("wordlist.json")
	if err != nil {
		log.Fatal(err)
	}
	var words []string
	err = json.Unmarshal(data, &words)
	if err != nil {
		log.Fatal(err)
	}
	word := words[rand.Intn(len(words))]
	lives := len(word) * 2
	var blanks []string
	for range word {
		blanks = append(blanks, "_")
	}
	for {
		fmt.Printf("♥️%d,Word: %s Letter: ", lives, strings.Join(blanks, " "))
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			log.Fatal(err)
		}

		input = strings.ToLower(input)
		fmt.Println(input)
		for _, inputLetter := range input {
			correctGuess := false
			for i, wordLetter := range word {
				if inputLetter == wordLetter {
					blanks[i] = string(inputLetter)
					correctGuess = true
				}
			}
			if !correctGuess {
				lives--
			}
		}
		if lives <= 0 {
			fmt.Printf("♥️ 0, Word: %s - sorry,you lost!", word)
			break
		}
		if word == strings.Join(blanks, "") {
			fmt.Printf("♥️ %d 🎉🎉🎉You guessed the word: %s - congratulations!", lives, word)
			break
		}
	}
}
