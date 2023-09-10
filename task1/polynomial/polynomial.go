package polynomial

import "math"

type Polynomial struct {
	Coefficients []float64
}

func (p *Polynomial) CalcPolynomialDerivative() Polynomial {
	polynomialDerivateCoefficients := []float64{}
	for index, value := range p.Coefficients[1:] {
		polynomialDerivateCoefficients = append(polynomialDerivateCoefficients, value*float64(index+1))
	}

	return Polynomial{
		Coefficients: polynomialDerivateCoefficients,
	}
}

// Function calculate value for polynomial P(x) = an x^n + a1 x^(n-1) + ... + a0 for value point.
func (p *Polynomial) CalcValue(point float64) float64 {
	res := 0.0
	for i, v := range p.Coefficients {
		res += v * math.Pow(point, float64(i))
	}

	return res
}

func (p *Polynomial) GetPolynomialOrder() int {
	for i := len(p.Coefficients) - 1; i >= 0; i-- {
		if p.Coefficients[i] != 0.0 {
			return i
		}
	}

	return 0
}
