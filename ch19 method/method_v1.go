package main

import "fmt"

type account struct {
	balance int
}

// 사용자 정의 별칭 타입
type myInt int

// myInt 별칭 타입을 리시버로 갖는 메서드
func (a myInt) add(b int) (result int) {
	return int(a) + b
}

// 일반 함수 표현
func withdrawFunc(a *account, amount int) {
	a.balance -= amount
}

// 메서드 표현
// func 키워드와 함수명 사이에 리시버 타입과 리시버 값을 갖는 변수가 있으면 메서드 없으면 함수

func (a *account) withdrwaMethod(amount int) {
	a.balance -= amount
}

func main() {
	a := &account{100}
	withdrawFunc(a, 30)
	a.withdrwaMethod(30)
	fmt.Printf("%d \n", a.balance)

	var b myInt = myInt(10)
	fmt.Println(b.add(30))
	var c int = 20
	fmt.Println(myInt(c).add(50))
}
