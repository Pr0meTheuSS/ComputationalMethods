package integrator

import (
	"main/interval"
	"math"
)

type trapezoidResolver struct{}

func (r trapezoidResolver) Integrate(f func(float64) float64, interval interval.Interval, accuracy float64) (float64, error) {
	n := int64(interval.Right-interval.Left+1) * 100
	h := (interval.Right - interval.Left) / float64(n)

	res1 := r.intrgrateOnInterval(f, interval, n, h)
	res2 := r.intrgrateOnInterval(f, interval, n*2, h/2)

	for (1.0/3.0)*math.Abs(res1-res2) > accuracy {
		n *= 2
		h := (interval.Right - interval.Left) / float64(n)

		res1 = r.intrgrateOnInterval(f, interval, n, h)
		res2 = r.intrgrateOnInterval(f, interval, n*2, h/2)
	}

	return res2, nil
}

func (r trapezoidResolver) intrgrateOnInterval(f func(float64) float64, interval interval.Interval, steps int64, stepSize float64) float64 {
	result := 0.0
	left := interval.Left
	for i := 0; i < int(steps); i++ {
		right := left + stepSize
		result += f(left) + f(right)
		left = right
	}

	return result * stepSize / 2
}
