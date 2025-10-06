package main
import (
	"fmt"
	"strings"
	"unicode"
	)

func Abbrevate(s string){
	words := strings.Fields(s)
	shortform := ""

	for _, word := range words {
		firstChar := rune(word[0])

		if unicode.IsUpper(firstChar){
			shortform += string(firstChar)
		}	
	}
	fmt.Println(shortform)
}

func main() {
	Abbrevate("College of Engineering Trivandrum")
}