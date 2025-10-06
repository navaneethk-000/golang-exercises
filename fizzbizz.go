package main
import "fmt"

func Fizzbiz(n int) {
	for i:=1;i<=n;i++ {

		if i%15 == 0 {
			fmt.Println("fizzbizz")
		} else if i%3 ==0 {
			fmt.Println("fizz")
		} else if i%5 ==0 {
			fmt.Println("bizz")
		}else {
			fmt.Println(i)
		}
	}
}

func main() {
	Fizzbiz(20)	
}