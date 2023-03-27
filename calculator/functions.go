package calculator

import (
	"fmt"
	"math"
)

func Abs(value int) int {
	if value <= 0 {
		return value
	}
	return value
}

func IsSquareNumber(value int) bool {

	if sqrt := math.Sqrt(float64(value)); sqrt != float64(int(sqrt)) {
		return false
	} else {
		return true
	}
}

func RunOperations(operations string, a, b int) {
	switch operations {
	case "add":
		fmt.Println(Add(a, b))
	case "divide":
		fmt.Println(Divide(a, b))
	case "abs":
		fmt.Println(Abs(a))
	case "square":
		fmt.Println(IsSquareNumber(a))
	default:
		fmt.Println("Invalid operation")
	}

}
