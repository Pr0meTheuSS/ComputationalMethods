package main

import (
	"errors"
	"math/big"
)

type Equatation struct {
	coefficients []big.Float
	roots        []big.Float
}

type Interval struct {
}

type EquatationResolverInterface interface {
	resolve(e *Equatation) error
}

type DichotomyMethod struct {
	epsilon big.Float
}

func (d *DichotomyMethod) resolve(e *Equatation) error {
	if nil == e {
		return errors.New("Catch empty ecuatation struct.")
	}

	referencePoints, err := d.getReferencePoint(e.coefficients)
	if nil != err {
		return err
	}

	intervals, err := d.getSearchIntervals(referencePoints)
	if nil != err {
		return err
	}

	for _, interval := range intervals {
		currentRoot := d.binaryRootSearch(interval)
		if nil != currentRoot {
			e.coefficients = append(e.coefficients, currentRoot...)
		}
	}

	return nil
}

func (d *DichotomyMethod) getReferencePoint(coefficients []big.Float) ([]big.Float, error) {
	panic("implement me")
	return make([]big.Float, 0), nil
}

func (d *DichotomyMethod) binaryRootSearch(intervalForSearch Interval) []big.Float {
	panic("implement me")
	return make([]big.Float, 0)
}

func (d *DichotomyMethod) getSearchIntervals(coefficients []big.Float) ([]Interval, error) {
	panic("implement me")
	return make([]Interval, 0), nil
}
