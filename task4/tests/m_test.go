package main_test

import (
	"main/common"
	"main/factory"
	"testing"

	"gonum.org/v1/gonum/mat"
)

var positiveTestCasesLU = []struct {
	A         *mat.Dense
	b         *mat.VecDense
	expectedX *mat.VecDense
}{
	{
		A:         mat.NewDense(2, 2, []float64{2, 1, 1, 3}),
		b:         mat.NewVecDense(2, []float64{5, 9}),
		expectedX: mat.NewVecDense(2, []float64{1.2, 2.6}),
	},
	{
		A:         mat.NewDense(3, 3, []float64{2, 1, -1, 1, 3, 2, 1, -1, 3}),
		b:         mat.NewVecDense(3, []float64{8, 11, 6}),
		expectedX: mat.NewVecDense(3, []float64{3.84, 1.56, 1.24}),
	},
	{
		A:         mat.NewDense(3, 3, []float64{4, -1, -1, -1, 4, -1, -1, -1, 4}),
		b:         mat.NewVecDense(3, []float64{2, 2, 2}),
		expectedX: mat.NewVecDense(3, []float64{1, 1, 1}),
	},
	{
		A:         mat.NewDense(3, 3, []float64{2, -1, 0, -1, 2, -1, 0, -1, 2}),
		b:         mat.NewVecDense(3, []float64{2, 2, 2}),
		expectedX: mat.NewVecDense(3, []float64{3, 4, 3}),
	},

	{
		A:         mat.NewDense(2, 2, []float64{2, -1, -1, 2}),
		b:         mat.NewVecDense(2, []float64{2, 2}),
		expectedX: mat.NewVecDense(2, []float64{2, 2}),
	},

	{
		A:         mat.NewDense(3, 3, []float64{5, 1, 0, 1, 5, 1, 0, 1, 5}),
		b:         mat.NewVecDense(3, []float64{1, 1, 1}),
		expectedX: mat.NewVecDense(3, []float64{4.0 / 23, 3.0 / 23, 4.0 / 23}),
	},
	{
		A:         mat.NewDense(3, 3, []float64{2, 1, -1, 1, -5, 4, 3, 2, 6}),
		b:         mat.NewVecDense(3, []float64{0, 10, 7}),
		expectedX: mat.NewVecDense(3, []float64{1, -1, 1}),
	},
	{
		A: mat.NewDense(3, 3, []float64{
			1, 1.0 / 2, 1.0 / 3,
			1.0 / 2, 1.0 / 3, 1.0 / 4,
			1.0 / 3, 1.0 / 4, 1.0 / 5}),

		b:         mat.NewVecDense(3, []float64{1, 1, 1}),
		expectedX: mat.NewVecDense(3, []float64{3, -24, 30}),
	},
}

// Матрицы с вырожденным ушловым минором.
var negativeTestCasesLU = []struct {
	A         *mat.Dense
	b         *mat.VecDense
	expectedX *mat.VecDense
}{
	{
		A:         mat.NewDense(3, 3, []float64{20, 20, 0, 15, 15, 5, 0, 1, 1}),
		b:         mat.NewVecDense(3, []float64{40, 35, 2}),
		expectedX: mat.NewVecDense(3, []float64{1, 1, 1}),
	},
}

func TestSolveLinearEquationsByLU(t *testing.T) {
	var f factory.SLAEResolverFactory = factory.SLAEResolverFactoryImpl{}
	resolver := f.Create("LU")

	for _, testCase := range positiveTestCasesLU {
		x, _ := resolver.Resolve(testCase.A, testCase.b)
		common.MatrixInfo(testCase.A)
		if !mat.EqualApprox(x, testCase.expectedX, 1e-6) {
			t.Errorf("Неверное решение СЛАУ. Ожидается %v, получено %v", testCase.expectedX, x)
		}
	}

	for _, testCase := range negativeTestCasesLU {
		_, err := resolver.Resolve(testCase.A, testCase.b)
		common.MatrixInfo(testCase.A)
		if err == nil {
			t.Error("Ожидалась ошибка вычисления методом LU, но err == nil.")
		}
	}

}

var positiveTestCasesQR = []struct {
	A         *mat.Dense
	b         *mat.VecDense
	expectedX *mat.VecDense
}{
	{
		A:         mat.NewDense(2, 2, []float64{2, 1, 1, 3}),
		b:         mat.NewVecDense(2, []float64{5, 9}),
		expectedX: mat.NewVecDense(2, []float64{1.2, 2.6}),
	},
	{
		A:         mat.NewDense(3, 3, []float64{2, 1, -1, 1, 3, 2, 1, -1, 3}),
		b:         mat.NewVecDense(3, []float64{8, 11, 6}),
		expectedX: mat.NewVecDense(3, []float64{3.84, 1.56, 1.24}),
	},
	{
		A:         mat.NewDense(3, 3, []float64{4, -1, -1, -1, 4, -1, -1, -1, 4}),
		b:         mat.NewVecDense(3, []float64{2, 2, 2}),
		expectedX: mat.NewVecDense(3, []float64{1, 1, 1}),
	},
	{
		A:         mat.NewDense(3, 3, []float64{2, -1, 0, -1, 2, -1, 0, -1, 2}),
		b:         mat.NewVecDense(3, []float64{2, 2, 2}),
		expectedX: mat.NewVecDense(3, []float64{3, 4, 3}),
	},

	{
		A:         mat.NewDense(2, 2, []float64{2, -1, -1, 2}),
		b:         mat.NewVecDense(2, []float64{2, 2}),
		expectedX: mat.NewVecDense(2, []float64{2, 2}),
	},

	{
		A:         mat.NewDense(3, 3, []float64{5, 1, 0, 1, 5, 1, 0, 1, 5}),
		b:         mat.NewVecDense(3, []float64{1, 1, 1}),
		expectedX: mat.NewVecDense(3, []float64{4.0 / 23, 3.0 / 23, 4.0 / 23}),
	},
	{
		A:         mat.NewDense(3, 3, []float64{2, 1, -1, 1, -5, 4, 3, 2, 6}),
		b:         mat.NewVecDense(3, []float64{0, 10, 7}),
		expectedX: mat.NewVecDense(3, []float64{1, -1, 1}),
	},
	{
		A:         mat.NewDense(3, 3, []float64{20, 20, 0, 15, 15, 5, 0, 1, 1}),
		b:         mat.NewVecDense(3, []float64{40, 35, 2}),
		expectedX: mat.NewVecDense(3, []float64{1, 1, 1}),
	},
	{
		A: mat.NewDense(3, 3, []float64{
			1, 1.0 / 2, 1.0 / 3,
			1.0 / 2, 1.0 / 3, 1.0 / 4,
			1.0 / 3, 1.0 / 4, 1.0 / 5}),

		b:         mat.NewVecDense(3, []float64{1, 1, 1}),
		expectedX: mat.NewVecDense(3, []float64{3, -24, 30}),
	},
	{
		A: mat.NewDense(3, 3, []float64{
			1, 1.001, 1.002,
			1.003, 1.001, 1.0,
			1.002, 1.003, 1.0}),

		b:         mat.NewVecDense(3, []float64{1, 1, 1}),
		expectedX: mat.NewVecDense(3, []float64{2000.0 / 6007, 1000.0 / 6007, 3000.0 / 6007}),
	},
	{
		A: mat.NewDense(3, 3, []float64{
			0, 1, 2,
			1, 0, 1,
			2, 1, 0}),

		b:         mat.NewVecDense(3, []float64{1, 2, 2}),
		expectedX: mat.NewVecDense(3, []float64{1, 0, 1}),
	},

}

func TestSolveLinearEquationsByQR(t *testing.T) {
	for _, testCase := range positiveTestCasesQR {
		var f factory.SLAEResolverFactory = factory.SLAEResolverFactoryImpl{}
		resolver := f.Create("QR")
		x, _ := resolver.Resolve(testCase.A, testCase.b)
		common.MatrixInfo(testCase.A)
		if !mat.EqualApprox(x, testCase.expectedX, 1e-6) {
			t.Errorf("Неверное решение СЛАУ. Ожидается %v, получено %v", testCase.expectedX, x)
		}
	}
}

var positiveTestCasesJacobi = []struct {
	A         *mat.Dense
	b         *mat.VecDense
	expectedX *mat.VecDense
}{
	{
		A:         mat.NewDense(2, 2, []float64{2, 1, 1, 3}),
		b:         mat.NewVecDense(2, []float64{5, 9}),
		expectedX: mat.NewVecDense(2, []float64{1.2, 2.6}),
	},
	{
		A:         mat.NewDense(3, 3, []float64{2, 1, -1, 1, 3, 2, 1, -1, 3}),
		b:         mat.NewVecDense(3, []float64{8, 11, 6}),
		expectedX: mat.NewVecDense(3, []float64{3.84, 1.56, 1.24}),
	},
	{
		A:         mat.NewDense(3, 3, []float64{4, -1, -1, -1, 4, -1, -1, -1, 4}),
		b:         mat.NewVecDense(3, []float64{2, 2, 2}),
		expectedX: mat.NewVecDense(3, []float64{1, 1, 1}),
	},
	{
		A:         mat.NewDense(3, 3, []float64{2, -1, 0, -1, 2, -1, 0, -1, 2}),
		b:         mat.NewVecDense(3, []float64{2, 2, 2}),
		expectedX: mat.NewVecDense(3, []float64{3, 4, 3}),
	},
	{
		A:         mat.NewDense(2, 2, []float64{2, -1, -1, 2}),
		b:         mat.NewVecDense(2, []float64{2, 2}),
		expectedX: mat.NewVecDense(2, []float64{2, 2}),
	},

	{
		A:         mat.NewDense(3, 3, []float64{5, 1, 0, 1, 5, 1, 0, 1, 5}),
		b:         mat.NewVecDense(3, []float64{1, 1, 1}),
		expectedX: mat.NewVecDense(3, []float64{4.0 / 23, 3.0 / 23, 4.0 / 23}),
	},
	{
		A:         mat.NewDense(3, 3, []float64{2, 1, -1, 1, -5, 4, 3, 2, 6}),
		b:         mat.NewVecDense(3, []float64{0, 10, 7}),
		expectedX: mat.NewVecDense(3, []float64{1, -1, 1}),
	},
	{
		A: mat.NewDense(3, 3, []float64{
			1, 1.0 / 2, 1.0 / 3,
			1.0 / 2, 1.0 / 3, 1.0 / 4,
			1.0 / 3, 1.0 / 4, 1.0 / 5}),

		b:         mat.NewVecDense(3, []float64{1, 1, 1}),
		expectedX: mat.NewVecDense(3, []float64{3, -24, 30}),
	},

	// Выкинул плохо обусловленную матрицу
}

func TestSolveLinearEquationsByJacobi(t *testing.T) {
	for _, testCase := range positiveTestCasesJacobi {
		var f factory.SLAEResolverFactory = factory.SLAEResolverFactoryImpl{}
		resolver := f.Create("J")
		x, _ := resolver.Resolve(testCase.A, testCase.b)
		common.MatrixInfo(testCase.A)
		if !mat.EqualApprox(x, testCase.expectedX, 1e-6) {
			t.Errorf("Неверное решение СЛАУ. Ожидается %v, получено %v", testCase.expectedX, x)
		}
	}
}

var positiveTestCasesTrig = []struct {
	A         *mat.Dense
	b         *mat.VecDense
	expectedX *mat.VecDense
}{
	{
		A:         mat.NewDense(2, 2, []float64{2, 1, 1, 3}),
		b:         mat.NewVecDense(2, []float64{5, 9}),
		expectedX: mat.NewVecDense(2, []float64{1.2, 2.6}),
	},
	{
		A:         mat.NewDense(3, 3, []float64{2, -1, 0, -1, 2, -1, 0, -1, 2}),
		b:         mat.NewVecDense(3, []float64{2, 2, 2}),
		expectedX: mat.NewVecDense(3, []float64{3, 4, 3}),
	},

	{
		A:         mat.NewDense(2, 2, []float64{2, -1, -1, 2}),
		b:         mat.NewVecDense(2, []float64{2, 2}),
		expectedX: mat.NewVecDense(2, []float64{2, 2}),
	},

	{
		A:         mat.NewDense(3, 3, []float64{5, 1, 0, 1, 5, 1, 0, 1, 5}),
		b:         mat.NewVecDense(3, []float64{1, 1, 1}),
		expectedX: mat.NewVecDense(3, []float64{4.0 / 23, 3.0 / 23, 4.0 / 23}),
	},
	{
		A: mat.NewDense(3, 3, []float64{
			1, 1.0 / 2, 1.0 / 3,
			1.0 / 2, 1.0 / 3, 1.0 / 4,
			1.0 / 3, 1.0 / 4, 1.0 / 5}),

		b:         mat.NewVecDense(3, []float64{1, 1, 1}),
		expectedX: mat.NewVecDense(3, []float64{3, -24, 30}),
	},
}

func TestSolveLinearEquationsByTrig(t *testing.T) {
	for _, testCase := range positiveTestCasesTrig {
		var f factory.SLAEResolverFactory = factory.SLAEResolverFactoryImpl{}
		resolver := f.Create("trig")
		x, _ := resolver.Resolve(testCase.A, testCase.b)
		common.MatrixInfo(testCase.A)
		if !mat.EqualApprox(x, testCase.expectedX, 1e-6) {
			t.Errorf("Неверное решение СЛАУ. Ожидается %v, получено %v", testCase.expectedX, x)
		}
	}
}

var positiveTestCasesCG = []struct {
	A         *mat.Dense
	b         *mat.VecDense
	expectedX *mat.VecDense
}{
	{
		A:         mat.NewDense(3, 3, []float64{2, -1, 0, -1, 2, -1, 0, -1, 2}),
		b:         mat.NewVecDense(3, []float64{2, 2, 2}),
		expectedX: mat.NewVecDense(3, []float64{3, 4, 3}),
	},

	{
		A:         mat.NewDense(2, 2, []float64{2, -1, -1, 2}),
		b:         mat.NewVecDense(2, []float64{2, 2}),
		expectedX: mat.NewVecDense(2, []float64{2, 2}),
	},
}

func TestSolveLinearEquationsByConjugateGradients(t *testing.T) {
	for _, testCase := range positiveTestCasesCG {
		var f factory.SLAEResolverFactory = factory.SLAEResolverFactoryImpl{}
		resolver := f.Create("grad")
		x, _ := resolver.Resolve(testCase.A, testCase.b)
		common.MatrixInfo(testCase.A)
		if !mat.EqualApprox(x, testCase.expectedX, 1e-3) {
			t.Errorf("Неверное решение СЛАУ. Ожидается %v, получено %v", testCase.expectedX, x)
		}
	}
}
