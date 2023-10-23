package slaeresolveriface

import "gonum.org/v1/gonum/mat"

/*
SLAE Resolver interface contains method for
resolving System of linear algebraic equation
in matrix form
(for example, Ax = b, where

	A - a square matrix with dims NxN, contains coeeficients of SLAE (a_ij).
	b - vector with dim N, contains left part of the SLAE (b_i).
	x - vector with dim N, contains root of the SLAE (x_i).
	).
*/
type SLAEResolver interface {
	Resolve(mat.Matrix, mat.Vector) (mat.Vector, error)
}
