package main

import (
	"main/calcs"
	"math"
	"testing"
)

func solution(x float64) float64 {
	return 1.0 / math.Exp(x)
}

const (
	eps = 1e-4
)

func TestCalcDiffFirstOrder(t *testing.T) {
	step := 0.0001
	left := 0.0
	right := 5.0
	n := (int64)((right - left) / step)

	x := make([]float64, n)
	for i := int64(0); i < n; i++ {
		x[i] = left + float64(i)*step
	}

	res := calcs.FirstOrderDiff(x, 1.0, step, func(x, y float64) float64 { return -y })
	for i, r := range res {
		if math.Abs(solution(x[i])-r) > eps {
			t.Fatalf("Неверное решение дифф. уравнения. Ожидается %v, получено %v", solution(x[i]), r)
		}
	}
}

func TestCalcDiffSecondOrder(t *testing.T) {
	step := 0.01
	left := 0.0
	right := 5.0
	n := (int64)((right - left) / step)

	x := make([]float64, n)
	for i := int64(0); i < n; i++ {
		x[i] = left + float64(i)*step
	}

	res := calcs.SecondOrderDiff(x, 1.0, step, func(x, y float64) float64 { return -y })
	for i, r := range res {
		if math.Abs(solution(x[i])-r) > eps {
			t.Fatalf("Неверное решение дифф. уравнения. Ожидается %v, получено %v. Итерация: %d", solution(x[i]), r, i)
		}
	}
}

func TestCalcDiffFourthOrder(t *testing.T) {
	step := 0.1
	left := 0.0
	right := 5.0
	n := (int64)((right - left) / step)

	x := make([]float64, n)
	for i := int64(0); i < n; i++ {
		x[i] = left + float64(i)*step
	}

	res := calcs.FourthOrderDiff(x, 1.0, step)
	for i, r := range res {
		if math.Abs(solution(x[i])-r) > eps {
			t.Fatalf("Неверное решение дифф. уравнения. Ожидается %v, получено %v. Итерация: %d", solution(x[i]), r, i)
		}
	}
}

func solution1(x float64) float64 {
	return math.Exp(x)*0.5*(math.Sin(x)+math.Cos(x)) - 0.5
}

func TestCalcSecondOrderCosExp(t *testing.T) {
	step := 0.1
	left := 0.0
	right := 2.5
	n := (int64)((right - left) / step)

	x := make([]float64, n)
	for i := int64(0); i < n; i++ {
		x[i] = left + float64(i)*step
	}

	res := calcs.SecondOrderDiffCos(x, .0, step, func(x, y float64) float64 { return math.Exp(x) * math.Cos(x) })

	for i, r := range res {
		if math.Abs(solution1(x[i])-r) > eps {
			t.Fatalf("Неверное решение дифф. уравнения в x_%d. Ожидается %v, получено %v", i, solution1(x[i]), r)
		}
	}
}

func TestCalcFourthOrderCosExp(t *testing.T) {
	step := 0.1
	left := 0.0
	right := 2.5
	n := (int64)((right - left) / step)

	x := make([]float64, n)
	for i := int64(0); i < n; i++ {
		x[i] = left + float64(i)*step
	}

	res := calcs.FourthOrderDiffWithG(x, 0.0, step, func(x, y float64) float64 { return math.Exp(x) * math.Cos(x) })

	for i, r := range res {
		if math.Abs(solution1(x[i])-r) > eps {
			t.Fatalf("Неверное решение дифф. уравнения в x_%d. Ожидается %v, получено %v", i, solution1(x[i]), r)
		}
	}
}
