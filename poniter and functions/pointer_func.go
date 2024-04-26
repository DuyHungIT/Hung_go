package main

import "fmt"

type hcn struct {
	cr float64
	cd float64
}

// thay đổi giá trị cr cd
func change(hcn *hcn) {
	hcn.cr = hcn.cr + 5
	hcn.cd = hcn.cd + 5

}

// lưu ý cách gọi lại hàm
func (hcn *hcn) scale(float64) {
	hcn.cd = hcn.cd + 10
	hcn.cr = hcn.cr + 10
}

// tính diện tích
func acreage(hcn hcn) float64 {
	return hcn.cd * hcn.cr
}

// tính chu vi

func main() {
	hcn := hcn{2, 6}
	fmt.Println("diện tích lúc chưa thay đổi: ", acreage(hcn))
	//
	change(&hcn)
	fmt.Println("diện tích lúc  thay đổi: ", acreage(hcn))
	// lưu ý cáhc gọi lại hàm
	hcn.scale(2) // Go tự hiểu là &hcn
	fmt.Println("diện tích final: ", acreage(hcn))
}
