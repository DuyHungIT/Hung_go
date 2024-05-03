package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// khai báo 1 struct Person
type Person struct {
	name         string
	birthdayYear int
	age          int
	email        string
	phone        string
}

// interface
type PersonInterface interface {
	setName(string) error
	setBirthdayYear(int) error
	setEmail(string) error
	setPhone(interface{}) error
}

// // setName : set vào name của Person
func (ob *Person) setName(name string) error {
	if name == "" {
		return errors.New("lỗi, không dược để trống")
	}
	// Check if the first letter is uppercase
	firstLetter := name[:1]
	if firstLetter != strings.ToUpper(firstLetter) {
		return errors.New("Lỗi, đọc lại yêu cầu")
	}
	// Check name format using regex
	nameRegex := regexp.MustCompile(`^[A-Z][a-zA-Z]*$`)
	if !nameRegex.MatchString(name) {
		return errors.New("tên không đúng định dạng")
	}
	ob.name = name
	return nil
}

// setBirthdayYear
func (ob *Person) setBirthdayYear(num int) error {
	if num > 1900 {
		ob.birthdayYear = num
		ob.age = time.Now().Year() - num
	} else {
		return errors.New("lỗi năm, đọc lại yêu cầu")
	}
	return nil
}

// setEmail
func (ob *Person) setEmail(email string) error {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return errors.New("lỗi email, định dạng")
	}
	ob.email = email
	return nil
}

// setPhone
func (ob *Person) setPhone(phone interface{}) error {
	// gán typePhone là kiểu đầu vào của phone
	switch phone := phone.(type) {
	case string:
		// strings.HasPerfix trả về bool khi bắt đầu chuỗi là đúng
		if !strings.HasPrefix(phone, "+") || len(phone) != 12 && len(phone) != 11 {
			return errors.New("lỗi SĐT, đọc lại yêu cầu")
		} else {
			ob.phone = phone
		}

	case int:
		//var myString = strconv.Itoa(myInt)
		phoneStr := strconv.Itoa(phone)
		if !strings.HasPrefix(phoneStr, "0") || (len(phoneStr) != 11 && len(phoneStr) != 10) {
			return errors.New("lỗi SĐT, đọc lại yêu cầu")
		}
		ob.phone = phoneStr

	default:
		return errors.New("Lỗi SĐT")
	}
	return nil
}

func main() {
	ob := new(Person)
	// Set name
	err := ob.setName("Hung")
	if err != nil {
		fmt.Println("Lỗi tên:", err)
	}
	// Set birthdayYear
	err = ob.setBirthdayYear(1883)
	if err != nil {
		fmt.Println("Lỗi năm:", err)
	}
	// Set email
	err = ob.setEmail("hung@asdf.com")
	if err != nil {
		fmt.Println("Lỗi email:", err)
	}
	// Set phone
	err = ob.setPhone("1234567890")
	if err != nil {
		fmt.Println("Lỗi số điệnh thoại:", err)
	}

	// In thông tin
	fmt.Println("Name:", ob.name)
	fmt.Println("Birthday Year:", ob.birthdayYear)
	fmt.Println("Age:", ob.age)
	fmt.Println("Email:", ob.email)
	fmt.Println("Phone:", ob.phone)
}
