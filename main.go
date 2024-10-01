package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	learning_rate := 0.3
	hidden_layers := []int{100, 50}
	activation := "s"
	epochs := 1
	var train_input_layers = []Matrix{}
	var test_input_layers = []Matrix{}
	var train_traget_list = []Matrix{}
	var test_target_list = []Matrix{}
	train_file, train_file_err := os.Open("C:/Users/ruffl/Desktop/mnist data set/mnist_train.csv")
	test_file, test_file_err := os.Open("C:/Users/ruffl/Desktop/mnist data set/mnist_test.csv")
	checkError(train_file_err)
	checkError(test_file_err)
	defer train_file.Close()
	defer test_file.Close()
	train_scanner := bufio.NewScanner(train_file)
	test_scanner := bufio.NewScanner(test_file)
	for test_scanner.Scan() {
		line := test_scanner.Text()
		input_list, target := convertFileValuesToMatrix(&line, ",")
		target_list := createTargetList(target)
		test_input_layers = append(test_input_layers, input_list)
		test_target_list = append(test_target_list, target_list)
	}
	if err := test_scanner.Err(); err != nil {
		fmt.Println("Error reading test files: ", err)
	}
	for train_scanner.Scan() {
		line := train_scanner.Text()
		input_list, target := convertFileValuesToMatrix(&line, ",")
		target_list := createTargetList(target)
		train_input_layers = append(train_input_layers, input_list)
		train_traget_list = append(train_traget_list, target_list)
	}
	if err := train_scanner.Err(); err != nil {
		fmt.Println("Error reading file: ", err)
	}
	nn := CreateNetwork(train_input_layers[0].cols, train_traget_list[0].rows, hidden_layers, learning_rate, epochs, activation)
	nn.input_layer = test_input_layers[0]
	a := nn.Predict()
	a.printMatrix()
	for i := range 10000 {
		nn.input_layer = train_input_layers[i]
		nn.target = train_traget_list[i]
		nn.Train()
	}
	fmt.Println()
	nn.input_layer = test_input_layers[0]
	b := nn.Predict()
	b.printMatrix()
	test_target_list[0].printMatrix()
}

func convertFileValuesToMatrix(input *string, delimiter string) (Matrix, float64) {
	strs := strings.Split(*input, delimiter)
	input_matrix := createMatrix(1, len(strs)-1)
	var target_value float64
	for i, str := range strs {
		val, err := strconv.ParseFloat(str, 64)
		checkError(err)
		if i == 0 {
			target_value = val
		} else {
			val = (val / 255 * 0.99)
			if val == 0 {
				val += 0.01
			}
			input_matrix.data[0][i-1] = val
		}
	}
	return input_matrix, target_value
}

func createTargetList(target float64) Matrix {
	target_list := Matrix{
		data: [][]float64{
			{0.01},
			{0.01},
			{0.01},
			{0.01},
			{0.01},
			{0.01},
			{0.01},
			{0.01},
			{0.01},
			{0.01},
		},
		rows: 10,
		cols: 1,
	}
	target_list.data[int(target)][0] = 0.99
	return target_list
}
