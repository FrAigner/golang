package main

import (
	"fmt"
	"math/rand"

	"github.com/FrAigner/golang/helloworld/calculator"
)

func main() {
	fmt.Println("Hello, World!", rand.Intn(10))
	fmt.Println("23 + 42 =", calculator.Add(23, 42))
	fmt.Println(calculator.Divide(17, 3))

	fmt.Println("Sum of 1 to 10 =", calculator.Sum(1, 10))
	fmt.Println(calculator.SumUntil(10))
	// calculator.SumInfinite()
	fmt.Println(calculator.IsSquareNumber(24))
	calculator.RunOperations("add", 1, 2)
}
