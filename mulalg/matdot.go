package mulalg

import "gonum.org/v1/gonum/mat"

func MatDot(m1 *mat.Dense, m2 *mat.Dense) *mat.Dense {
	l1, n := m1.Dims()
	if _, n1 := m2.Dims(); n != n1 {
		panic(mat.ErrShape)
	}
	_, l2 := m2.Dims()

	res := mat.NewDense(l1, l2, nil)
	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			for k := 0; k < n; k++ {
				res.Set(i, j, res.At(i, j)+m1.At(i, k)*m2.At(k, j))
			}
		}
	}
	return res
}

// a goroutine calculate a row's value of new matrix
// create l1 number of thread
// DOP Θ(n3/n2) = Θ(n)
func MatDotParallel(m1 *mat.Dense, m2 *mat.Dense) *mat.Dense {
	l1, n := m1.Dims()
	if _, n1 := m2.Dims(); n != n1 {
		panic(mat.ErrShape)
	}
	_, l2 := m2.Dims()

	res := mat.NewDense(l1, l2, nil)
	ch := make(chan int)
	for i := 0; i < l1; i++ {
		go func(row int, sync chan int) {
			for j := 0; j < l2; j++ {
				for k := 0; k < n; k++ {
					res.Set(row, j, res.At(row, j)+m1.At(row, k)*m2.At(k, j))
				}
			}
			ch <- 1
		}(i, ch)
	}

	for i := 0; i < l1; i++ {
		<-ch
	}
	return res
}

// a goroutine calculate one value in result matrix
// create l1 * l2 number of thread
// too much threads lead to bad performance
// DOP Θ(n3/n) = Θ(n2)
func MatDotParallel2(m1 *mat.Dense, m2 *mat.Dense) *mat.Dense {
	l1, n := m1.Dims()
	if _, n1 := m2.Dims(); n != n1 {
		panic(mat.ErrShape)
	}
	_, l2 := m2.Dims()

	res := mat.NewDense(l1, l2, nil)
	ch := make(chan int)
	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			go func(row, col int, sync chan int) {
				for k := 0; k < n; k++ {
					res.Set(row, col, res.At(row, col)+m1.At(row, k)*m2.At(k, col))
				}
				ch <- 1
			}(i, j, ch)
		}
	}

	for i := 0; i < l1*l2; i++ {
		<-ch
	}
	return res
}
