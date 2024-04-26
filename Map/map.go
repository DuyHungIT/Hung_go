package main

import (
	"fmt"
)

type Vertex struct {
	Kinhdo, Vido float64
}

func nhapMap() map[string]Vertex {
	var name string
	var lat, long float64 // lat kinh độ, long vĩ độ
	fmt.Println("nhập khóa: ")
	fmt.Scanln(&name)
	fmt.Println("nhập kinh độ: ")
	fmt.Scan(&lat)
	fmt.Scanln()
	fmt.Println("nhập vĩ độ: ")
	fmt.Scan(&long)
	fmt.Scanln()
	return map[string]Vertex{
		name: {
			Kinhdo: lat,
			Vido:   long,
		},
	}
}
func main() {
	// KHAI BÁO một map tên qp
	var qp = map[string]Vertex{
		"Quynh Phu": Vertex{
			17, 18,
		},
		"Nam Định": Vertex{
			18, 19,
		},
	}
	fmt.Println(qp)

	// khai bao một mảng nhiều map
	n := 4
	maps := make([]map[string]Vertex, n)
	for i := 0; i < n; i++ {
		maps[i] = nhapMap()
	}

	// in ra các map
	fmt.Println("các map: ")
	for i, m := range maps {
		fmt.Printf("map thứ %d : %v", i+1, m)
	}
}
