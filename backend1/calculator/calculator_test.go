package calculator

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b, want float64
	}{
		{2, 3, 5},
		{-1, 1, 0},
		{0, 0, 0},
		{1.5, 2.5, 4},
		{-3.5, -2.5, -6},
	}
	for _, tc := range tests {
		got := Add(tc.a, tc.b)
		if got != tc.want {
			t.Errorf("Add(%v, %v) = %v; want %v", tc.a, tc.b, got, tc.want)
		}
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		a, b, want float64
	}{
		{5, 3, 2},
		{0, 0, 0},
		{-1, -1, 0},
		{10.5, 0.5, 10},
	}
	for _, tc := range tests {
		got := Subtract(tc.a, tc.b)
		if got != tc.want {
			t.Errorf("Subtract(%v, %v) = %v; want %v", tc.a, tc.b, got, tc.want)
		}
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		a, b, want float64
	}{
		{2, 3, 6},
		{0, 100, 0},
		{-2, 3, -6},
		{-2, -3, 6},
		{1.5, 2, 3},
	}
	for _, tc := range tests {
		got := Multiply(tc.a, tc.b)
		if got != tc.want {
			t.Errorf("Multiply(%v, %v) = %v; want %v", tc.a, tc.b, got, tc.want)
		}
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		a, b    float64
		want    float64
		wantErr error
	}{
		{6, 3, 2, nil},
		{7, 2, 3.5, nil},
		{0, 5, 0, nil},
		{5, 0, 0, ErrDivisionByZero},
	}
	for _, tc := range tests {
		got, err := Divide(tc.a, tc.b)
		if err != tc.wantErr {
			t.Errorf("Divide(%v, %v) error = %v; want %v", tc.a, tc.b, err, tc.wantErr)
		}
		if got != tc.want {
			t.Errorf("Divide(%v, %v) = %v; want %v", tc.a, tc.b, got, tc.want)
		}
	}
}

func TestExponentiate(t *testing.T) {
	tests := []struct {
		base, exp, want float64
	}{
		{2, 3, 8},
		{5, 0, 1},
		{9, 0.5, 3},
		{2, -1, 0.5},
	}
	for _, tc := range tests {
		got := Exponentiate(tc.base, tc.exp)
		if math.Abs(got-tc.want) > 1e-9 {
			t.Errorf("Exponentiate(%v, %v) = %v; want %v", tc.base, tc.exp, got, tc.want)
		}
	}
}

func TestSquareRoot(t *testing.T) {
	tests := []struct {
		a       float64
		want    float64
		wantErr error
	}{
		{9, 3, nil},
		{0, 0, nil},
		{2, math.Sqrt(2), nil},
		{-1, 0, ErrNegativeSquareRoot},
	}
	for _, tc := range tests {
		got, err := SquareRoot(tc.a)
		if err != tc.wantErr {
			t.Errorf("SquareRoot(%v) error = %v; want %v", tc.a, err, tc.wantErr)
		}
		if math.Abs(got-tc.want) > 1e-9 {
			t.Errorf("SquareRoot(%v) = %v; want %v", tc.a, got, tc.want)
		}
	}
}

func TestPercentage(t *testing.T) {
	tests := []struct {
		a, b, want float64
	}{
		{200, 10, 20},
		{50, 50, 25},
		{0, 100, 0},
		{100, 0, 0},
	}
	for _, tc := range tests {
		got := Percentage(tc.a, tc.b)
		if got != tc.want {
			t.Errorf("Percentage(%v, %v) = %v; want %v", tc.a, tc.b, got, tc.want)
		}
	}
}
