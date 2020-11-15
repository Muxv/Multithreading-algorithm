package test

import (
	"mulalg"
	"testing"
)

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mulalg.Sum()
	}
}

func BenchmarkSumParallel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mulalg.SumParallel()
	}
}
