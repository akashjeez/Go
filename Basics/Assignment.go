package main
import "fmt"
func main() {
	a := 21
	fmt.Println(" Value of A :",a)
	a += 2
	fmt.Println(" A += 2 :",a)
	a -= 3
	fmt.Println(" A -= 3 :",a)
	a *= 10
	fmt.Println(" A *= 10 :",a)
	a /= 2
	fmt.Println(" A /= 2 :",a)
	a <<= 2
	fmt.Println(" A <<= 2 :",a)
	a >>= 3
	fmt.Println(" A >>= 3 :",a)
}
