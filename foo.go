package main
import(
	"fmt"
	"strings"
	)
type celsius float64
type Fahrenhite float64

func F2C(f Fahrenhite) celsius {
	c:= (f - 32) / 1.8
	return celsius(c)
}

func C2F(c celsius) Fahrenhite {
	f := (c * 9.0/5.0) + 32
	return Fahrenhite(f)
}

func main() {
	var c1 celsius = 100.0
	fmt.Printf("Celcius : %f, Fahrenhite : %f\n",c1,C2F(c1))

	s:= "golang"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c : %d %T \n",s[i],s[i],s[i])
	}

	if strings.Contains(s,"la"){
		fmt.Println(s, "contains x")
	} else {
		fmt.Println(s,"does not contain x")
	}
	fmt.Println(s)
}