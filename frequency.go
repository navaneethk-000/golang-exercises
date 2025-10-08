package main
import (
	"fmt"
	"strings"
)

func Frequency(s string) {

	s = strings.ToLower(s)
	frequency := make(map[string]int)

	for i:=0; i<len(s); i++ {
		ch := string(s[i])
		frequency[ch] += 1
	}

	fmt.Println(frequency)

}

func main(){
	Frequency("Okkey")
}