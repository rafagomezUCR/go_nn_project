package main

import (
	"fmt"
	"os"
)

func checkError(e error) {
	if e != nil {
		fmt.Print(e)
		os.Exit(1)
	}
}

type MatrixOperatioError struct {
	leftMatrix  Matrix
	rightMatrix Matrix
	operation   string
}

func (e MatrixOperatioError) Error() string {
	return fmt.Sprintf("Matrix Dimensions Don't Match doing %s.\n Left Matrix Dimension\n rows: %d, cols: %d\n Right Matrix Dimension\n rows: %d, cols: %d\n", e.operation, e.leftMatrix.rows, e.leftMatrix.cols, e.rightMatrix.rows, e.rightMatrix.cols)
}
