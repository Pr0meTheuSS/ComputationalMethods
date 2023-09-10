package polynomial

import (
	"testing"
)

func TestCalcPolynomialDerivative(t *testing.T) {
	type derivateTest struct {
		polynomialCoefficients          []float64
		derivatedPolynomialCoefficients []float64
	}

	derivateTests := []derivateTest{
		// (3 + 2x + x^2)' -> 2 + 2x
		{
			[]float64{3.0, 2.0, 1.0},
			[]float64{2.0, 2.0},
		},
		// (1 - 2x + 3x^2 + 4x^3)' -> -2 + 6x + 12x^2
		{
			[]float64{1.0, -2.0, 3.0, 4.0},
			[]float64{-2.0, 6.0, 12.0},
		},
	}

	for _, test := range derivateTests {
		inputPolynomial := Polynomial{
			Coefficients: test.polynomialCoefficients,
		}
		expectedPolynomial := Polynomial{
			Coefficients: test.derivatedPolynomialCoefficients,
		}

		result := inputPolynomial.CalcPolynomialDerivative()

		if result.GetPolynomialOrder() != expectedPolynomial.GetPolynomialOrder() {
			t.Errorf("Expected %d coefficients, but got %d", expectedPolynomial.GetPolynomialOrder(), result.GetPolynomialOrder())
		}

		for i, coeff := range result.Coefficients {
			if coeff != expectedPolynomial.Coefficients[i] {
				t.Errorf("Coefficient %d is expected to be %f, but got %f", i, expectedPolynomial.Coefficients[i], coeff)
			}
		}
	}
}

func TestCalcValue(t *testing.T) {
	p := Polynomial{
		Coefficients: []float64{3.0, 2.0, 1.0},
	}

	point := 2.0
	result := p.CalcValue(point)

	expectedValue := 11.0
	if result != expectedValue {
		t.Errorf("Expected value %f, but got %f", expectedValue, result)
	}
}

func TestGetPolynomialOrder(t *testing.T) {
	p := Polynomial{
		Coefficients: []float64{3.0, 2.0, 1.0},
	}

	order := p.GetPolynomialOrder()

	expectedOrder := 2
	if order != expectedOrder {
		t.Errorf("Expected polynomial order %d, but got %d", expectedOrder, order)
	}
}
