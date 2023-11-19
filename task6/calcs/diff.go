package calcs

// Решение дифференциального уравнения y'(x) = -y(x)
func differentialEquation(y float64) float64 {
	return -y
}

func FirstOrderDiff(x []float64, y0 float64) []float64 {
	n := len(x)

	h := (x[n-1] - x[0]) / float64(n)
	y := make([]float64, n)

	// Начальные условия
	y[0] = y0

	// Решение уравнения с разностной схемой
	for i := 0; i < n-1; i++ {
		y[i+1] = y[i] + h*differentialEquation(y[i])
	}

	return y
}

// Функция для разностной схемы второго порядка аппроксимации
func SecondOrderDiff(x []float64, y0 float64, f func(float64) float64) []float64 {
	n := len(x)
	h := (x[n-1] - x[0]) / float64(n-1)
	y := make([]float64, n)

	y[0] = y0

	for i := 0; i < n-1; i++ {
		y[i+1] = y[i] + h*f(x[i]) + (h*h/2)*f(x[i])
	}

	return y
}

// Функция для разностной схемы четвертого порядка аппроксимации
func FourthOrderDiff(x []float64, y0 float64, f func(float64) float64) []float64 {
	n := len(x)

	h := (x[n-1] - x[0]) / float64(n)
	y := make([]float64, n)

	y[0] = y0

	// Решение уравнения с разностной схемой
	for i := 0; i < n-1; i++ {
		k1 := h * differentialEquation(y[i])
		k2 := h * differentialEquation(y[i]+k1/2)
		k3 := h * differentialEquation(y[i]+k2/2)
		k4 := h * differentialEquation(y[i]+k3)

		y[i+1] = y[i] + (k1+2*k2+2*k3+k4)/6
	}

	return y
}

// Центральная разностная схема второго порядка (k=2)
func SecondOrderCentralDifference(x []float64, h float64, f func(float64) float64) []float64 {
	n := len(x)
	y := make([]float64, n)

	// Начальные условия
	y[0] = 0

	// Решение уравнения с разностной схемой
	for i := 0; i < n-1; i++ {
		y[i+1] = y[i] + h*f(y[i]) + (h*h/2)*f(y[i])
	}

	// Главный член невязки для второй схемы
	// residual := make([]float64, n)
	// for i := 1; i < n-1; i++ {
	// 	residual[i] = y[i] - math.Exp(x[i])*math.Cos(x[i])
	// }

	return y
}

// Центральная разностная схема четвертого порядка (k=4)
func FourthOrderCentralDifference(x []float64, y0 float64) []float64 {
	n := len(x)

	h := (x[n-1] - x[0]) / float64(n)
	y := make([]float64, n)

	y[0] = y0
	for i := 2; i < n-2; i++ {
		y[i] = (-x[i+2] + 8*x[i+1] - 8*x[i-1] + x[i-2]) / (12 * h)
	}

	// Главный член невязки для второй схемы
	// residual := make([]float64, n)
	// for i := 2; i < n-2; i++ {
	// 	residual[i] = y[i] - math.Exp(x[i])*math.Cos(x[i])
	// }

	return y
}
