package test

import (
	"gonum.org/v1/gonum/mat"
	"mulalg"
	"reflect"
	"testing"
)

var (
	testTransposeN  = 3
	benchTransposeN = 2000
	benchmarkM2     = mat.NewDense(benchTransposeN,
		benchTransposeN,
		makeRange(1, benchTransposeN*benchTransposeN+1))
	testM2 = mat.NewDense(testTransposeN,
		testTransposeN,
		makeRange(1, testTransposeN*testTransposeN+1))
)

func TestTranspose(t *testing.T) {
	e := testM2.T()
	expected := mat.DenseCopyOf(e)
	actual := mulalg.Transpose(testM2)

	if !reflect.DeepEqual(expected.RawMatrix(), actual.RawMatrix()) {
		fraw := mat.Formatted(testM2)
		fe := mat.Formatted(expected, mat.Squeeze())
		fa := mat.Formatted(actual, mat.Squeeze())
		t.Errorf("\n%v.T\n equals \n%v\n expected \n%v;\n",
			fraw, fa, fe)
	}
}

func TestTransposeParallel(t *testing.T) {
	e := testM2.T()
	expected := mat.DenseCopyOf(e)
	actual := mulalg.TransposeParallel(testM2)

	if !reflect.DeepEqual(expected.RawMatrix(), actual.RawMatrix()) {
		fraw := mat.Formatted(testM2)
		fe := mat.Formatted(expected, mat.Squeeze())
		fa := mat.Formatted(actual, mat.Squeeze())
		t.Errorf("\n%v.T\n equals \n%v\n expected \n%v;\n",
			fraw, fa, fe)
	}
}

func BenchmarkTranspose(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mulalg.Transpose(benchmarkM2)
	}
}

func BenchmarkTransposeParallel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mulalg.TransposeParallel(benchmarkM2)
	}
}
