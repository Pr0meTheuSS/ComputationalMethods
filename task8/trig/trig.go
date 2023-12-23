package trig

import (
	"gonum.org/v1/gonum/mat"
)

type triDiagonalMatrixSolver struct{}

func NewTrigResolver() *triDiagonalMatrixSolver {
	return &triDiagonalMatrixSolver{}
}

func (solver triDiagonalMatrixSolver) Resolve(A mat.Matrix, b mat.Vector) (mat.Vector, error) {
	//TODO: rename arrays like in explanation reference.

	// You can find explanation there:
	// https://pro-prof.com/forums/topic/sweep-method-for-solving-systems-of-linear-algebraic-equations
	// log.Println("=====================Trig method resolver was called.")
	// defer log.Println("Trig method resolver was finished.=====================")
	n, _ := A.Dims()
	x := make([]float64, n)

	// log.Println("Trig resolver set values for arrays.")
	var a, c, d []float64
	for i := 0; i < n; i++ {
		d = append(d, A.At(i, i))
	}

	a = append(a, 0)
	for i := 0; i < n-1; i++ {
		a = append(a, A.At(i, i+1))
	}

	for i := 0; i < n-1; i++ {
		c = append(c, A.At(i+1, i))
	}
	c = append(c, 0)

	alpha := make([]float64, n)
	beta := make([]float64, n)

	// straight
	alpha[0] = -c[0] / d[0]
	beta[0] = b.AtVec(0) / d[0]
	for i := 1; i < n-1; i++ {
		curgamma := (d[i] + a[i]*alpha[i-1])
		alpha[i] = -c[i] / curgamma
		beta[i] = (b.AtVec(i) - a[i]*beta[i-1]) / curgamma
	}
	beta[n-1] = (b.AtVec(n-1) - a[n-1]*beta[n-2]) / (d[n-1] + a[n-1]*alpha[n-2])

	// reverse
	x[n-1] = beta[n-1]
	for i := n - 2; i >= 0; i-- {
		x[i] = alpha[i]*x[i+1] + beta[i]
	}

	return mat.NewVecDense(n, x), nil
}
