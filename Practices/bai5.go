package main

import (
	"fmt"
)

func printFibonacci(n int) {
	arr := make([]int, n)
	arr[0] = 0
	arr[1] = 1
	for i := 2; i < n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	for i := 0; i < n; i++ {
		fmt.Printf(" %v  ", arr[i])
	}
}
func main() {
	// Hiển thị n số đầu của dãy fibonacci
	var n int
	fmt.Println("Nhập vào N số fibonacci bạn muốn trả ra: ")
	fmt.Scan(&n)
	printFibonacci(n)
}
