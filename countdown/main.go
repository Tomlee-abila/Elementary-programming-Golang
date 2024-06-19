package main

import "github.com/01-edu/z01"

func main() {
	digits := "9876543210"

	for _, digit := range digits {
		z01.PrintRune(digit)
	}
	z01.PrintRune('\n')
}
