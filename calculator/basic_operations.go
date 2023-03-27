package calculator

const Pi = 3.121592654

func Add(a int, b int) int {
	// var sum int = a + b
	// var sum = a + b
	sum := a + b
	return sum
}

func Divide(a, b int) (quotient int, remainder int) {
	quotient = a / b
	remainder = a % b
	return
}
