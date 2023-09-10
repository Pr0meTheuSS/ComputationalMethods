package equatation_resolver

import (
	"errors"
	"log"
	"math"
	"sort"
	"task1/polynomial"
)

type Interval struct {
	left  float64
	right float64
}

type EquatationResolverInterface interface {
	// Method resolve polynomial 3rd order-equatation.
	resolve(p polynomial.Polynomial) ([]float64, error)
}

type DichotomyMethod struct {
	epsilon float64
}

func (d *DichotomyMethod) resolve(polynomial polynomial.Polynomial) ([]float64, error) {
	log.Printf("eq order: %d\n", polynomial.GetPolynomialOrder())
	switch polynomial.GetPolynomialOrder() {
	case 1:
		return []float64{-polynomial.Coefficients[0] / polynomial.Coefficients[1]}, nil
	case 2:
		return calcEquatationRoots(polynomial)
	case 3:
		return d.resolveCubicEquatation(polynomial)
	default:
		return nil, errors.New("Cannot resolve this equatation. Get a wrong equatation order.")
	}
}

func (d *DichotomyMethod) resolveCubicEquatation(polynomial polynomial.Polynomial) ([]float64, error) {
	referencePoints, err := d.getReferencePoint(polynomial)
	if nil != err {
		return nil, err
	}

	intervals := d.getSearchIntervals(referencePoints)
	roots := []float64{}

	for _, interval := range intervals {
		currentRoots := d.binarySearch(polynomial, interval)
		if nil != currentRoots {
			roots = append(roots, currentRoots...)
		}
	}

	roots = collapseCloseRoots(roots, d.epsilon)
	return roots, nil

}

func collapseCloseRoots(roots []float64, distance float64) []float64 {
	if len(roots) == 0 {
		return roots
	}

	collapsedRoots := []float64{roots[0]}

	for i := 0; i < len(roots)-1; i++ {
		if math.Abs(roots[i]-roots[i+1]) <= distance {
			continue
		}

		collapsedRoots = append(collapsedRoots, roots[i+1])
	}

	return collapsedRoots
}

func calcEquatationRoots(polynomial polynomial.Polynomial) ([]float64, error) {
	if polynomial.GetPolynomialOrder() != 2 {
		return nil, errors.New("Cannot resolve equation more than second order.")
	}

	a := polynomial.Coefficients[2]
	b := polynomial.Coefficients[1]
	c := polynomial.Coefficients[0]

	D := b*b - 4*a*c

	if D > 0 {
		return []float64{(-b - math.Sqrt(D)) / (2 * a), (-b + math.Sqrt(D)) / (2 * a)}, nil
	} else if D == 0 {
		return []float64{-b / (2 * a)}, nil
	} else {
		return []float64{}, nil
	}
}

func (d *DichotomyMethod) getReferencePoint(polynomial polynomial.Polynomial) ([]float64, error) {
	derivate := polynomial.CalcPolynomialDerivative()

	referencePoints, err := calcEquatationRoots(derivate)
	if nil != err {
		return nil, err
	}

	return referencePoints, err
}

// Function search root in interval (using binary search)
func (d *DichotomyMethod) binarySearch(polynomial polynomial.Polynomial, interval Interval) []float64 {
	roots := []float64{}

	for math.Abs(interval.right-interval.left) >= 10e-12 {
		mid := (interval.left + interval.right) / 2

		if polynomial.CalcValue(mid)*polynomial.CalcValue(interval.left) <= 0.0 {
			interval.right = mid
		} else {
			interval.left = mid
		}
	}

	if math.Abs(polynomial.CalcValue(interval.left)) <= d.epsilon {
		roots = append(roots, interval.left)
	}

	return roots
}

func (d *DichotomyMethod) getSearchIntervals(referencePoints []float64) []Interval {
	sort.Float64s(referencePoints)

	intervals := []Interval{}
	if len(referencePoints) == 0 {
		return []Interval{
			{left: -1e10, right: 1e10},
		}
	}

	// handle first and last points.
	intervals = append(intervals, Interval{left: -1e10, right: referencePoints[0]})
	intervals = append(intervals, Interval{left: referencePoints[len(referencePoints)-1], right: 1e10})

	for i, point := range referencePoints[:len(referencePoints)-1] {
		intervals = append(intervals, Interval{left: point, right: referencePoints[i+1]})
	}

	return intervals
}
