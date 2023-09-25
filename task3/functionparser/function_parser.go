package functionparser

func ParseFunction(expression string) (func(float64) float64, error) {
	return func(x float64) float64 { return x }, nil
}
