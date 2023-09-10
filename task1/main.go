package main

import (
	"fmt"
	"os"
	"task1/equatation_resolver"
	"task1/polynomial"
)

func main() {
	// Запросить ввод выражения (полинома) от пользователя
	fmt.Print("Введите коэффициенты полинома от меньшей степени параметра к большей: \n")
	coefficients := make([]float64, 4)
	fmt.Scan(&coefficients[0], &coefficients[1], &coefficients[2], &coefficients[3])

	fmt.Println("Коэффициенты полинома:", coefficients)

	p := polynomial.Polynomial{
		Coefficients: coefficients,
	}

	epsilon := 1e-6
	resolver := &equatation_resolver.DichotomyMethod{
		Epsilon: epsilon,
	}

	roots, err := resolver.Resolve(p)
	if nil != err {
		fmt.Printf("При вычислении произошла ошибка: %+v\n", err)
		os.Exit(1)
	}

	fmt.Println("Корни уравнения: ", roots)

}
