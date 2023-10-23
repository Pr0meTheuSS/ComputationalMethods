package qr

import (
	"fmt"
	"main/common"
	"main/logger"

	"gonum.org/v1/gonum/mat"
)

type qrResolver struct {
	Q mat.Matrix
	R mat.Matrix
}

func NewQRResolver() *qrResolver {
	return &qrResolver{
		Q: nil,
		R: nil,
	}
}

func (qrr qrResolver) Resolve(A mat.Matrix, b mat.Vector) (mat.Vector, error) {
	qrr.QR(A)
	return qrr.SolveQR(b), nil
}

func (qrr *qrResolver) SolveQR(b mat.Vector) mat.Vector {
	l := logger.GetLoggerInstance()
	l.Info("=====================QR resolver method was called.")
	defer l.Info("QR resolver was finished.=====================")

	_, n := qrr.R.Dims()
	x := mat.NewVecDense(n, nil)
	Qt := qrr.Q.T()

	Qtb := mat.NewVecDense(n, nil)
	Qtb.MulVec(Qt, b)

	for i := n - 1; i >= 0; i-- {
		x.SetVec(i, Qtb.AtVec(i))
		for j := i + 1; j < n; j++ {
			x.SetVec(i, x.AtVec(i)-qrr.R.At(i, j)*x.AtVec(j))
		}
		x.SetVec(i, x.AtVec(i)/qrr.R.At(i, i))
	}

	l.Debug("Resolved Rx = Qt * b.")

	return x
}

func (qrr *qrResolver) QR(A mat.Matrix) {
	l := logger.GetLoggerInstance()
	l.Info("=====================QR resolver called method QR to separate A like Q * R.")
	defer l.Info("QR resolver finished separation A to Q * R.=====================")

	l.Debug("Q and R initialized.")
	Q := gramSchmidt(A)
	r, c := A.Dims()
	R := mat.NewDense(r, c, nil)
	Qt := Q.T()
	R.Mul(Qt, A)
	l.Debug("QR decomposition completed.")

	qrr.Q, qrr.R = Q, R

	// Выводим значения Q и R (можете заменить на logger)
	fmt.Println("Q:")
	common.PrintMatrix(Q)
	fmt.Println("R:")
	common.PrintMatrix(R)
}

func getColumn(A mat.Matrix, col int) mat.VecDense {
	m, _ := A.Dims()
	data := []float64{}
	for i := 0; i < m; i++ {
		data = append(data, A.At(i, col))
	}
	return *mat.NewVecDense(m, data)
}

func normalize(v mat.Vector) mat.VecDense {
	data := make([]float64, v.Len())
	copy(data, mat.Col(nil, 0, v))
	norm := mat.Norm(v, 2)
	for i := range data {
		data[i] /= norm
	}
	return *mat.NewVecDense(v.Len(), data)
}

func gramSchmidt(A mat.Matrix) *mat.Dense {
	r, c := A.Dims()
	Q := mat.NewDense(r, c, nil)
	Q.Copy(A)
	V := getColumn(A, 0)
	normV := normalize(&V)
	Q.SetCol(0, normV.RawVector().Data)

	for j := 1; j < c; j++ {
		// Получаем текущий столбец из матрицы A
		col := mat.Col(nil, j, A)
		v := mat.NewVecDense(r, col)

		// Ортогонализация
		for i := 0; i < j; i++ {
			prev := getColumn(Q, i)
			v.SubVec(v, proj(v, &prev))
		}

		// Нормализация вектора
		norm := mat.Norm(v, 2)
		if norm == 0 {
			// Если нулевой вектор, пропускаем его
			continue
		}
		v.ScaleVec(1/norm, v)

		// Устанавливаем ортогонализированный вектор в матрицу Q
		Q.SetCol(j, v.RawVector().Data)
	}

	return Q
}

func proj(a, b *mat.VecDense) *mat.VecDense {
	// Вычисляем скалярное произведение векторов a и b
	dotProduct := mat.Dot(a, b)

	// Вычисляем квадрат нормы вектора b
	bNormSquared := mat.Dot(b, b)

	// Вычисляем проекцию вектора a на вектор b
	projection := mat.NewVecDense(b.Len(), nil)
	projection.ScaleVec(dotProduct/bNormSquared, b)

	return projection
}
