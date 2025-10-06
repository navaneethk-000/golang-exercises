package main
import "fmt"

func Palindrome(s string) {
	
	reverse := ""

	for i:=len(s)-1;i>=0;i--{
		
		reverse += string(s[i])
	}
	if reverse == s{
		fmt.Println("Palindrome")
	}else {
		fmt.Println("Not Palindrome")
	}
}

func main()  {
	Palindrome("oko")
}