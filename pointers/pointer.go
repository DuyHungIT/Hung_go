package main

import (
	"fmt"
)

func test(x *int) int {
	*x += 5
	return *x
}
func main() {

	x, y := 10, 50

	p := &x                                                // khai báo con trỏ p bằng giá trị của x
	fmt.Println("giá trị của i thông qua con trỏ p: ", *p) // in giá trị của x thông qua toán tử *

	p = &y // đổi con trỏ p sang giá trị của y
	fmt.Println("giá trị của y thông p: ", *p)

	//
	fmt.Println("giá trị của ", x, " sau khi +5 la: ", test(&x))
	fmt.Println(x)
}
