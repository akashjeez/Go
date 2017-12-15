package main
import "fmt"
func main() {
	a, b, c, d := 20, 10, 15, 5
	var e int;
	fmt.Println(" Value of A, B, C, D : ", a, b, c, d)
	e = (a + b) * c / d
	fmt.Println(" Value of (a + b) * c / d is :", e)
	e = ((a + b) * c) / d
	fmt.Println(" Value of ((a + b) * c) / d is :" , e)
	e = (a + b) * (c / d)
	fmt.Println(" Value of (a + b) * (c / d) is :", e)
	e = a + (b * c) / d
	fmt.Println(" Value of a + (b * c) / d is :" , e)
}
