package calculator_test

import (
	"go-calculate/calculator"
	"math"
	"testing"
)

func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestAdd(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 4, want: 6},
		{a: 0, b: -3, want: -3},
		{a: 5, b: 45, want: 50},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	type tesCase struct {
		a, b float64
		want float64
	}
	testCases := []tesCase{
		{a: 4, b: 4, want: 0},
		{a: 20, b: 5, want: 15},
		{a: -5, b: -5, want: 0},
	}
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Subtract(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 3, want: 6},
		{a: 0, b: 5, want: 0},
		{a: -5, b: 6, want: -30},
		{a: 2, b: .4, want: .8},
	}
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Mutiply(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 9, b: 3, want: 3},
		{a: -1, b: -1, want: 1},
		{a: 10, b: 2, want: 5},
		{a: 1, b: 3, want: 0.333333},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		if err != nil {
			t.Fatalf("Divide(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
		if !closeEnough(tc.want, got, 0.001) {
			t.Errorf("Divide(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}

}

func TestDivideInvalid(t *testing.T) {
	t.Parallel()
	_, err := calculator.Divide(6, 0)
	if err == nil {
		t.Error("Divide(6,0): want error for invalid input, got nil")
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	type tesCase struct {
		num, want float64
	}
	testCases := []tesCase{
		{num: 49, want: 7},
		{num: 36, want: 6},
		{num: 64, want: 8},
		{num: 10, want: 3.162278},
	}
	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.num)
		if err != nil {
			t.Fatalf("Sqrt(%f): want %f, got %f", tc.num, tc.want, got)
		}
		if !closeEnough(tc.want, got, 0.001) {
			t.Errorf("Sqrt(%f): want %f, got %f", tc.num, tc.want, got)
		}
	}
}

func TestSqrtInvalid(t *testing.T) {
	t.Parallel()
	_, err := calculator.Sqrt(-36)
	if err == nil {
		t.Error("Sqrt(-36) want error for invalid, for nil")
	}
}
