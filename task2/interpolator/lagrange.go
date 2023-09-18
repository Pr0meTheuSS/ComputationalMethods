package interpolator

// Node представляет узел (точку) с координатами x и y.
type Node struct {
	X float64
	Y float64
}

type InterpolatorInterface interface {
	CalcPolynomialValue(x float64, nodes []Node) float64
}

type LagrangeInterpolator struct{}

// Вычисляет значения полинома Лагранжа для заданного значения и заданных узлов.
func (li LagrangeInterpolator) CalcPolynomialValue(x float64, nodes []Node) float64 {
	n := len(nodes)
	result := 0.0

	for i := 0; i < n; i++ {
		term := nodes[i].Y
		for j := 0; j < n; j++ {
			if i != j {
				term *= (x - nodes[j].X) / (nodes[i].X - nodes[j].X)
			}
		}
		result += term
	}

	return result
}
