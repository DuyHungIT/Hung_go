package main

import "fmt"

func main() {
	var i interface{} = 56

	s, ok := i.(int)
	fmt.Println(s, ok) // biến ok dùng để xem chuyển đổi thành công không
	
}
