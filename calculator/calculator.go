// Package calulato provides a library for simple
// Calculations in Go
package calculator

import (
	"errors"
	"math"
)

func Add(a, b float64) float64 {
	return a + b
}

func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero not allowed")
	}
	return a / b, nil
}

func Sqrt(num float64) (float64, error) {
	if math.Signbit(num) {
		return 0, errors.New("negetive numbers are invalid")
	}
	return math.Sqrt(num), nil
}
