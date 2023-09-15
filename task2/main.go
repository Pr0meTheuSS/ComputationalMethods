package main

import (
	"image/color"
	"main/interpolator"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// lg(x) + 7 / (2x + 6)
func originalFunction(x float64) float64 {
	return math.Log(x) + 7/(2*x+6)
}

type Interval struct {
	left  float64
	right float64
}

const (
	nodesAmount = 2
)

func main() {
	nodes := []interpolator.Node{}
	interval := Interval{
		left:  1.0,
		right: 4.0,
	}

	step := (interval.right - interval.left) / nodesAmount
	currentX := interval.left
	for i := 0; i <= nodesAmount; i++ {
		nodes = append(nodes, interpolator.Node{X: currentX, Y: originalFunction(currentX)})
		currentX += step
	}

	p := plot.New()

	p.Title.Text = "Functions"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	originalFuncGraph := plotter.NewFunction(originalFunction)
	originalFuncGraph.Color = color.RGBA{B: 255, A: 255}

	p.Add(originalFuncGraph)
	p.Legend.Add("lg(x)+7/(2x+6)", originalFuncGraph)
	p.X.Min = interval.left
	p.X.Max = interval.right
	p.Y.Min = 0
	p.Y.Max = 3

	interpolator := interpolator.LagrangeInterpolator{}

	interpolarResultGraph := plotter.NewFunction(func(x float64) float64 {
		return interpolator.CalcPolynomialValue(x, nodes)
	})

	interpolarResultGraph.Dashes = []vg.Length{vg.Points(2), vg.Points(2)}
	interpolarResultGraph.Width = vg.Points(2)
	interpolarResultGraph.Color = color.RGBA{R: 255, A: 255}

	p.Add(interpolarResultGraph)
	p.Legend.Add("interpolar", interpolarResultGraph)

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "lab2.png"); err != nil {
		panic(err)
	}
}
