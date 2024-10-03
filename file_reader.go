package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func readFile(filePath string) ([]Matrix, []Matrix) {
	var complete_input_list []Matrix
	var complete_target_list []Matrix
	file, err := os.Open(filePath)
	checkError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input_list, target := convertFileValuesToMatrix(&line, ",")
		target_list := createTargetList(target)
		complete_input_list = append(complete_input_list, input_list)
		complete_target_list = append(complete_target_list, target_list)
	}
	if err := scanner.Err(); err != nil {
		checkError(err)
	}
	return complete_input_list, complete_target_list
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
				val = 0.01
			}
			input_matrix.data[0][i-1] = val
		}
	}
	return input_matrix, target_value
}

func createTargetList(target float64) Matrix {
	target_list := Matrix{
		rows: 10,
		cols: 1,
		data: [][]float64{
			{0.1},
			{0.1},
			{0.1},
			{0.1},
			{0.1},
			{0.1},
			{0.1},
			{0.1},
			{0.1},
			{0.1},
		},
	}
	target_list.data[int(target)][0] = 0.99
	return target_list
}
