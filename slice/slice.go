package main

import (
	"fmt"
)

func printSlice(s []int) {

	if s == nil {
		fmt.Println("No ^_^")
	} else {
		fmt.Println("	=> length = ", len(s), ",cap = ", cap(s), " ", s)
	}
}
func main() {
	// khai báo một mảng gồn 7 số nguyên
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("mảng: ", arr)
	// slice length and capacity
	printSlice(arr)
	// cắt mảng, lấy từ giá trị 2 - 8
	arr = arr[2:9]
	fmt.Println("mảng: ", arr)
	printSlice(arr)

	// mảng 2 chiều
	var n, m int
	fmt.Scan(&n)
	fmt.Scan(&m)
	arr2 := make([][]int, n) // khai báo mảng 2 chiều arr với n hàng
	for i := range arr2 {
		arr2[i] = make([]int, m)
		for j := range arr2[i] {
			fmt.Println("nhập vào phần từ thứ %d, và cột %d", i+1, j+1)
			fmt.Scan(&arr2[i][j])
		}
	}
	// In ra mảng 2 chiều
	fmt.Println("Mảng 2 chiều vừa nhập là:")
	for i := range arr2 {
		fmt.Println(arr2[i])
	}
}
