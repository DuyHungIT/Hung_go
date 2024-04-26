package main

import (
	"fmt"
	"math"
)

// định dạng kiểu dữ liệu mới
type errorNegative float64

// in ra thông báo lỗi
func (v errorNegative) Error() string {
	return fmt.Sprintf("i cant sqrt %v ", v)
}

// kiểm tra
func test(v float64) (float64, error) {
	if v < 0 {
		return 0, errorNegative(v)
	} else {
		num := math.Sqrt(v)
		return num, nil
	}

}
func main() {
	fmt.Println(test(-2))
	fmt.Println(test(2))
}
