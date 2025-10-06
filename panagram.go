package main
import (
	"fmt"
	"strings"
)

func Panagram(s string){

	s = strings.ToLower(s)
	alphabets := "abcdefghijklmnopqrstuvwxyz"

	for i:=0; i<len(alphabets); i++ {
		if !strings.Contains(s, string(alphabets[i])) {
			fmt.Println("Not Panagram")
			return
		}
	}
	fmt.Println("Panagram")
}

func main() {
	Panagram("The quick brown fox jumps over the lazy dog")
}