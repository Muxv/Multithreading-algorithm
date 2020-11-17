package mulalg

import (
	"runtime"
)

const limit = 1e8

func Sum() int {
	total := 0
	for i := 0; i < limit; i++ {
		total += i
	}
	return total
}

// DOP Θ(n/1) = Θ(n)
func SumParallel() int {
	nCPU := runtime.GOMAXPROCS(0)
	ch := make(chan int)
	for i := 0; i < nCPU; i++ {
		go func(i int, out chan<- int) {
			sum := 0
			start := (limit / nCPU) * i
			end := start + (limit / nCPU)
			for j := start; j < end; j += 1 {
				sum += j
			}
			out <- sum
		}(i, ch)
	}

	total := 0
	for i := 0; i < nCPU; i++ {
		total += <-ch
	}
	return total
}
