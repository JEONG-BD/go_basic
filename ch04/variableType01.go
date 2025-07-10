package main

import "fmt"

func main() {

	/*
		프로그래밍 언어를 구분할 대 타입 검사를 하는가 안 하는가에 따라 강타입 언어와 약 타입 언어로 구분
		Go는 가장 강하게 타입 검사를 하는 강타입 언어
	*/

	a := 3
	var b float64 = 3.5

	var c int = b // Error
	d := a * b    //Error
	fmt.Println(c)
	fmt.Println(d)

	var e int64 = 7
	f := a * e

}
