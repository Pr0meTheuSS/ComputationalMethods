package integrator

import "main/interval"

type Integrator interface {
	Integrate(f func(float64) float64, interval interval.Interval, accuracy float64) (float64, error)
}

// TODO: separate and implement
type rectangleResolver struct{}

func (r rectangleResolver) Integrate(f func(float64) float64, interval interval.Interval, accuracy float64) (float64, error) {
	panic("implement me")
	// return 0.0, nil
}

func IntegratorFactory(methodName string) Integrator {
	panic("implement me")
	// switch methodName {
	// case "rectangle":
	// 	return rectangleResolver{}
	// default:
	// 	return rectangleResolver{}
	// }
}
