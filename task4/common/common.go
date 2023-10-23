package common

import (
	"bufio"
	"fmt"
	"io"
	"main/logger"
	"strings"

	"gonum.org/v1/gonum/mat"
)

func PrintMatrix(A mat.Matrix) {
	fa := mat.Formatted(A, mat.Squeeze())
	fmt.Println(fa)
}

func PrintVector(v mat.Vector) {
	fv := mat.Formatted(v, mat.Squeeze())
	fmt.Println(fv)
}

func MatrixInfo(A *mat.Dense) {
	l := logger.GetLoggerInstance()
	fa := mat.Formatted(A, mat.Squeeze())
	fmt.Printf("матрица:=\n%v \n", fa)

	norm := mat.Norm(A, 2)
	fmt.Printf("Норма матрицы: %f\n", norm)

	var eig mat.Eigen
	ok := eig.Factorize(A, mat.EigenLeft)
	if !ok {
		l.Error("Eigendecomposition failed")
		fmt.Printf("Собственные значения не найдены.\n")
	} else {
		fmt.Printf("Собственные значения:\n%v\n", eig.Values(nil))
	}

	conditionNumber := mat.Cond(A, 2)
	fmt.Printf("Число обусловленности матрицы: %f\n", conditionNumber)
}

func ReadMatrix(in io.Reader) *mat.Dense {
	var lines []string
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var matrix [][]float64
	for _, line := range lines {
		fields := strings.Fields(line)
		var row []float64
		for _, field := range fields {
			value := 0.0
			fmt.Sscanf(field, "%f", &value)
			row = append(row, value)
		}
		matrix = append(matrix, row)
	}

	rows := len(matrix)
	cols := len(matrix[0])
	data := make([]float64, rows*cols)
	ld := cols
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			data[i*ld+j] = matrix[i][j]
		}
	}

	return mat.NewDense(rows, cols, data)
}

func ReadVector(in io.Reader) *mat.VecDense {
	var values []float64
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		value := 0.0
		fmt.Sscanf(line, "%f", &value)
		values = append(values, value)
	}

	return mat.NewVecDense(len(values), values)
}
