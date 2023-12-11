package main

import "fmt"

type User struct {
	Name string
	Age  int
}

// 구조체를 이용하여 여러개의 반환 값 변경하기
func nameAndAge(uid int) User {
	switch uid {
	case 0:
		return User{"Baheer Kamal", 24}
	case 1:
		return User{"Tamay Bakshi", 16}
	default:
		return User{"", -1}
	}
}

func main() {
	user := nameAndAge(1)
	fmt.Println(user.Name)
	fmt.Println(user.Age)
}
