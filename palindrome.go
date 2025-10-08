package main

import "fmt"

func Palindrome(s string) bool {

	reverse := ""

	for i := len(s) - 1; i >= 0; i-- {
		reverse += string(s[i])
	}

	return reverse == s
}

func main() {

	word := "malayalam"

	if Palindrome(word) {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not Palindrome")
	}

}
