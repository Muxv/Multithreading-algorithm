package test

import "math/rand"

// [a, b)
func makeRange(start int, end int) []float64 {
	a := make([]float64, end-start)
	for i, _ := range a {
		a[i] = float64(start + i)
	}
	return a
}

func makeRandomArray(length int) []int {
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(length)
	}
	return arr
}
