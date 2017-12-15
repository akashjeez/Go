package main
import "fmt"
func main() {
	var grade string = "B"
	var marks int = 90
	switch marks {
		case 90: grade = "A"
		case 80: grade = "B"
		case 50,60,70 : grade = "C"
		default: grade = "D"
	}

	switch {
		case grade == "A" :
			fmt.Println(" Excellent!")
		case grade == "B", grade == "C" :
			fmt.Println(" Well done")
		case grade == "D" :
			fmt.Println(" You passed")
		case grade == "F":
			fmt.Println(" Better try again")
		default:
			fmt.Println(" Invalid grade")
	}
	
	fmt.Println(" Your grade :", grade)
}
