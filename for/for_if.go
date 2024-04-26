package main

import "fmt"

// tính tổng các số lẻ từ 1 đến n. n nhập từ bàn phím
func add(n int) int {
	sum := 0
	for i := 1; i <= n; i++ { // duyệt hết các số từ 1 đến n
		if i%2 != 0 { // kiểm tra số lẻ
			sum += i
		}
	}
	return sum
}

// kiểm tra 3 số nhập vào có phải 3 cạnh tam giác ko
// trả về 1 nếu đúng, 0 nếu sai
func check(a, b, c int) int {
	if a+b > c && a+c > b && b+c > a {
		return 1
	} else {
		return 0
	}
}
func main() {
	var n int
	fmt.Println("Nhập vào số n: ")
	fmt.Scan(&n) // nhập n từ bàn phím
	fmt.Scanln() // sử lý kí tự xuống dòng, sau câu lệnh nhập vào số nguyên,...
	fmt.Println("tổng các số lẻ từ 1 đến ", n, " là: ", add(n))

	// tính tổng sum đến khi nhỏ hơn 1000
	sum := 1
	for sum < 1000 {
		sum += sum
	}

	fmt.Println("sum: ", sum)

	// vòng lặp vô hạn thì cần bỏ qua điều kiện lặp
	// for{
	//
	//}

	// kiểm tra 3 cạnh tam giác
	fmt.Println("==: ", check(3, 4, 5))
}
