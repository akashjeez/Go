package main
import "fmt"

//bool, string
//int  int8  int16  int32  int64
//uint uint8 uint16 uint32 uint64 uintptr
//byte // alias for uint8
//rune // alias for int32, represents a Unicode code point
//float32 float64
//complex64 complex128

//func add(x float64, y float64) float64 {
func add(x, y float64) float64 {
	return x+y
}

func multiple(a,b string) (string,string) {
	return a,b
}

func main() {
	//var num1 float64 = 3.5
	//var num2 float64 = 5.4
	//var a int = 4
	//var num1, num2 float64 = 3.6, 5.4
	num1, num2 := 3.6, 5.4
	fmt.Println(" num2 float64:",num2," converted into num3 int:",int(num2))
	fmt.Println(" Addition of num1 & num 2: ",add(num1,num2))
	w1, w2 := "Hey", "there"
	fmt.Println(multiple(w1,w2))
}
