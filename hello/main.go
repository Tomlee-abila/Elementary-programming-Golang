package main

import "github.com/01-edu/z01"

func main(){
	greeting := "Hello World!"
	for _, letter := range greeting{
		z01.PrintRune(letter)
	}	
	z01.PrintRune('\n')
}