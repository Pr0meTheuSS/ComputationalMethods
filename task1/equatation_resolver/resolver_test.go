package equatation_resolver

import (
	"sort"
	"task1/polynomial"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type resolveEquationTest struct {
	coeffs   []float64
	expected []float64
}

func TestSolveCubic(t *testing.T) {
	epsilon := 1e-6
	resolver := &DichotomyMethod{
		Epsilon: epsilon,
	}

	testCases := []struct {
		coefficients []float64
		solutions    []float64
	}{
		{[]float64{-2.0, -1.0, 2.0, 1.0}, []float64{-2.0, -1.0, 1.0}},
		{[]float64{-100.0, 1.0, -100.0, 1.0}, []float64{100.0}},
		{[]float64{4.0, 3.0, 2.0, 1.0}, []float64{-1.6506291}},

		{[]float64{-6.0, 11.0, -6.0, 1.0}, []float64{1.0, 2.0, 3.0}},
		{[]float64{-6.0, -7.0, 0.0, 1.0}, []float64{-2.0, -1.0, 3.0}},
		{[]float64{0.0, 0.0, 0.0, 3.0}, []float64{0.0}},
		{[]float64{0.0, 1.0, 2.0, 3.0}, []float64{0.0}},
		{[]float64{1.0, 1.0, 1.0, 1.0}, []float64{-1.0}},
		{[]float64{0.0, 12.0, 0.0, -10.0}, []float64{-1.095445, 0.0, 1.095445}},

		{[]float64{1.0, 2.0, 1.0, 0.0}, []float64{-1.0}},
		{[]float64{1.0, 2.0, 0.0, 0.0}, []float64{-.5}},
		{[]float64{1.0, 2.0}, []float64{-.5}},
	}

	for _, tc := range testCases {
		p := polynomial.Polynomial{
			Coefficients: tc.coefficients,
		}

		t.Run("Positive tests cubic resolver", func(t *testing.T) {
			solutions, _ := resolver.Resolve(p)
			sort.Float64s(solutions)
			sort.Float64s(tc.solutions)

			// Проверяем, что количество корней совпадает с ожидаемым количеством
			require.Equal(t, len(tc.solutions), len(solutions))

			// Проверяем, что каждый корень близок к ожидаемому с точностью до небольшой погрешности
			for i, expected := range tc.solutions {
				actual := solutions[i]
				assert.InDelta(t, expected, actual, epsilon)
			}
		})
	}
}
