package main

import (
	"fmt"
	"runtime"
)

func main() {
	// defer  - trì hoãn một câu lệnh đến khi hàm bao quanh trả
	// về
	defer fmt.Println("The end")
	fmt.Print("Go runs on ")
	os := runtime.GOOS // trả về hệ điều hành đang dùng
	switch os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

}
