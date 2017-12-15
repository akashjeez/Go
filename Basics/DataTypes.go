package main
import "fmt"
func main() {
	//Datatypes
	//Integer Types - uint8, uint16, uint32, uint64, int8, int16, int32, int64
	//Floating Types - float32, float64
	//Complex Types - complex64, complex128
	//Other Numeric Types - byte, rune, uint, int, uintptr
	var fval float64
	fval = 20.0
	x1,x2 := 12,23
	fmt.Println("Addition of x1 and x2 variable : ",x1+x2) 
	fmt.Println("Value of fval variable:",fval)
	fmt.Printf("Type of fval variable: %T\n",fval)
}
