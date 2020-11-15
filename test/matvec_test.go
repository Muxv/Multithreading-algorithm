package test

import (
	"gonum.org/v1/gonum/mat"
	"mulalg"
	"testing"
)

var (
	n          = 1000
	matContent = makeRange(0, n*n)
	vecContent = makeRange(0, n)
	testMatrix = mat.NewDense(n, n, matContent)
	testVec    = mat.NewVecDense(n, vecContent)
)

func makeRange(start int, end int) []float64 {
	a := make([]float64, end-start)
	for i, _ := range a {
		a[i] = float64(start + i)
	}
	return a
}

func BenchmarkMatvec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mulalg.Matvec(testMatrix, testVec)
	}
}

func BenchmarkMatvecParallel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mulalg.MatvecParallel(testMatrix, testVec)
	}
}
