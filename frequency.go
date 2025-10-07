package main
import (
	"fmt"
	"strings"
)

func Frequency(s string) {

	frequency := make(map[string]int)
	s = strings.ToLower(s)

	for i:=0; i<len(s); i++ {
		ch := string(s[i])
		frequency[ch] += 1
	}

	fmt.Println(frequency)

}

func main(){
	Frequency("Okkey")
}