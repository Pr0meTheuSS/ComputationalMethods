// package main

// func main() {
// }

package main

import (
	"fmt"
	"main/calcs"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func solution(x float64) float64 {
	return 1.0 / math.Exp(x)
}

func solution1(x float64) float64 {
	return math.Exp(x)*0.5*(math.Sin(x)+math.Cos(x)) - 0.5
}

func main() {
	step := 0.1
	left := 0.0
	right := 5.0
	n := (int64)((right - left) / step)

	x := make([]float64, n)
	for i := int64(0); i < n; i++ {
		x[i] = left + float64(i)*step
	}
	// y := calcs.SecondOrderDiff(x, .0, step, func(x, y float64) float64 { return math.Exp(x) * math.Cos(x) })
	y := calcs.FourthOrderDiff(x, 1, step)

	//	y := calcs.FourthOrderDiff(x, 1.0, step)
	// Создание точек для графика
	points1 := createPoints(x, y)

	ans := []float64{}
	for i := range x {
		ans = append(ans, solution(x[i]))
	}
	points2 := createPoints(x, ans)

	// Создание графика
	p := plot.New()
	p.Title.Text = "Functions"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	p.X.Min = -1
	p.X.Max = 6
	p.Y.Min = -1
	p.Y.Max = 1
	// Добавление линии для второго графика
	line2, err := plotter.NewLine(points2)
	if err != nil {
		fmt.Println("Ошибка создания линии:", err)
		return
	}
	line2.Color = plotutil.Color(1)
	p.Add(line2)

	// Добавление линии для первого графика
	line1, err := plotter.NewLine(points1)
	if err != nil {
		fmt.Println("Ошибка создания линии:", err)
		return
	}
	line1.Color = plotutil.Color(0)
	p.Add(line1)

	// Сохранение графика в файл (png)
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "plot.png"); err != nil {
		fmt.Println("Ошибка сохранения графика:", err)
		return
	}
}

// createPoints создает точки для графика на основе функции
func createPoints(x, y []float64) plotter.XYs {
	var points plotter.XYs
	for i := range x {
		points = append(points, plotter.XY{X: x[i], Y: y[i]})
	}
	return points
}
