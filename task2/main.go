package main

import (
	"fmt"
	"image/color"
	"log"
	"main/interpolator"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// lg(x) + 7 / (2x + 6)
func originalFunction(x float64) float64 {
	if x <= 0 {
		log.Fatal("Invalid function param")
	}
	return math.Log(x) + 7/(2*x+6)
}

type Interval struct {
	left  float64
	right float64
}

func main() {
	nodes := []interpolator.Node{}
	interval := Interval{
		left:  1.0,
		right: 4.0,
	}
	fmt.Println("Enter [a,b] like two numbers:")
	if _, err := fmt.Scan(&interval.left, &interval.right); err != nil {
		log.Fatal("Invalid input")
	}
	nodesAmount := 0
	fmt.Println("Enter nodes amount:")
	if _, err := fmt.Scan(&nodesAmount); err != nil {
		log.Fatal("Invalid input")
	}

	errorPointX := interval.left
	fmt.Println("Error point (x):")
	if _, err := fmt.Scan(&errorPointX); err != nil {
		log.Fatal("Invalid input")
	}

	step := (interval.right - interval.left) / float64(nodesAmount)
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
	p.Y.Min = originalFunction(p.X.Min)
	p.Y.Max = originalFunction(p.X.Max)

	var interpolator interpolator.InterpolatorInterface = interpolator.LagrangeInterpolator{}

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

	diff := originalFunction(errorPointX) - interpolator.CalcPolynomialValue(errorPointX, nodes)
	fmt.Printf("(function(x) - interpolator(x))^2 %f: %f\n", errorPointX, diff*diff)
}
