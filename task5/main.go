package main

import (
	"fmt"
	"math"
)

const epsilon = 1e-6
const maxIterations = 1000

func f1(x, a float64) float64 {
	return x - a - math.Log(x+1)
}

// Производная
func derivative1(x, a float64) float64 {
	return 1 / (1 + x)
}

// x-a = ln(x+1) -> e^(x-a) = x+1 -> x = e^(x-a)-1
func rev_f1(x, a float64) float64 {
	return math.Exp(x-a) - 1 - x
}

func rev_f1_derivate(x, a float64) float64 {
	return math.Exp(x-a) - 1
}

// Метод простой итерации
func NewtonMethod(a float64, f func(x, a float64) float64, d func(x, a float64) float64, x0 float64) (float64, error) {
	x := x0
	lambda0 := 1.0 / d(x, a)
	fmt.Println("lambda:", lambda0)
	for i := 0; i < maxIterations; i++ {
		// Если решение достаточно близко:
		if math.Abs(f(x, a)) <= epsilon {
			return x, nil
		}

		xNext := x - lambda0*f(x, a)
		// fmt.Println(x, xNext)
		if math.Abs(xNext-x) < epsilon {
			return xNext, nil
		}

		x = xNext
	}

	return 0, fmt.Errorf("не удалось найти решение после %d итераций. Текущий ответ: %f", maxIterations, x)
}

// Генерируем подотрезки вида [-Pi/2 + Pi*k, Pi/2 + Pi*k]
func generateSubIntervalsAboutZero(x float64) []part {
	// Известно, что синус ведет себя монотонно на участках [-Pi/2 + Pi*k; Pi/2 + Pi*k]
	// Считаем количество входящих участков монотонности.
	/*

												^ Y
												|
												|       ***
												|     *
												|   *
												| *
		    			      -Pi      -Pi/2	|*      Pi/2
				-------------- *--------|-------*--------------------------------> X
		    				    *		|      *|
								 * 		|     * |
								   *    |    *  |
									 *  |  *    |
									   ***      |
												|

	*/

	var subIntervals []part
	left := -x
	right := x
	step := math.Pi

	k_min := -math.Floor((x - step/2.0) / step)
	k_max := -k_min

	if left < -step/2.0+step*k_min {
		subIntervals = append(subIntervals, part{left: left, right: -step/2.0 + step*k_min})
	}

	for k := k_min; k <= k_max; k++ {
		subIntervals = append(subIntervals, part{left: -step/2.0 + step*k, right: step/2.0 + step*k})
	}

	if right > step/2.0+step*k_max {
		subIntervals = append(subIntervals, part{left: step/2.0 + step*k_max, right: right})
	}

	return subIntervals
}

func secantMethod(a float64, f func(x, a float64) float64, x0, x1 float64) (float64, error) {
	var x2, fx0, fx1, fx2 float64

	for i := 0; i < maxIterations; i++ {
		fx0 = f(x0, a)
		fx1 = f(x1, a)

		if math.Abs(fx1-fx0) < epsilon {
			return x1, nil // Успех: достигнута достаточная точность
		}

		// Метод секущих: x2 = x1 - f(x1) * (x1 - x0) / (f(x1) - f(x0))
		x2 = x1 - (fx1 * (x1 - x0) / (fx1 - fx0))
		fx2 = f(x2, a)

		if math.Abs(fx2) < epsilon {
			return x2, nil // Успех: найден корень с достаточной точностью
		}
		// fmt.Println(x2)
		// Переносим точки для следующей итерации
		x0, x1 = x1, x2
	}

	return 0, fmt.Errorf("Не удалось найти корень после %d итераций", maxIterations)
}

type part struct {
	left  float64
	right float64
}

func main() {
	var a float64
	fmt.Print("Введите значение параметра a: ")
	fmt.Scan(&a)

	// Метод простой итерации для x-a = log(x + 1)
	// resultSimple, errSimple := NewtonMethod(a, f1, derivative1)
	// if errSimple != nil {
	// 	fmt.Println(errSimple)
	// } else {
	// 	fmt.Printf("Метод простой итерации: x = %.6f\n", resultSimple)
	// }

	// // Метод простой итерации для x-a = log(x + 1)
	// resultSimple, errSimple = NewtonMethod(a, rev_f1, rev_f1_derivate)
	// if errSimple != nil {
	// 	fmt.Println(errSimple)
	// } else {
	// 	fmt.Printf("Метод простой итерации: x = %.6f\n", resultSimple)
	// }

	// // Метод секущих для x-a = log(x + 1)
	resultSecant, errSecant := secantMethod(a, f1, 0, a*a)
	if errSecant != nil {
		fmt.Println(errSecant)
	} else {
		fmt.Printf("Метод секущих: x = %.6f\n", resultSecant)
	}

	// Метод секущих для x-a = log(x + 1)
	resultSecant, errSecant = secantMethod(a, rev_f1, -a*a, 0)
	if errSecant != nil {
		fmt.Println(errSecant)
	} else {
		fmt.Printf("Метод секущих: x = %.6f\n", resultSecant)
	}

	// Метод простой итерации для x = l*sin(x)
	// Разбиваем весь интервал,
	// где возможно пересечение линейной функции с синусойидой,
	// на участки монотонности.
	// monotonicityIntervals := generateSubIntervalsAboutZero(a)

	// // интервал, содержащий 0.
	// center := len(monotonicityIntervals) / 2
	// for i, m := range monotonicityIntervals {
	// 	fmt.Println(m)
	// 	// Обработка прямой функции
	// 	resultSimple, errSimple := secantMethod(a, func(x, a float64) float64 { return a*math.Sin(x) - x }, m.left, m.right)
	// 	if errSimple != nil {
	// 		fmt.Println(errSimple)
	// 	} else {
	// 		fmt.Printf("Метод секущих: x = %.6f\n", resultSimple)
	// 	}

	// 	// Обработка обратной функции
	// 	period := float64((i - center)) * math.Pi
	// 	fmt.Println(period)
	// 	// resultSimple, errSimple = secantMethod(a, func(x, a float64) float64 { return -math.Asin(x/a) + period }, m.left, m.right)
	// 	resultSimple, errSimple = NewtonMethod(a, func(x, a float64) float64 { return -math.Asin(x/a) + period }, func(x, a float64) float64 { return -1.0 / (a * math.Sqrt(1.0-(x/a)*(x/a))) }, m.right)
	// 	if errSimple != nil {
	// 		fmt.Println(errSimple)
	// 	} else {
	// 		fmt.Printf("Метод секущих для обратной функции: x = %.6f\n", resultSimple)
	// 	}
	// }
}
