package main

import (
	"fmt"
	"log"
	"main/trig"
	"math"

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

const (
	a          = 1.0         // константа для f(u) = cu
	dx         = 0.01        // шаг по пространству
	dt         = 0.01        // шаг по времени
	r          = a * dt / dx // число Куранта
	tMax       = 5.0         // максимальное время
	spaceStart = 1.0
	space      = 5.0
)

func initialCondition(x float64) float64 {
	switch {
	case 0 <= x && x < 1:
		return 0
	case 1 <= x && x < 4:
		return math.Sin(math.Pi * (x - 1) / 3.0)
	case 4 <= x && x < 5:
		return 0
	}
	return 0
	// switch {
	// case -1 <= x && x < 0:
	// 	return 3
	// case 0 <= x && x < 1:
	// 	return 3
	// }

	// return 0
}

func laxNewlayer(uOld []float64) []float64 {
	lenth := len(uOld)
	uNewLayer := make([]float64, lenth)

	// Перенос краевых значений
	uNewLayer[0] = uOld[0]
	uNewLayer[lenth-1] = uOld[lenth-1]

	for i := 0; i < len(uOld)-3; i++ {
		uNewLayer[i+1] = 0.5 * (uOld[i]*(1+r) + uOld[i+2]*(1-r))
	}

	return uNewLayer
}

func laxNewlayerQuasiLinear(uOld []float64, f func(u_ij float64) float64) []float64 {
	lenth := len(uOld)
	uNewLayer := make([]float64, lenth)

	// Перенос краевых значений
	uNewLayer[0] = uOld[0]
	uNewLayer[lenth-1] = uOld[lenth-1]

	for i := 0; i < len(uOld)-3; i++ {
		uNewLayer[i+1] = 0.5 * (uOld[i]*(1+f(uOld[i])*dt/dx) + uOld[i+2]*(1-f(uOld[i])*dt/dx))
	}

	return uNewLayer
}

func buildVecB(uOld []float64, uNew []float64) mat.Vector {
	data := make([]float64, len(uOld)-2)
	data[0] = uOld[1] + (r/2.0)*uOld[0]
	data[len(data)-1] = uOld[len(uOld)-2] - (r/2.0)*uOld[len(uOld)-1]

	for i := 1; i < len(data)-1; i++ {
		data[i] = uOld[i+1]
	}

	return mat.NewVecDense(len(data), data)
}

func buildMatrixA(size int, old []float64, f func(float64) float64) mat.Matrix {
	data := make([]float64, size*size)

	for i := 0; i < size; i++ {
		data[i*size+i] = 1
		if i < size-1 {
			data[i*size+i+1] = -f(old[i]) * dt / (2.0 * dx)
		}
		if i > 0 {
			data[i*size+i-1] = f(old[i]) * dt / (2.0 * dx)
		}
	}

	return mat.NewDense(size, size, data)
}

func implicitScheme(uOld []float64, f func(float64) float64) []float64 {
	uNew := make([]float64, len(uOld))

	result, _ := trig.NewTrigResolver().Resolve(buildMatrixA(len(uOld)-2, uOld, f), buildVecB(uOld, uNew))

	ret := make([]float64, len(uOld))

	// перернос значений с краевых точек.
	ret[0] = uOld[0]
	ret[len(uOld)-1] = uOld[len(uOld)-1]

	for i := 1; i < len(uOld)-1; i++ {
		ret[i] = result.AtVec(i - 1)
	}

	return ret
}

func plotResults(x []float64, u []float64, title, filename string, nodes []float64) {
	p := plot.New()

	p.Title.Text = title
	p.X.Label.Text = "x"
	p.Y.Label.Text = "u"

	points := make(plotter.XYs, len(x))
	for i := range points {
		points[i].X = x[i]
		points[i].Y = u[i]
	}

	line, err := plotter.NewLine(points)
	if err != nil {
		log.Fatal(err)
	}

	p.Add(line)

	if err := p.Save(6*vg.Inch, 4*vg.Inch, filename); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Инициализация сетки
	nx := int(space / dx)
	x := make([]float64, nx)
	u := make([]float64, nx)

	for i := 0; i < nx; i++ {
		x[i] = spaceStart + float64(i)*dx
		u[i] = initialCondition(x[i])
	}

	// Анимация для f(u) = cu
	var framesCu [][]float64
	for t := 0.0; t <= tMax; t += dt {
		// u = laxNewlayer(u)
		//u = laxNewlayerQuasiLinear(u, func(u float64) float64 { return a })
		u = implicitScheme(u, func(u float64) float64 { return a })
		frame := make([]float64, len(u))
		copy(frame, u)
		framesCu = append(framesCu, frame)
	}

	// Сохранение анимации для f(u) = cu
	for i, frame := range framesCu {
		filename := fmt.Sprintf("frame_cu_%03d.png", i)
		plotResults(x, frame, fmt.Sprintf("Шаг времени %d", i), filename, u)
	}

	// // Анимация для f(u) = u^2/2
	// u = make([]float64, nx) // сбросим значения
	// for i := 0; i < nx; i++ {
	// 	u[i] = initialCondition(x[i])
	// }
	// var framesU2 [][]float64
	// for t := 0.0; t <= tMax; t += dt {
	// 	u = laxNewlayerQuasiLinear(u, func(u float64) float64 { return u * u / 2.0 })
	// 	// u = implicitScheme(u, func(u float64) float64 { return u * u / 2.0 })
	// 	frame := make([]float64, len(u))
	// 	copy(frame, u)
	// 	framesU2 = append(framesU2, frame)
	// }

	// // Сохранение анимации для f(u) = u^2/2
	// for i, frame := range framesU2 {
	// 	filename := fmt.Sprintf("frame_u2_%03d.png", i)
	// 	plotResults(x, frame, fmt.Sprintf("Шаг времени %d", i), filename, u)
	// }
}
