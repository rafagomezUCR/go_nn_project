package main

import (
	"errors"
	"fmt"
	"math"
)

type Matrix struct {
	data [][]float64
	rows int
	cols int
}

func createMatrix(rows int, cols int) Matrix {
	matrix := Matrix{
		rows: rows,
		cols: cols,
	}
	matrix.data = make([][]float64, rows)
	for i := range rows {
		matrix.data[i] = make([]float64, cols)
	}
	return matrix
}

func (matrix *Matrix) printMatrix() {
	for i := range matrix.rows {
		for j := range matrix.cols {
			fmt.Print(matrix.data[i][j], " ")
		}
		fmt.Println()
	}
}

func transpose(matrix *Matrix) Matrix {
	new_matrix := createMatrix(matrix.cols, matrix.rows)
	for i := range matrix.rows {
		for j := range matrix.cols {
			new_matrix.data[j][i] = matrix.data[i][j]
		}
	}
	return new_matrix
}

func matrixMult(left *Matrix, right *Matrix) (Matrix, error) {
	if left.cols != right.rows {
		return Matrix{}, MatrixOperatioError{leftMatrix: *left, rightMatrix: *right, operation: "Matrix Multiplication"}
	}
	res := createMatrix(left.rows, right.cols)
	for i := range left.rows {
		for j := range right.cols {
			for k := range right.rows {
				res.data[i][j] += left.data[i][k] * right.data[k][j]
			}
		}
	}
	return res, nil
}

func matrixSubtraction(left *Matrix, right *Matrix) Matrix {
	result_matrix := createMatrix(left.rows, left.cols)
	for i := range left.rows {
		for j := range left.cols {
			result_matrix.data[i][j] = left.data[i][j] - right.data[i][j]
		}
	}
	return result_matrix
}

func elementwiseMatrixMult(matrix1 *Matrix, matrix2 *Matrix) Matrix {
	result_matrix := createMatrix(matrix1.rows, matrix1.cols)
	for i := range matrix1.rows {
		for j := range matrix1.cols {
			result_matrix.data[i][j] = matrix1.data[i][j] * matrix2.data[i][j]
		}
	}
	return result_matrix
}

func scalarMinusMatrix(scalar float64, matrix *Matrix) Matrix {
	result_matrix := createMatrix(matrix.rows, matrix.cols)
	for i := range matrix.rows {
		for j := range matrix.cols {
			result_matrix.data[i][j] = scalar - matrix.data[i][j]
		}
	}
	return result_matrix
}

func scalarMult(matrix *Matrix, scalar float64) Matrix {
	result_matrix := createMatrix(matrix.rows, matrix.cols)
	for i := range matrix.rows {
		for j := range matrix.cols {
			result_matrix.data[i][j] = scalar * matrix.data[i][j]
		}
	}
	return result_matrix
}

func activation(matrix *Matrix, activation string) (Matrix, error) {
	switch activation {
	case "s":
		return sigmoid(matrix), nil
	default:
		return Matrix{}, errors.New("proper activation function was not given ")
	}
}

func sigmoid(matrix *Matrix) Matrix {
	result_matrix := createMatrix(matrix.rows, matrix.cols)
	for i := range matrix.rows {
		for j := range matrix.cols {
			result_matrix.data[i][j] = 1 / (1 + math.Exp(-matrix.data[i][j]))
		}
	}
	return result_matrix
}
