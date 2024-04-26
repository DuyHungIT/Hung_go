package main

import (
	"fmt"
	"math"
)

func add(fn func(float64, float64) float64) float64 {
	return fn(5, 8)
}
func main() {
	test := func(x, y float64) float64 {
		return x + y
	}

	fmt.Println(test(8, 7)) // kết quả là 8 +7
	fmt.Println(add(test))  // kết quả sẽ trả về 5+8
	fmt.Println(add(math.Pow)) // 5 mũ 8
}
