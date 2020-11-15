package test

import (
	"mulalg"
	"testing"
)

type intCase struct {
	in       int
	expected int
}

var fibTests = []intCase{
	{1, 1},
	{2, 1},
	{3, 2},
	{7, 13},
	{11, 89},
	{20, 6765},
}

var extremeTest = intCase{40, 102334155}

func TestFib(t *testing.T) {
	for _, tt := range fibTests {
		if actual := mulalg.Fib(tt.in); actual != tt.expected {
			t.Errorf("Fib{%d} = %d; expected %d", tt.in, actual, tt.expected)
		}
	}
}

func TestFibP(t *testing.T) {
	for _, tt := range fibTests {
		if actual := mulalg.FibParallel(tt.in); actual != tt.expected {
			t.Errorf("Fib{%d} = %d; expected %d", tt.in, actual, tt.expected)
		}
	}
}

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mulalg.Fib(35)
	}
}

func BenchmarkFibParallel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mulalg.FibParallel(35)
	}
}
