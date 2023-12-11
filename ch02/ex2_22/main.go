package main

import "fmt"

type User struct {
	Name string
	Age  int
}

// 값 리시버를 이용한 구조체 변화
func (user User) incrementAge() {
	user.Age++
	fmt.Println(user.Age)
}

func main() {
	kathy := User{"Kathy", 19}
	kathy.incrementAge()
	fmt.Println(kathy.Age)
}
