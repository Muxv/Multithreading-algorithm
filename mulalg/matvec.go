package mulalg

import "gonum.org/v1/gonum/mat"

func Matvec(m *mat.Dense, x *mat.VecDense) *mat.VecDense {
	l := x.Len()
	if l != m.ColView(0).Len() || l != m.RowView(0).Len() {
		panic("Input size does not match")
	}
	v := mat.NewVecDense(l, make([]float64, l))
	for i := 0; i < l; i++ {
		sum := 0.0
		for j := 0; j < l; j++ {
			sum += m.At(i, j) * x.AtVec(i)
		}
		v.SetVec(i, sum)
	}
	return v
}

func MatvecParallel(m *mat.Dense, x *mat.VecDense) *mat.VecDense {
	l := x.Len()
	if l != m.ColView(0).Len() || l != m.RowView(0).Len() {
		panic("Input size does not match")
	}
	v := mat.NewVecDense(l, make([]float64, l))

	for i := 0; i < l; i++ {
		go func(row int) {
			sum := 0.0
			for column := 0; column < l; column++ {
				sum += m.At(row, column) * x.AtVec(row)
			}
			v.SetVec(row, sum)
		}(i)
	}
	return v
}
