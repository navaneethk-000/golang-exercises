package main
import (
	"fmt"
	"strings"
)

func Panagram(s string){

	s = strings.ToLower(s)
	alphabets := "abcdefghijklmnopqrstuvwxyz"

	for _, ch := range alphabets {
		if !strings.Contains(s, string(ch)){
			fmt.Println("Not Panagram")
			return
		}
	}
	fmt.Println("Panagram")
}

func main() {
	Panagram("The quick brown fox jumps over the lazy dog")
}