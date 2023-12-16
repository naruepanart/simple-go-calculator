package main

import (
	"abc/calculator"
	"fmt"
)

func main() {
	calc := calculator.NewSimpleCalculator()

	add := calc.Add(5, 3)
	fmt.Println("5 + 3 =", add)

	subtract := calc.Subtract(10, 2)
	fmt.Println("10 - 2 =", subtract)

	multiply := calc.Multiply(8, 3)
	fmt.Println("8 * 3 =", multiply)

	divide := calc.Divide(8.0, 3)
	fmt.Println("8 / 3 =", divide)
}