// cho số  z tính bình phương của x gần nhất
package main

import (
	"fmt"
	"math"
)

func main() {
	// nhập vào số nguyên z
	var z float64
	fmt.Scan(&z)

	// tính căn bậc 2 của z
	sqrtZ := math.Sqrt(z)
	d := sqrtZ - math.Floor(sqrtZ) // tính phần số thập phân của sqrtZ
	// math.floor làm tròn xuống, math.ceil làm tròn lên
	var x float64
	// so sánh số thập phân rồi làm tròn
	// sẽ ra được x gần nhất
	if d >= 0.5 {
		x = math.Ceil(sqrtZ)
	} else {
		x = math.Floor(sqrtZ)
	}

	fmt.Println("số bình phương gần nhất là: ", math.Pow(x, 2))
}
