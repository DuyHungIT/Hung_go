package main

import (
	//các gói math/rand sẽ bao gồm các câu lệnh bắt đầu bằng rand
	//khi import thì các lệnh import phải xuống dòng
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("my number favourite is", rand.Intn(10)) // rand.Intn(x) chọn số nguyên ngẫu nhiên từ 1 - 10

}
