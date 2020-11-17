package test

import (
	"gonum.org/v1/gonum/mat"
	"mulalg"
	"reflect"
	"testing"
)

var (
	testMatdotN  = 3
	benchMatdotN = 50
	benchmarkMl  = mat.NewDense(benchMatdotN,
		benchMatdotN,
		makeRange(0, benchMatdotN*benchMatdotN))
	benchmarkMr = mat.DenseCopyOf(benchmarkMl)

	testMl = mat.NewDense(testMatdotN, testMatdotN,
		makeRange(0, testMatdotN*testMatdotN))
	testMr = mat.DenseCopyOf(testMl)
)

func TestMatDot(t *testing.T) {
	expected := mat.NewDense(testMatdotN, testMatdotN, make([]float64, testMatdotN*testMatdotN))
	expected.Mul(testMl, testMr)
	actual := mulalg.MatDot(testMl, testMr)
	if !reflect.DeepEqual(actual.RawMatrix(), expected.RawMatrix()) {
		fml := mat.Formatted(testMl, mat.Squeeze())
		fmr := mat.Formatted(testMr, mat.Squeeze())
		fa := mat.Formatted(actual, mat.Squeeze())
		fe := mat.Formatted(expected, mat.Squeeze())

		t.Errorf("\n%v dot \n%v equals \n%v, expected \n%v;",
			fml, fmr, fa, fe)
	}
}

func TestMatDotParallel(t *testing.T) {
	expected := mat.NewDense(testMatdotN, testMatdotN, make([]float64, testMatdotN*testMatdotN))
	expected.Mul(testMl, testMr)
	actual := mulalg.MatDotParallel(testMl, testMr)
	if !reflect.DeepEqual(actual.RawMatrix(), expected.RawMatrix()) {
		fml := mat.Formatted(testMl, mat.Squeeze())
		fmr := mat.Formatted(testMr, mat.Squeeze())
		fa := mat.Formatted(actual, mat.Squeeze())
		fe := mat.Formatted(expected, mat.Squeeze())

		t.Errorf("\n%v dot \n%v equals \n%v, expected \n%v;",
			fml, fmr, fa, fe)
	}
}

func BenchmarkMatDot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mulalg.MatDot(benchmarkMl, benchmarkMr)
	}
}

func BenchmarkMatDotParallel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mulalg.MatDotParallel(benchmarkMl, benchmarkMr)
	}
}

func BenchmarkMatDotParallel2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mulalg.MatDotParallel2(benchmarkMl, benchmarkMr)
	}
}
