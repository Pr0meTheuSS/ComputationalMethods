package main

import (
	"fmt"
	"log"
	"main/integrator"
	"main/interval"
	"math"
)

/* TODO:
 *
- write tests.
- add function expression parser.
*
*/

func main() {
	// var functionExpr string
	// if _, err := fmt.Scan(&functionExpr); err != nil {
	// 	log.Fatal("Wrong input. Expected expression with function for integraton.\n For example: sin(x), ln(x + 12x^2)\n")
	// }

	function := func(x float64) float64 { return math.Exp(x) * math.Cos(x) }
	// function, err := functionparser.ParseFunction(functionExpr)
	// if err != nil {
	// 	log.Fatalf("Function expression parser failed with error: %s\n", err.Error())
	// }

	fmt.Println("Enter interval for integration (two numbers - left and right bounds):")
	interval := interval.Interval{}
	if _, err := fmt.Scan(&interval.Left, &interval.Right); err != nil {
		log.Fatal("Wrong input. Expected expression with two float values.\n For example: -3.12 10\n")
	}

	fmt.Println("Enter accuracy:")
	var accuracy float64
	if _, err := fmt.Scan(&accuracy); err != nil {
		log.Fatal("Wrong input. Expected expression with one float value.\n For example: 0.001\n")
	}

	resolver := integrator.IntegratorFactory("simpson")

	result, err := resolver.Integrate(function, interval, accuracy)
	if err != nil {
		log.Fatalf("Integrator failed with error: %s\n", err.Error())
	}

	fmt.Println("Integral value:", result)
}
