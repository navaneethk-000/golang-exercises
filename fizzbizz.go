package main

import (
	"fmt"
	"strconv"
)

func Fizzbiz(n int) []string {
	var output []string
	for i := 1; i <= n; i++ {

		if i%15 == 0 {
			output = append(output, "fizzbizz")
		} else if i%3 == 0 {
			output = append(output, "fizz")
		} else if i%5 == 0 {
			output = append(output, "bizz")
		} else {
			output = append(output, strconv.Itoa(i))
		}
	}
	return output
}

func main() {
	fmt.Println(Fizzbiz(20))
}
