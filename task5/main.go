package main

import (
	"fmt"
	"math"
)

const epsilon = 1e-6
const maxIterations = 1000

func f1(x, a float64) float64 {
	return x - a - math.Log(x+1)
}

// Производная
func derivative1(x, a float64) float64 {
	return 1 / (1 + x)
}

// x-a = ln(x+1) -> e^(x-a) = x+1 -> x = e^(x-a)-1
func rev_f1(x, a float64) float64 {
	return math.Exp(x-a) - 1 - x
}

func rev_f1_derivate(x, a float64) float64 {
	return math.Exp(x-a) - 1
}

// Метод простой итерации
func NewtonMethod(a float64, f func(x, a float64) float64, d func(x, a float64) float64) (float64, error) {
	x := 0.0
	lambda0 := 1.0 / d(x, a)

	for i := 0; i < maxIterations; i++ {
		xNext := x - lambda0*f(x, a)
		fmt.Println(xNext)
		if math.Abs(xNext-x) < epsilon {
			return xNext, nil
		}
		x = xNext
	}

	return 0, fmt.Errorf("не удалось найти решение после %d итераций", maxIterations)
}

func secantMethod(a float64, f func(x, a float64) float64, x0, x1 float64) (float64, error) {
	var x2, fx0, fx1, fx2 float64

	for i := 0; i < maxIterations; i++ {
		fx0 = f(x0, a)
		fx1 = f(x1, a)

		if math.Abs(fx1-fx0) < epsilon {
			return x1, nil // Успех: достигнута достаточная точность
		}

		// Метод секущих: x2 = x1 - f(x1) * (x1 - x0) / (f(x1) - f(x0))
		x2 = x1 - (fx1 * (x1 - x0) / (fx1 - fx0))
		fx2 = f(x2, a)

		if math.Abs(fx2) < epsilon {
			return x2, nil // Успех: найден корень с достаточной точностью
		}
		fmt.Println(x2)
		// Переносим точки для следующей итерации
		x0, x1 = x1, x2
	}

	return 0, fmt.Errorf("Не удалось найти корень после %d итераций", maxIterations)
}

func main() {
	var a float64
	fmt.Print("Введите значение параметра a: ")
	fmt.Scan(&a)

	// Метод простой итерации
	resultSimple, errSimple := NewtonMethod(a, f1, derivative1)
	if errSimple != nil {
		fmt.Println(errSimple)
	} else {
		fmt.Printf("Метод простой итерации: x = %.6f\n", resultSimple)
	}

	resultSimple, errSimple = NewtonMethod(a, rev_f1, rev_f1_derivate)
	if errSimple != nil {
		fmt.Println(errSimple)
	} else {
		fmt.Printf("Метод простой итерации: x = %.6f\n", resultSimple)
	}

	// Метод секущих
	resultSecant, errSecant := secantMethod(a, f1, 0, a*a)
	if errSecant != nil {
		fmt.Println(errSecant)
	} else {
		fmt.Printf("Метод секущих: x = %.6f\n", resultSecant)
	}

	resultSecant, errSecant = secantMethod(a, rev_f1, -a*a, 0)
	if errSecant != nil {
		fmt.Println(errSecant)
	} else {
		fmt.Printf("Метод секущих: x = %.6f\n", resultSecant)
	}
}
