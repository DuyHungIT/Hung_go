package main

import (
	"fmt"
	"time"
)

// khai báo kiểu dữ liệu myerrors
type Myerror struct {
	when time.Time
	what string // lỗi
}

func (v *Myerror) Error() string {
	return fmt.Sprintf("thời gian %v, : %v ", v.when, v.what)
}

func run() error {
	return &Myerror{
		time.Now(), "dont work",
	}
}
func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
	}
}
