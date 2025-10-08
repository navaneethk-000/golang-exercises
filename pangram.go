package main

import (
	"fmt"
	"strings"
)

func Pangram(s string) bool {
	s = strings.ToLower(s)
	alphabets := "abcdefghijklmnopqrstuvwxyz"

	for i := 0; i < len(alphabets); i++ {
		if !strings.Contains(s, string(alphabets[i])) {
			return false
		}
	}
	return true
}

func main() {

	sentence := "The quick brown fox jumps over the lazy dog"
	fmt.Println(Pangram(sentence))

}
