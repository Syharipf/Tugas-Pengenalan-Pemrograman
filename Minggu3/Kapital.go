package main

import (
	"fmt"
	"unicode"
)

func isCapitalLetter(ch rune) bool {
	return unicode.IsUpper(ch)
}

func main() {
	var input rune

	fmt.Print("Masukkan sebuah huruf: ")
	fmt.Scanf("%c", &input)

	result := isCapitalLetter(input)

	fmt.Printf("Apakah huruf %c kapital? %t\n", input, result)
}
