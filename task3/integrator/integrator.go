package integrator

import "main/interval"

type Integrator interface {
	Integrate(f func(float64) float64, interval interval.Interval, accuracy float64) (float64, error)
}

func IntegratorFactory(methodName string) Integrator {
	switch methodName {
	case "trapezoid":
		return trapezoidResolver{}
	case "simpson":
		return simpsonResolver{}
	default:
		return trapezoidResolver{}
	}
}
