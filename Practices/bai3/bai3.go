package main

import (
	"fmt"
)

func Det(x [][]float64) float64 {
	det := (x[0][0] * (x[1][1]*x[2][2] - x[1][2]*x[2][1])) -
		(x[0][1] * (x[1][0]*x[2][2] - x[1][2]*x[2][0])) +
		(x[0][2] * (x[1][0]*x[2][1] - x[1][1]*x[2][0]))
	return det
}
func main() {

	myArrays := [][]float64{{1, 2, 3}, {4, 10, 6}, {7, 8, 9}}
	det := Det(myArrays)
	fmt.Println("det = ", det)

}
