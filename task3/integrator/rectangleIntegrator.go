package integrator

import "main/interval"

type rectangleResolver struct{}

func (r rectangleResolver) Integrate(f func(float64) float64, interval interval.Interval, accuracy float64) (float64, error) {
	rectanglesAmount := int64(interval.Right-interval.Left+1) * 100
	h := (interval.Right - interval.Left) / float64(rectanglesAmount)
	result := 0.0

	left := interval.Left
	for i := 0; i < int(rectanglesAmount); i++ {
		right := left + h
		result += f((left + right) / 2)
		left = right
	}

	return result * h, nil
}
