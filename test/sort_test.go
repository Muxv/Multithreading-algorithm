package test

import (
	"mulalg"
	"sort"
	"testing"
)

var (
	length      int = 10
	benchlength int = 100000
	array           = makeRandomArray(length)
)

func TestMergeSort(t *testing.T) {
	actual := make([]int, length)
	copy(actual, array)
	mulalg.MergeSort(array)

	if !sort.SliceIsSorted(actual, func(i, j int) bool { return i <= j }) {
		t.Errorf("%v after sorting: \nget %v, which is not sorted", array, actual)
	}
}

func TestMergeSortParallel(t *testing.T) {
	actual := make([]int, length)
	copy(actual, array)
	mulalg.MergeSortParallel(array)

	if !sort.SliceIsSorted(actual, func(i, j int) bool { return i <= j }) {
		t.Errorf("%v after sorting: \nget %v, which is not sorted", array, actual)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := makeRandomArray(benchlength)
		mulalg.MergeSort(arr)
	}
}

func BenchmarkMergeSortParallel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := makeRandomArray(benchlength)
		mulalg.MergeSortParallel(arr)
	}
}
