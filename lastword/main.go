package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main(){
	if len(os.Args) != 2{
		return
	}
	word := lastWord(os.Args[1])
	for _, letter := range word{
		z01.PrintRune(letter)
	}
	z01.PrintRune('\n')
}

func lastWord(s string)string{
	var words []string
	firstletter := false
	word := ""

	for i, letter := range s{
		if letter != ' '{
			firstletter = true
			word += string(letter)
		}

		if letter == ' ' && firstletter{
			firstletter = false
			words = append(words, word)
			word = ""
		}
		if i == len(s)-1 && firstletter{
			words = append(words, word)
		}
	}

	return words[len(words)-1]
}