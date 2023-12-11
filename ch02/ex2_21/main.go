package main

import "fmt"

type User struct {
	Name string
	Age  int
}

// 함수에 포인터로 전달된 구조체의 변화
func incrementAge(user *User) {
	user.Age++
	fmt.Println(user.Age)
}

func main() {
	kathy := User{"Kathy", 19}
	incrementAge(&kathy)
	fmt.Println(kathy.Age)
}
