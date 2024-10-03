package main

import (
	"fmt"
)

func main() {
	learning_rate := 0.3
	hidden_layers := []int{100, 50}
	activation := "s"
	epochs := 1
	train_input_list, train_target_list := readFile("C:/Users/ruffl/Desktop/mnist data set/mnist_train.csv")
	test_input_list, test_target_list := readFile("C:/Users/ruffl/Desktop/mnist data set/mnist_test.csv")
	nn := CreateNetwork(train_input_list[0].cols, train_target_list[0].rows, hidden_layers, learning_rate, epochs, activation)
	nn.input_layer = test_input_list[0]
	a := nn.Predict()
	a.printMatrix()
	for i := range 10000 {
		nn.input_layer = train_input_list[i]
		nn.target = train_target_list[i]
		nn.Train()
	}
	fmt.Println()
	nn.input_layer = test_input_list[0]
	b := nn.Predict()
	b.printMatrix()
	test_target_list[0].printMatrix()
}
