package test

import (
	"gonum.org/v1/gonum/mat"
	"mulalg"
	"reflect"
	"testing"
)

var (
	testMatvecN  = 3
	benchMatvecN = 2000
	benchmarkM1  = mat.NewDense(benchMatvecN,
		benchMatvecN,
		makeRange(0, benchMatvecN*benchMatvecN))
	benchmarkV1 = mat.NewVecDense(benchMatvecN,
		makeRange(0, benchMatvecN))
	testM1 = mat.NewDense(testMatvecN,
		testMatvecN,
		makeRange(0, testMatvecN*testMatvecN))
	testV1 = mat.NewVecDense(testMatvecN,
		makeRange(0, testMatvecN))
)

func TestMatvec(t *testing.T) {
	expected := mat.NewVecDense(testMatvecN, make([]float64, testMatvecN))
	expected.MulVec(testM1, testV1)
	actual := mulalg.Matvec(testM1, testV1)
	if !reflect.DeepEqual(actual.RawVector(), expected.RawVector()) {
		fm := mat.Formatted(testM1, mat.Squeeze())
		fv := mat.Formatted(testV1, mat.Squeeze())
		fa := mat.Formatted(actual, mat.Squeeze())
		fe := mat.Formatted(expected, mat.Squeeze())

		t.Errorf("\n%v dot \n%v equals \n%v, expected \n%v;",
			fm, fv, fa, fe)
	}
}

func TestMatvecParallel(t *testing.T) {
	expected := mat.NewVecDense(testMatvecN, make([]float64, testMatvecN))
	expected.MulVec(testM1, testV1)
	actual := mulalg.MatvecParallel(testM1, testV1)
	if !reflect.DeepEqual(actual.RawVector(), expected.RawVector()) {
		fm := mat.Formatted(testM1, mat.Squeeze())
		fv := mat.Formatted(testV1, mat.Squeeze())
		fa := mat.Formatted(actual, mat.Squeeze())
		fe := mat.Formatted(expected, mat.Squeeze())

		t.Errorf("\n%v dot \n%v equals \n%v, expected \n%v;",
			fm, fv, fa, fe)
	}
}

func TestMatvecRecursiveParallel(t *testing.T) {
	expected := mat.NewVecDense(testMatvecN, make([]float64, testMatvecN))
	expected.MulVec(testM1, testV1)
	actual := mulalg.MatvecRecursiveParallel(testM1, testV1)
	if !reflect.DeepEqual(actual.RawVector(), expected.RawVector()) {
		fm := mat.Formatted(testM1, mat.Squeeze())
		fv := mat.Formatted(testV1, mat.Squeeze())
		fa := mat.Formatted(actual, mat.Squeeze())
		fe := mat.Formatted(expected, mat.Squeeze())

		t.Errorf("\n%v dot \n%v equals \n%v, expected \n%v;",
			fm, fv, fa, fe)
	}
}

func TestMatvecRecursiveParallelLogn(t *testing.T) {
	expected := mat.NewVecDense(testMatvecN, make([]float64, testMatvecN))
	expected.MulVec(testM1, testV1)
	actual := mulalg.MatvecRecursiveParallelLogn(testM1, testV1)
	if !reflect.DeepEqual(actual.RawVector(), expected.RawVector()) {
		fm := mat.Formatted(testM1, mat.Squeeze())
		fv := mat.Formatted(testV1, mat.Squeeze())
		fa := mat.Formatted(actual, mat.Squeeze())
		fe := mat.Formatted(expected, mat.Squeeze())
		t.Errorf("\n%v dot \n%v equals \n%v, expected \n%v;",
			fm, fv, fa, fe)
	}
}

func BenchmarkMatvec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mulalg.Matvec(benchmarkM1, benchmarkV1)
	}
}

func BenchmarkMatvecParallel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mulalg.MatvecParallel(benchmarkM1, benchmarkV1)
	}
}

func BenchmarkMatvecRecursiveParallel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mulalg.MatvecRecursiveParallel(benchmarkM1, benchmarkV1)
	}
}

func BenchmarkMatvecRecursiveParallelLogn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mulalg.MatvecRecursiveParallelLogn(benchmarkM1, benchmarkV1)
	}
}
