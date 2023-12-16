package calculator

import "fmt"

// Calculator defines the methods for basic arithmetic operations
type Calculator interface {
	Add(int, int) int
	Subtract(int, int) int
	Multiply(int, int) int
	Divide(float64, int) float64
}

type simpleCalculator struct{}

// NewSimpleCalculator returns a new instance of the simpleCalculator struct
func NewSimpleCalculator() Calculator {
	return simpleCalculator{}
}

// Add performs addition
func (c simpleCalculator) Add(a, b int) int {
	return a + b
}

// Subtract performs subtraction
func (c simpleCalculator) Subtract(a, b int) int {
	return a - b
}

// Multiply performs multiplication
func (c simpleCalculator) Multiply(a, b int) int {
	return a * b
}

// Divide performs division
func (c simpleCalculator) Divide(a float64, b int) float64 {
	if b == 0 {
		fmt.Println("Error: cannot divide by zero")
		return 0
	}
	return a / float64(b)
}