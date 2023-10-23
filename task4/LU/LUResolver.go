package lu

import (
	"errors"
	"fmt"
	"main/common"
	"main/logger"
	"math"

	"gonum.org/v1/gonum/mat"
)

type luResolver struct {
	L mat.Matrix
	U mat.Matrix
}

func NewLUResolver() *luResolver {
	return &luResolver{
		U: nil,
		L: nil,
	}
}

func (lur luResolver) Resolve(A mat.Matrix, b mat.Vector) (mat.Vector, error) {
	if err := lur.LU(A); err != nil {
		return nil, err
	}

	return lur.SolveLU(b), nil
}

// TODO: расставить уточняющие комментарии
func (lur *luResolver) SolveLU(b mat.Vector) mat.Vector {
	l := logger.GetLoggerInstance()
	l.Info("=====================LU resolver method was called.")
	defer l.Info("LU resolver was finished.=====================")

	// Решение Ly = b
	n, _ := lur.L.Dims()
	y := mat.NewVecDense(n, nil)
	for i := 0; i < n; i++ {
		y.SetVec(i, b.At(i, 0))
		for j := 0; j < i; j++ {
			y.SetVec(i, y.At(i, 0)-lur.L.At(i, j)*y.At(j, 0))
		}
		y.SetVec(i, y.At(i, 0)/lur.L.At(i, i))
	}
	l.Debug("Resolved Ly = b.")

	// Решение Ux = y
	x := mat.NewVecDense(n, nil)
	for i := n - 1; i >= 0; i-- {
		x.SetVec(i, y.At(i, 0))
		for j := i + 1; j < n; j++ {
			x.SetVec(i, x.At(i, 0)-lur.U.At(i, j)*x.At(j, 0))
		}
		x.SetVec(i, x.At(i, 0)/lur.U.At(i, i))
	}
	l.Debug("Resolved Ux = y.")

	return x
}

func (lur *luResolver) LU(A mat.Matrix) error {
	l := logger.GetLoggerInstance()
	l.Info("=====================LU resolver called method LU to separate A like L*U.")
	defer l.Info("LU resolver finished separation A to L*U.=====================")

	r, c := A.Dims()
	U := mat.NewDense(r, c, nil)
	L := mat.NewDense(r, c, nil)

	U.CloneFrom(A)
	if !checkAbleLU(A) {
		err := errors.New("it is impossible to perform lu decomposition of this matrix. One of the corner minors is equal 0.")
		l.Error(err.Error())
		return err
	}
	l.Debug("U cloned from A.")

	for i := 0; i < r; i++ {
		for j := i; j < c; j++ {
			L.Set(j, i, U.At(j, i)/U.At(i, i))
		}
	}
	l.Debug("Init L values.")
	fmt.Println("Промежуточное значение Д")
	common.PrintMatrix(L)

	for k := 1; k < r; k++ {
		for i := k - 1; i < r; i++ {
			for j := i; j < c; j++ {
				L.Set(j, i, U.At(j, i)/U.At(i, i))
			}
		}

		fmt.Println("Промежуточное значение Д")
		common.PrintMatrix(L)

		for i := k; i < r; i++ {
			for j := k - 1; j < c; j++ {
				U.Set(i, j, U.At(i, j)-L.At(i, k-1)*U.At(k-1, j))
			}
		}
	}
	l.Debug("Values calculated for L and U.")

	lur.L, lur.U = L, U
	l.Debug("Values moved into structure.")

	return nil
}

func checkAbleLU(A mat.Matrix) bool {
	r, c := A.Dims()
	minSize := r
	if c < r {
		minSize = c
	}

	for i := 1; i <= minSize; i++ {
		submatrix := submatrix(A, 0, 0, i, i)
		det := mat.Det(submatrix)

		if math.Abs(det) <= 1e-9 {
			fmt.Println("Order", i)
			fmt.Println(det)
			return false
		}
	}

	return true
}

func submatrix(A mat.Matrix, i, j, m, n int) *mat.Dense {
	submatrix := mat.NewDense(m, n, nil)
	for ii := 0; ii < m; ii++ {
		for jj := 0; jj < n; jj++ {
			submatrix.Set(ii, jj, A.At(i+ii, j+jj))
		}
	}

	return submatrix
}
