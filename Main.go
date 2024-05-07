package main

import (
	"fmt"
	person "main/Person"
)

func main() {

	ob := new(person.Person)
	// Set name
	err := ob.SetName("Hung")
	if err != nil {
		fmt.Println("Lỗi tên:", err)
	}
	// Set birthdayYear
	err = ob.SetBirthdayYear(1883)
	if err != nil {
		fmt.Println("Lỗi năm:", err)
	}
	// Set email
	err = ob.SetEmail("hung@asdf.com")
	if err != nil {
		fmt.Println("Lỗi email:", err)
	}
	// Set phone
	err = ob.SetPhone("1234567890")
	if err != nil {
		fmt.Println("Lỗi số điệnh thoại:", err)
	}

	// In thông tin
	fmt.Println("Name:", ob.Name)
	fmt.Println("Birthday Year:", ob.BirthdayYear)
	fmt.Println("Age:", ob.Age)
	fmt.Println("Email:", ob.Email)
	fmt.Println("Phone:", ob.Phone)
}
