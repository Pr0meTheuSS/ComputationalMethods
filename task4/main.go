package main

import (
	"fmt"
	slaeresolveriface "main/SLAEResolverIface"
	"main/common"
	"main/factory"
	"os"
)

func main() {
	fm, err := os.Open("matrix.data")
	if err != nil {
		return
	}
	defer fm.Close()

	A := common.ReadMatrix(fm)
	common.MatrixInfo(A)

	fv, err := os.Open("vector.dat")
	if err != nil {
		return
	}
	defer fm.Close()
	b := common.ReadVector(fv)

	var factory factory.SLAEResolverFactory = factory.SLAEResolverFactoryImpl{}
	var r slaeresolveriface.SLAEResolver = factory.Create("lu")

	res, err := r.Resolve(A, b)
	if err != nil {
		fmt.Printf("Error while resolving: %s\n", err.Error())
		return
	}

	fmt.Println("matrix: ")
	common.PrintMatrix(A)
	fmt.Println("vector: ")
	common.PrintVector(b)

	fmt.Println("result: ")
	common.PrintVector(res)
}
