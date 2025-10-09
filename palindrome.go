package main

import (
	"strings"
)

func Palindrome(s string) bool {
	var sb strings.Builder

	for i := len(s) - 1; i >= 0; i-- {
		sb.WriteByte(s[i])
	}
	reverse := sb.String()

	return reverse == s
}
