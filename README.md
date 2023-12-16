# Calculator Package

The Calculator package provides a simple implementation of a calculator with basic arithmetic operations. It includes methods for addition, subtraction, multiplication, and division.

## Installation

To use this package in your Go project, you can use the following go get command:

```bash
go get -u yourmodule/calculator
```

Replace 'yourmodule' with your actual module path.

## Usage

```go
package main

import (
	"fmt"
	"yourmodule/calculator"
)

func main() {
	// Create a new instance of the calculator
	calc := calculator.NewSimpleCalculator()

	// Perform arithmetic operations
	addResult := calc.Add(5, 3)
	fmt.Println("5 + 3 =", addResult)

	subtractResult := calc.Subtract(10, 2)
	fmt.Println("10 - 2 =", subtractResult)

	multiplyResult := calc.Multiply(8, 3)
	fmt.Println("8 * 3 =", multiplyResult)

	divideResult := calc.Divide(8.0, 3)
	fmt.Println("8 / 3 =", divideResult)
}
```

## API Reference

### `calculator.NewSimpleCalculator() Calculator`

Creates a new instance of the simple calculator.

### `calculator.Add(a, b int) int`

Adds two integers and returns the result.

### `calculator.Subtract(a, b int) int`

Subtracts the second integer from the first and returns the result.

### `calculator.Multiply(a, b int) int`

Multiplies two integers and returns the result.

### `calculator.Divide(a float64, b int) float64`

Divides the first integer by the second, handling division by zero, and returns the result.

## License

This calculator package is distributed under the MIT License. See the `LICENSE` file for more information.