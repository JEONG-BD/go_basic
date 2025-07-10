package main

import "fmt"

// 변수 선언의 다양한 형태

func main() {
	var a int = 3 // 기본 형태
	var b int     // 초기값 생략
	var c = 4     // 타입 생략 변수 타입은 우변 값의 타입이 된다.
	d := 5        // 선언 대입문

	fmt.Println(a, b, c, d)

}
