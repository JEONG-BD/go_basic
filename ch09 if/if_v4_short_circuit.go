package main

import "fmt"

/*
* short circuit
* && 연산은 좌변이 false 면 우변을 검사하지 않고 false 처리
* || 연산은 좌변이 true면 우변을 검사하지 않고 true 처리를 한다
* 이를 short circuit 라고 한다
 */

var cnt int = 0

func IncreaseAndReturn() int {
	fmt.Println("IncreaseAndReturn()", cnt)
	cnt++
	return cnt
}

func main() {
	// 좌변이 false 이기 때문에 우변 호출하지 않음
	if false && IncreaseAndReturn() < 5 {
		fmt.Println("1 증가")
	}
	// 좌변 true 이기 때문에 우변 호출하지 않음
	if true || IncreaseAndReturn() < 5 {
		fmt.Println("2 증가")
	}

	fmt.Println(cnt)
}
