package mulalg

import (
	"gonum.org/v1/gonum/mat"
)

// transpose a nxn matrix
func Transpose(m *mat.Dense) *mat.Dense {
	r, c := m.Dims()
	if r != c {
		panic(mat.ErrShape)
	}
	res := mat.NewDense(r, c, nil)
	for i := 0; i < r; i++ {
		for j := 0; j <= i; j++ {
			res.Set(i, j, m.At(j, i))
			res.Set(j, i, m.At(i, j))
		}
	}
	return res
}

// DOP Θ(n2/n) = Θ(n)
func TransposeParallel(m *mat.Dense) *mat.Dense {
	r, c := m.Dims()
	if r != c {
		panic(mat.ErrShape)
	}
	res := mat.NewDense(r, c, nil)
	ch := make(chan int)
	for i := 0; i < r; i++ {
		go func(row int, sync chan int) {
			for j := 0; j <= row; j++ {
				res.Set(row, j, m.At(j, row))
				res.Set(j, row, m.At(row, j))
			}
			ch <- 1
		}(i, ch)
	}
	for i := 0; i < r; i++ {
		<-ch
	}
	return res
}
