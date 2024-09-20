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

func CreateNetwork(input Matrix, target Matrix, hidden_layers []int, lr float64, ep int, activation string) NeuralNetwork {
	weights := InitializeWeights(len(input.data), hidden_layers, len(target.data))
	return NeuralNetwork{
		weights:       weights,
		input_layer:   input,
		target:        target,
		hidden_layers: hidden_layers,
		learning_rate: lr,
		epochs:        ep,
		activation:    activation,
	}
}

func (nn *NeuralNetwork) Predict() Matrix {
	layer_input := nn.input_layer
	var layer_output Matrix
	for i := range len(nn.weights) {
		weighted_sum, e := MatrixMult(&nn.weights[i], &layer_input)
		checkError(e)
		layer_output, e = Activation(&weighted_sum, nn.activation)
		checkError(e)
		layer_input = layer_output
	}
	return layer_output
}

func (nn *NeuralNetwork) FeedFoward() []Matrix {
	layer_output_list := make([]Matrix, 0)
	layer_output_list = append(layer_output_list, nn.input_layer)
	layer_input := nn.input_layer
	for i := range len(nn.weights) {
		weightem_sum, e := MatrixMult(&nn.weights[i], &layer_input)
		checkError(e)
		layer_output, e := Activation(&weightem_sum, nn.activation)
		checkError(e)
		layer_output_list = append(layer_output_list, layer_output)
		layer_input = layer_output
	}
	return layer_output_list
}

func (nn *NeuralNetwork) Train() {
	layer_output_list := nn.FeedFoward()
	output_error, e := CalculateError(&layer_output_list[len(layer_output_list)-1], &nn.target)
	checkError(e)
	nn.Backpropogate(layer_output_list, &output_error)
}

func (nn *NeuralNetwork) Backpropogate(layer_output_list []Matrix, output_error *Matrix) {

}

func InitializeWeights(input_layer_size int, neurons_per_layer []int, output_layer_size int) []Matrix {
	weightMatrices := make([]Matrix, 0)
	for i := range len(neurons_per_layer) + 1 {
		if i == 0 {
			w := CreateMatrix(neurons_per_layer[i], input_layer_size)
			weightMatrices = append(weightMatrices, w)
		} else if i == len(neurons_per_layer) {
			w := CreateMatrix(output_layer_size, neurons_per_layer[i-1])
			weightMatrices = append(weightMatrices, w)
		} else {
			w := CreateMatrix(neurons_per_layer[i], neurons_per_layer[i-1])
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
	error_matrix := CreateMatrix(output.rows, output.cols)
	for i := range error_matrix.rows {
		for j := range error_matrix.cols {
			error_matrix.data[i][j] = math.Pow(output.data[i][j]-target.data[i][j], 2)
		}
	}
	return error_matrix, nil
}
