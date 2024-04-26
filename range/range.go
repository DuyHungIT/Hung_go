package main

import (
	"fmt"
)

func main() {
	var pow = []int{0, 3, 6, 9, 12, 15}
	for i, v := range pow {
		fmt.Printf("2 * %d = %d\n", i, v)
	}

	for h := range pow {
		fmt.Print("test: ", pow[h])
		fmt.Println()
	}
	// cách in gán giá trị của mảng vào biến
	for _, value := range pow {
		fmt.Print(value)
	}
}
