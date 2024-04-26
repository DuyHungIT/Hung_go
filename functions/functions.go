package main

import (
	"fmt"
)

// hàm trả về tổng 2 số
func add(x int, y int) int {
	return x + y
}

// trả về số lơn hơn
func max(x int, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

// đổi 2 dữ liệu kiểu string
func swap_string(x, y string) (string, string) {
	var temp string = x
	x = y
	y = temp
	return x, y
}

// trả về kiểu bool
func return_bool(x, y int) bool {
	if (x+y)%2 == 0 {
		return true
	} else {
		return false
	}
}

// ------------------------------------
func main() {
	fmt.Println(add(45, 1))
	// khai báo kiểu 1 x, y theo kiểu dữ liệu nguyên
	var x, y int
	//var z = 1
	//var u int = 1
	// nhập giá trị x, y từ bàn phím
	fmt.Scan(&x)
	fmt.Scan(&y)

	//trả về số lớn hơn
	fmt.Println("giá trị lớn hơn: ", max(x, y))

	// trả về kiểu bool
	if return_bool(x, y) == true {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	// khai báo kiểu 2
	a, b := "hung", "Duy"
	fmt.Println("giá trị a : ", a)
	fmt.Println("giá trị b : ", b)

	// đổi chỗ string
	a, b = swap_string(a, b)
	fmt.Println("sau khi đổi chỗ: ", a, b)

	// khai báo hằng số
	const country = "Viet Nam"
	const number = 12345

	fmt.Println("hằng số: ", country, number)
}
