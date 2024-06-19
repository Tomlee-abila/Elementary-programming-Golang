package main

import "github.com/01-edu/z01"

func main(){
	alphabet := "zYxWvUtSrQpOnMlKjIhGfEdCbA"
	for _, letter := range alphabet{
		z01.PrintRune(letter)
	}	
	z01.PrintRune('\n')
}