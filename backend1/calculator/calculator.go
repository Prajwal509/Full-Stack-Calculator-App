// Package calculator provides pure arithmetic functions.
// All business logic lives here, decoupled from HTTP transport.
package calculator

import (
	"errors"
	"math"
)

var (
	ErrDivisionByZero     = errors.New("division by zero")
	ErrNegativeSquareRoot = errors.New("cannot compute square root of a negative number")
	ErrInvalidPercentage  = errors.New("invalid percentage operation")
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
		return 0, ErrDivisionByZero
	}
	return a / b, nil
}

func Exponentiate(base, exponent float64) float64 {
	return math.Pow(base, exponent)
}

func SquareRoot(a float64) (float64, error) {
	if a < 0 {
		return 0, ErrNegativeSquareRoot
	}
	return math.Sqrt(a), nil
}

// Percentage returns (a * b / 100), e.g. "what is b% of a?"
func Percentage(a, b float64) float64 {
	return a * b / 100
}
