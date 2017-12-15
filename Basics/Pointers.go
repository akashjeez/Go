package main
import "fmt"
func main() {
	a := 21
	var ptr *int
	fmt.Println(" Value of A :", a);
	fmt.Printf(" Type of A : %T\n", a);
	ptr = &a
	fmt.Println(" ptr = &a -> ptr : ", ptr)
	fmt.Println(" *ptr :", *ptr);
}
