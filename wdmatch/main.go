package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	firstWord := os.Args[1]
	secondWord := os.Args[2]

	result := wordMatch2(firstWord, secondWord)

	if result {
		for _, letter := range firstWord {
			z01.PrintRune(letter)
		}
		z01.PrintRune('\n')
	}
}

func wordMatch(firstWord string, secondWord string) bool {
	result := false

	for i, letter1 := range firstWord {
		for j, leletter2 := range secondWord {
			if letter1 == leletter2 {
				if i == len(firstWord)-1 {
					return true
				} else {
					return wordMatch(firstWord[i+1:], secondWord[j+1:])
				}
			}

			if i == len(secondWord)-1 {
				return false
			}
		}
	}

	return result
}

func wordMatch2(firstWord, secondWord string) bool {
	i, j := 0, 0

	for i < len(firstWord) && j < len(secondWord) {
		if firstWord[i] == secondWord[j] {
			i++
		}
		j++
	}

	return i == len(firstWord)
}
