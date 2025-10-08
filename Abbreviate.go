package main

import (
	"fmt"
	"strings"
	"unicode"
)

func Abbreviate(s string) string {
	words := strings.Fields(s)
	shortform := ""

	for _, word := range words {
		firstChar := rune(word[0])

		if unicode.IsUpper(firstChar) {
			shortform += string(firstChar)
		}
	}
	return shortform
}

func main() {
	result := Abbreviate("College of Engineering Trivandrum")
	fmt.Println(result)
}
