package main

import (
	"fmt"
)

type SinhVien struct {
	ten  string
	tuoi int
}

// khai báo 4 sinh viên
var (
	hs1 = SinhVien{}                    // tên rỗng, tuổi 0
	hs2 = SinhVien{ten: "Nguyen Van A"} // tuoi 0\
	hs3 = SinhVien{"Van b", 29}
	hs4 = &SinhVien{"u", 98}
)

func main() {
	hung := SinhVien{"Hưng", 19} // hưng
	fmt.Println(hung)
	// sử dụng con trỏ trong struct
	p := &hung

	p.tuoi = 20 // gán cho tuổi của hưung bằng 20
	fmt.Println(hung)
	// in ra 4 sinh vien
	fmt.Println(
		"sinh vien 1", hs1, "\n",
		"sinh viên 2", hs2, "\n",
		"sinhh viên 3", hs3, "\n",
		"sinh vien 4", hs4,
	)
	// gán thông tin của hưng cho 1 sinh viên mới
	a := hung
	fmt.Println("sinh viên a: ", a)
	// thay đổi thông tin của sinh viên hưng
	fmt.Println("thông tin ban đầu của Hưng: ", hung)
	hung.ten = "Duy Hung"
	hung.tuoi = 30
	fmt.Println("thông tin sau khi thay đổi: ", hung)

	// khai báo một mảng gồm 3 sinh viên
	var mang_SV [3]SinhVien
	// gán giá trị cho mảng sv
	mang_SV[0] = hs2
	mang_SV[1] = hs3
	mang_SV[2] = hung
	// in mảng
	fmt.Println("thông tin mảng sinh viên:", mang_SV)
	// thử dùng vòng for để gán trị cho mảng sv gồm n sinh viên
	fmt.Println("-------------------------------------------")
	var n int
	fmt.Scan(&n)
	// khai báo mangsv1 có n phần tử sinh viên
	mang_SV1 := make([]SinhVien, n)
	for i := 0; i < n; i++ {
		var sv SinhVien
		fmt.Println("nhập thông tin sinh viên ", i+1)
		fmt.Print("tên: ")
		fmt.Scan(&sv.ten)
		fmt.Print("tuổi: ")
		fmt.Scan(&sv.tuoi)
		// gán giá trị mảng = sv
		mang_SV1[i] = sv
	}
	// in ra mảng sv1
	fmt.Println("thông tin mảng SV1: ", mang_SV1)
}
