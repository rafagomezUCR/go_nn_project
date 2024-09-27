package main

import (
	"math"
	"math/rand/v2"
)

type NeuralNetwork struct {
	weights       []Matrix
	input_layer   Matrix
	target        Matrix
	hidden_layers []int
	learning_rate float64
	epochs        int
	activation    string
}

func CreateNetwork(input_list_size int, target_list_size int, hidden_layers []int, lr float64, ep int, activation string) NeuralNetwork {
	weights := InitializeWeights(input_list_size, hidden_layers, target_list_size)
	return NeuralNetwork{
		weights:       weights,
		hidden_layers: hidden_layers,
		learning_rate: lr,
		epochs:        ep,
		activation:    activation,
	}
}

func (nn *NeuralNetwork) Predict() Matrix {
	layer_input := transpose(&nn.input_layer)
	var layer_output Matrix
	for i := range len(nn.weights) {
		weighted_sum, e := matrixMult(&nn.weights[i], &layer_input)
		checkError(e)
		layer_output, e = activation(&weighted_sum, nn.activation)
		checkError(e)
		layer_input = layer_output
	}
	return layer_output
}

func (nn *NeuralNetwork) FeedFoward() []Matrix {
	layer_output_list := make([]Matrix, 0)
	layer_input := transpose(&nn.input_layer)
	layer_output_list = append(layer_output_list, layer_input)
	for i := range len(nn.weights) {
		weightem_sum, e := matrixMult(&nn.weights[i], &layer_input)
		checkError(e)
		layer_output, e := activation(&weightem_sum, nn.activation)
		checkError(e)
		layer_output_list = append(layer_output_list, layer_output)
		layer_input = layer_output
	}
	return layer_output_list
}

func (nn *NeuralNetwork) Train() {
	for range nn.epochs {
		layer_output_list := nn.FeedFoward()
		output_error, e := CalculateError(&layer_output_list[len(layer_output_list)-1], &nn.target)
		checkError(e)
		nn.Backpropogate(layer_output_list, &output_error)
	}
}

func (nn *NeuralNetwork) Backpropogate(layer_output_list []Matrix, output_error *Matrix) {
	for i := len(nn.weights) - 1; i >= 0; i-- {
		weight_t := transpose(&(nn.weights)[i])
		error_hidden, e := matrixMult(&weight_t, output_error)
		checkError(e)
		// right here goes gradient descent
		one_subtract_o := scalarSubtraction(&layer_output_list[i+1], 1.0)
		de_dw := elementwiseMatrixMult(&layer_output_list[i+1], &one_subtract_o)
		de_dw = elementwiseMatrixMult(output_error, &de_dw)
		hidden_outputs := transpose(&layer_output_list[i])
		de_dw, e = matrixMult(&de_dw, &hidden_outputs)
		checkError(e)
		de_dw = scalarMult(&de_dw, nn.learning_rate)
		nn.weights[i] = matrixSubtraction(&(nn.weights)[i], &de_dw)
		*output_error = error_hidden
		// weight_t := transpose(&nn.weights[i])
		// hidden_error, e := matrixMult(&weight_t, output_error)
		// checkError(e)
		// do_dw := scalarSubtraction(&layer_output_list[i+1], 1.0)
		// do_dw = elementwiseMatrixMult(&do_dw, &layer_output_list[i+1])
		// de_dw := elementwiseMatrixMult(output_error, &do_dw)
		// previous_output_t := transpose(&layer_output_list[i])
		// de_dw, e = matrixMult(&de_dw, &previous_output_t)
		// checkError(e)
		// de_dw = scalarMult(&de_dw, nn.learning_rate)
		// nn.weights[i] = matrixSubtraction(&nn.weights[i], &de_dw)
		// output_error = &hidden_error
	}
}

func InitializeWeights(input_layer_size int, neurons_per_layer []int, output_layer_size int) []Matrix {
	weightMatrices := make([]Matrix, 0)
	for i := range len(neurons_per_layer) + 1 {
		if i == 0 {
			w := createMatrix(neurons_per_layer[i], input_layer_size)
			weightMatrices = append(weightMatrices, w)
		} else if i == len(neurons_per_layer) {
			w := createMatrix(output_layer_size, neurons_per_layer[i-1])
			weightMatrices = append(weightMatrices, w)
		} else {
			w := createMatrix(neurons_per_layer[i], neurons_per_layer[i-1])
			weightMatrices = append(weightMatrices, w)
		}
	}
	for i := range len(weightMatrices) {
		for j := range weightMatrices[i].rows {
			for k := range weightMatrices[i].cols {
				weightMatrices[i].data[j][k] = rand.Float64() - 0.5
			}
		}
	}
	return weightMatrices
}

func CalculateError(output *Matrix, target *Matrix) (Matrix, error) {
	if output.rows != target.rows || output.cols != target.cols {
		return Matrix{}, MatrixOperatioError{leftMatrix: *output, rightMatrix: *target, operation: "Error Calculation"}
	}
	error_matrix := createMatrix(output.rows, output.cols)
	for i := range error_matrix.rows {
		for j := range error_matrix.cols {
			error_matrix.data[i][j] = math.Pow(target.data[i][j]-output.data[i][j], 2)
		}
	}
	return error_matrix, nil
}
