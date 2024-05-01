package main

import (
	"fmt"
	"math"
)

func calculateCircumference(x int) float64 { // tính chu vi hình tròn, truyền vào bán kính
	return math.Pi * float64(x*2)
}
func main() {

	var i int
	var cv1 float64

	fmt.Print("Nhập vào giá trị bán kính: ")
	fmt.Scan(&i)
	fmt.Scanln()

	// gọi lại hàm tính chu vi
	cv := calculateCircumference(i)
	fmt.Println("chu vi= ", cv)
	// làm tròn chu vi

	cv_down := math.Floor(cv) // làm tròn xuống giá trị nguyên
	new_cv := cv - cv_down
	if new_cv >= 0.5 {
		cv1 = math.Ceil(cv)
	} else {
		cv1 = math.Floor(cv)
	}

	fmt.Println("Gía trị làm tròn chu vi: ", cv1)
}
