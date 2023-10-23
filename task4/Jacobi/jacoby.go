package jacobi

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

type JacobiSolver struct {
	MaxIterations int
	Tolerance     float64
}

func NewJacobiResolver(maxIterations int, tolerance float64) *JacobiSolver {
	return &JacobiSolver{
		MaxIterations: maxIterations,
		Tolerance:     tolerance,
	}
}

// https://ru.wikipedia.org/wiki/%D0%9C%D0%B5%D1%82%D0%BE%D0%B4_%D0%AF%D0%BA%D0%BE%D0%B1%D0%B8
func (solver *JacobiSolver) Resolve(A mat.Matrix, b mat.Vector) (mat.Vector, error) {
	n, _ := A.Dims()
	x := mat.NewVecDense(n, nil)

	for k := 0; k < solver.MaxIterations; k++ {
		xNext := mat.NewVecDense(n, nil)

		for i := 0; i < n; i++ {
			sum := 0.0
			for j := 0; j < n; j++ {
				if i != j {
					sum += A.At(i, j) * x.At(j, 0)
				}
			}
			xNext.SetVec(i, (b.At(i, 0)-sum)/A.At(i, i))
		}

		diff := mat.NewVecDense(n, nil)
		diff.SubVec(xNext, x)

		if mat.Norm(diff, 2) < solver.Tolerance {
			fmt.Println(k)
			return xNext, nil
		}

		x = xNext
	}

	return x, nil
}
