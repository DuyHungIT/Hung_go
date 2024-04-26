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
		fmt.Printf("%d không xác định được\n", t)
	}
}

type IPaddr [4]byte

func (ip IPaddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	do(15)
	do("Hung")
	do(false)
	// khai báo map
	network := map[string]IPaddr{
		"ip_network":  {10, 100, 100, 115},
		"ip_network2": {10, 120, 120, 25},
	}
	for name, ip := range network {
		fmt.Printf("%v: %v\n", name, ip)
	}

}
