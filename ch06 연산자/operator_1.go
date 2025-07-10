package main

import "fmt"

/*
* Go  언어에서 모든 연산자의 각 항의 타입은 같아야 한다(시프트 연산 예와)
* 연산의 결과 타입도 인수 타입과 같다.
 */
func main() {
	var x int32 = 7
	var y int32 = 3

	var s float32 = 3.14
	var t float32 = 5

	fmt.Println("x + y = ", x+y)
	fmt.Println("x - y = ", x-y)
	fmt.Println("x * y = ", x-y)
	fmt.Println("x / y = ", x-y)
	fmt.Println("x % y = ", x-y)

}
