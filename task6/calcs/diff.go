package calcs

import "math"

func solution(x float64) float64 {
	return 1.0 / math.Exp(x)
}

func solution1(x float64) float64 {
	return math.Exp(x)*0.5*(math.Sin(x)+math.Cos(x)) - 0.5
}

func FirstOrderDiff(x []float64, y0, h float64, g func(x, y float64) float64) []float64 {
	/*
		y(x+h) - y(x)
		-------------  = g(x)
			  h

		y[i+1] = g(x) *h + y[i]
	*/
	n := len(x)
	y := make([]float64, n)

	// Начальные условия
	y[0] = y0

	// Решение уравнения с разностной схемой
	for i := 0; i < n-1; i++ {
		y[i+1] = y[i] + h*g(x[i], y[i])
	}

	return y
}

// Функция для разностной схемы второго порядка аппроксимации
func SecondOrderDiff(x []float64, y0, h float64, g func(x, y float64) float64) []float64 {
	/*

		y(x+h) - y(x-h)     g(x+h) + g(x-h)         |
		-------------   -   ----------------  = 0   |  * 2h
			  2h                   2                |


		y(x+h) = h(g(x+h) + g(x-h)) + y(x-h)


		y[i+1] = h * (g(x[i+1]) + g(x[i-1])) + y[i-1]



		y(x+h) - y(x-h)    				 |
		-------------   -  g(x)  =   0   |  * 2h
			  2h                         |

			  y[i+1] = y[i-1] +2*h*g(x)


	*/
	n := len(x)
	y := make([]float64, n)

	y[0] = y0
	y[1] = y[0] + h*g(x[0], y[0])

	for i := 1; i < n-1; i++ {
		y[i+1] = (1.0-h)*y[i] + (h/2.0)*(y[i-1]-y[i])
	}

	return y
}

// Функция для разностной схемы второго порядка аппроксимации c g(x)
func SecondOrderDiffCos(x []float64, y0, h float64, g func(x, y float64) float64) []float64 {
	n := len(x)
	y := make([]float64, n)

	y[0] = y0
	y[1] = y[0] + h*g(x[0], y[0])

	for i := 1; i < n-1; i++ {
		y[i+1] = h*(g(x[i+1], y[i+1])+g(x[i-1], y[i-1])) + y[i-1]
	}

	return y
}

// Функция для разностной схемы четвертого порядка аппроксимации
func FourthOrderDiff(x []float64, y0, h float64) []float64 {
	n := len(x)
	y := make([]float64, n)

	y[0] = y0
	y[1] = solution(x[1])

	for i := 1; i < n-1; i++ {
		y[i+1] = (y[i-1]*(3-h) - 4*h*y[i]) / (3.0 + h)
	}

	return y
}

// Функция для разностной схемы четвертого порядка аппроксимации
func FourthOrderDiffWithG(x []float64, y0, h float64, g func(x, y float64) float64) []float64 {
	n := len(x)
	y := make([]float64, n)

	y[0] = y0
	y[1] = y[0] + h*g(x[0], y[0])

	for i := 1; i < n-1; i++ {
		y[i+1] = 2*h*(g(x[i-1], y[i-1])/6.+2*g(x[i], y[i])/3+g(x[i+1], y[i+1])/6) + y[i-1]
	}

	return y
}
