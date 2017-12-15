package main
import "fmt"
func main() {
	a, b := 21, 10
	fmt.Println(" Value of A and B :",a, b)
	fmt.Println(" Addition of A and B :", a+b)
	fmt.Println(" Subraction of A and B :", a-b)
	fmt.Println(" Multiplication of A and B :", a*b)
	fmt.Println(" Division of A and B :", a/b)
	b++	
	fmt.Println(" After Increment of B :", b)
	a--
	fmt.Println(" After Decrement of A :", a)
}
