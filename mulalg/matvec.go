package mulalg

import (
	"gonum.org/v1/gonum/mat"
	"sync"
)

func Matvec(m *mat.Dense, x *mat.VecDense) *mat.VecDense {
	l := x.Len()
	if l != m.ColView(0).Len() || l != m.RowView(0).Len() {
		panic("Input size does not match")
	}
	v := mat.NewVecDense(l, make([]float64, l))
	for i := 0; i < l; i++ {
		sum := 0.0
		for j := 0; j < l; j++ {
			sum += m.At(i, j) * x.AtVec(j)
		}
		v.SetVec(i, sum)
	}
	return v
}

// DOP Θ(n2/n) = Θ(n)
func MatvecParallel(m *mat.Dense, x *mat.VecDense) *mat.VecDense {
	l := x.Len()
	if l != m.ColView(0).Len() || l != m.RowView(0).Len() {
		panic("Input size does not match")
	}
	v := mat.NewVecDense(l, make([]float64, l))
	ch := make(chan int)

	for i := 0; i < l; i++ {
		go func(row int, out chan int) {
			sum := 0.0
			for column := 0; column < l; column++ {
				sum += m.At(row, column) * x.AtVec(column)
			}
			v.SetVec(row, sum)
			out <- 0
		}(i, ch)
	}
	for i := 0; i < l; i++ {
		<-ch
	}
	return v
}

func MatvecRecursiveParallel(m *mat.Dense, x *mat.VecDense) *mat.VecDense {
	l := x.Len()
	if l != m.ColView(0).Len() || l != m.RowView(0).Len() {
		panic("Input size does not match")
	}
	y := mat.NewVecDense(l, make([]float64, l))
	matvecMainLoop(m, x, y, 0, l-1)
	return y
}

// T(n) = T(n/2), T(1) = Θ(n)
// => T(n) = Θ(n)
// DOP Θ(n2/n) = Θ(n)
// but use Recursive
func matvecMainLoop(m *mat.Dense,
	x *mat.VecDense,
	y *mat.VecDense,
	start int,
	end int) {
	if start == end {
		l := x.Len()
		for j := 0; j < l; j++ {
			y.SetVec(start, y.AtVec(start)+m.At(start, j)*x.AtVec(j))
		}
	} else {
		mid := (start + end) / 2
		ch := make(chan int)

		go func(start, end int, sync chan int) {
			matvecMainLoop(m, x, y, start, mid)
			sync <- 0
		}(start, mid, ch)

		go func(start, end int, sync chan int) {
			matvecMainLoop(m, x, y, mid+1, end)
			sync <- 0
		}(mid+1, end, ch)

		<-ch
		<-ch
	}
}

//
func MatvecRecursiveParallelLogn(m *mat.Dense, x *mat.VecDense) *mat.VecDense {
	l := x.Len()
	if l != m.ColView(0).Len() || l != m.RowView(0).Len() {
		panic("Input size does not match")
	}
	y := mat.NewVecDense(l, make([]float64, l))
	wg := sync.WaitGroup{}
	for i := 0; i < l; i++ {
		wg.Add(1)
		go func(row int, done func()) {
			defer done()
			v := matvecMainLoopLogn(mat.VecDenseCopyOf(m.RowView(row)), x, 0, l-1)
			y.SetVec(row, v)
		}(i, wg.Done)
	}
	wg.Wait()
	return y
}

// DOP Θ(n2/logn)
func matvecMainLoopLogn(m_part *mat.VecDense,
	x *mat.VecDense,
	start int,
	end int) float64 {
	if start == end {
		return m_part.AtVec(start) * x.AtVec(start)
	} else {
		mid := (start + end) / 2
		chl := make(chan float64)
		chr := make(chan float64)
		go func(ch chan float64) {
			chl <- matvecMainLoopLogn(m_part, x, start, mid)
		}(chl)

		go func(ch chan float64) {
			chr <- matvecMainLoopLogn(m_part, x, mid+1, end)
		}(chr)

		return <-chl + <-chr
	}
}
