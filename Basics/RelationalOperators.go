package main
import "fmt"
func main() {
	a, b := 21, 10
	
	if( a == b ) {
		fmt.Printf("Line 1 - a is equal to b\n" )
	} else {
		fmt.Printf("Line 1 - a is not equal to b\n" )
	}

	if ( a < b ) {
		fmt.Printf("Line 2 - a is less than b\n" )
	} else {
		fmt.Printf("Line 2 - a is not less than b\n" )
	}

	if ( a > b ) {
		fmt.Printf("Line 3 - a is greater than b\n" )
	} else {
		fmt.Printf("Line 3 - a is not greater than b\n" )
	}

	if ( a <= b ) {
		fmt.Printf("Line 4 - a is either less than or equal to b\n" )
	}
	if ( b >= a ) {
		fmt.Printf("Line 5 - b is either greater than or equal to b\n" )
	}
}
