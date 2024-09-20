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

func CreateMatrix(rows int, cols int) Matrix {
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

func (matrix *Matrix) PrintMatrix() {
	for i := range matrix.rows {
		for j := range matrix.cols {
			fmt.Print(matrix.data[i][j], " ")
		}
		fmt.Println()
	}
}

func Transpose(matrix *Matrix) Matrix {
	new_matrix := CreateMatrix(matrix.cols, matrix.rows)
	for i := range matrix.rows {
		for j := range matrix.cols {
			new_matrix.data[j][i] = matrix.data[i][j]
		}
	}
	return new_matrix
}

func (matrix *Matrix) Transpose() {
	temp_matrix := CreateMatrix(matrix.cols, matrix.rows)
	for i := range matrix.rows {
		for j := range matrix.cols {
			temp_matrix.data[j][i] = matrix.data[i][j]
		}
	}
	matrix = &temp_matrix
}

func MatrixMult(left *Matrix, right *Matrix) (Matrix, error) {
	if left.cols != right.rows {
		return Matrix{}, MatrixOperatioError{leftMatrix: *left, rightMatrix: *right, operation: "Matrix Multiplication"}
	}
	res := CreateMatrix(left.rows, right.cols)
	for i := range left.rows {
		for j := range right.cols {
			for k := range right.rows {
				res.data[i][j] += left.data[i][k] * right.data[k][j]
			}
		}
	}
	return res, nil
}

func Activation(matrix *Matrix, activation string) (Matrix, error) {
	switch activation {
	case "s":
		return Sigmoid(matrix), nil
	default:
		return Matrix{}, errors.New("proper activation function was not given ")
	}
}

func Sigmoid(matrix *Matrix) Matrix {
	result_matrix := CreateMatrix(matrix.rows, matrix.cols)
	for i := range matrix.rows {
		for j := range matrix.cols {
			result_matrix.data[i][j] = 1 / (1 + math.Exp(-matrix.data[i][j]))
		}
	}
	return result_matrix
}
