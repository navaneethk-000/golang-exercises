package main

import (
	"strings"
	"unicode"
)

func Abbreviate(s string) string {
	words := strings.Fields(s)
	var sb strings.Builder

	for _, word := range words {
		firstChar := rune(word[0])

		if unicode.IsUpper(firstChar) {
			sb.WriteRune(firstChar)
		}
	}

	return sb.String()
}
