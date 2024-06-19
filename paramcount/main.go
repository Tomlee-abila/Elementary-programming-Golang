package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main(){
	printNumber(len(os.Args)-1)
	z01.PrintRune('\n')
}

func printNumber(number int){
	var digits []rune

	for number != 0{
		digits = append(digits, rune(number%10 + '0'))
		number /= 10
	}

	for i := 0; i < len(digits); i++{
		z01.PrintRune(digits[len(digits)-1-i])
	}
	
}