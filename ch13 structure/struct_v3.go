package main

import "fmt"

// 구조체를 포함하는 구조체

type User struct {
	Name string
	ID   string
	Age  int
}

type VIPUser struct {
	UserInfo User
	VIPLevel int
	Price    int
}

func main() {
	user := User{"철수", "cheolsu", 20}
	vip := VIPUser{User{"철수2", "cheolsu", 30}, 3, 250}

	vip2 := VIPUser{user, 3, 250}
	fmt.Printf("vip: %v\n", vip)
	fmt.Printf("vip2: %v\n", vip2)
	fmt.Printf("vip2.UserInfo.Name: %v\n", vip2.UserInfo.Name)
}
