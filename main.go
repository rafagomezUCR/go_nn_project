package main

func main() {
	a := Matrix{
		data: [][]float64{
			{3, 3},
			{2, 2},
		},
		rows: 2,
		cols: 2,
	}
	b := Matrix{
		data: [][]float64{
			{3, 3},
			{3, 3},
			{3, 3},
		},
		rows: 3,
		cols: 2,
	}
	// fmt.Println(a.rows, a.cols, b.rows, b.cols)
	// c, e := MatrixMult(&a, &b)
	// checkError(e)
	// c.PrintMatrix()
	h := []int{100, 50}
	a.Transpose()
	nn := CreateNetwork(a, b, h, 0.3, 1, "s")
	m := nn.Predict()
	m.PrintMatrix()
}
