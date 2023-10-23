package conjugate

import (
	"errors"
	"fmt"

	"gonum.org/v1/gonum/mat"
)

type ConjugateGradientSolver struct {
	MaxIterations int
	Tolerance     float64
}

func NewConjugateGradientSolver(maxIterations int, tolerance float64) *ConjugateGradientSolver {
	return &ConjugateGradientSolver{
		MaxIterations: maxIterations,
		Tolerance:     tolerance,
	}
}

func (solver *ConjugateGradientSolver) Resolve(A mat.Matrix, b mat.Vector) (mat.Vector, error) {
	if !isSymmetric(A) {
		return nil, errors.New("Conjugate gradients method for SLAE resolving works only with symmetric matrix.")
	}

	n, _ := A.Dims()
	x := mat.NewVecDense(n, nil)

	r := mat.NewVecDense(n, nil)
	rNext := mat.NewVecDense(n, nil)
	p := mat.NewVecDense(n, nil)

	Ax := mat.NewVecDense(n, nil)
	Ax.MulVec(A, x)
	r.SubVec(b, Ax)

	p.CloneFromVec(r)

	rsold := mat.Dot(r, r)

	for k := 0; k < solver.MaxIterations; k++ {
		fmt.Println(k)
		alpha := rsold / mat.Dot(p, p)
		x.AddScaledVec(x, alpha, p)

		rNext.CloneFromVec(r)
		Ax.MulVec(A, x)
		rNext.SubVec(b, Ax)
		rsnew := mat.Dot(rNext, rNext)

		if rsnew < solver.Tolerance {
			return x, nil
		}

		beta := rsnew / rsold
		p.AddScaledVec(rNext, beta, p)
		r.CloneFromVec(rNext)
		rsold = rsnew
	}

	return x, nil
}

func isSymmetric(mat mat.Matrix) bool {
	rows, cols := mat.Dims()
	if rows != cols {
		return false
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if mat.At(i, j) != mat.At(j, i) {
				return false
			}
		}
	}

	return true
}
