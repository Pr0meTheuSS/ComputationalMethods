package integrator

import (
	"main/interval"
	"math"
)

type simpsonResolver struct{}

func (r simpsonResolver) Integrate(f func(float64) float64, interval interval.Interval, accuracy float64) (float64, error) {
	n := int64(interval.Right-interval.Left) * 100
	h := (interval.Right - interval.Left) / float64(n)

	res1 := intrgrateOnInterval(f, interval, n, h)
	res2 := intrgrateOnInterval(f, interval, n*2, h/2)

	for (1.0/15.0)*math.Abs(res2-res1) > accuracy {
		n *= 2
		h = (interval.Right - interval.Left) / float64(n)

		res1 = intrgrateOnInterval(f, interval, n, h)
		res2 = intrgrateOnInterval(f, interval, n*2, h/2)
	}

	return res2, nil
}

func intrgrateOnInterval(f func(float64) float64, interval interval.Interval, steps int64, stepSize float64) float64 {
	result := f(interval.Left) + f(interval.Right)

	// Сумма четных узлов
	evenSum := 0.0
	for i := int64(1); i < steps; i += 2 {
		evenSum += f(interval.Left + float64(i)*stepSize)
	}

	// Сумма нечетных узлов
	oddSum := 0.0
	for i := int64(2); i < steps-1; i += 2 {
		oddSum += f(interval.Left + float64(i)*stepSize)
	}

	result += 4*evenSum + 2*oddSum // Умножаем на соответствующие коэффициенты
	result *= stepSize / 3         // Умножаем на h/3

	return result
}
