package factory

import (
	jacobi "main/Jacobi"
	lu "main/LU"
	qr "main/QR"
	slaeresolveriface "main/SLAEResolverIface"
	"main/conjugate"
	"main/trig"
)

/*
Simple SLAE Resolver's Factory interface.
Method Create returns ready to work instance of
SLAEResolver interface implementation.
*/
type SLAEResolverFactory interface {
	Create(string) slaeresolveriface.SLAEResolver
}

type SLAEResolverFactoryImpl struct {
}

// TODO: Write doc. Add logs.
func (rf SLAEResolverFactoryImpl) Create(resolvingMethodName string) slaeresolveriface.SLAEResolver {
	switch resolvingMethodName {
	case "LU", "lu":
		return lu.NewLUResolver()
	case "Trig", "TRIG", "trig":
		return trig.NewTrigResolver()
	case "QR", "qr":
		return qr.NewQRResolver()
	case "J", "j", "Jacobi", "jacobi", "jac":
		return jacobi.NewJacobiResolver(10e3, 1e-9)
	case "con", "conjugate", "grad", "conjugate gradients":
		return conjugate.NewConjugateGradientSolver(1e3, 1e-6)
	default:
		return nil
	}
}
