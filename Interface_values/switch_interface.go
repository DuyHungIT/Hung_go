package main

import (
	"fmt"
)

func do(i interface{}) {
	switch t := i.(type) { // kiểu dữ liệu của i
	case int:
		fmt.Printf("%d đay là kiểu int\n", t)
	case string:
		fmt.Printf(" %d đây là kiểu string \n", t)
	default:
		fmt.Printf("%d không xác định được", t)
	}
}
func main() {
	do(15)
	do("Hung")
	do(false)
}
