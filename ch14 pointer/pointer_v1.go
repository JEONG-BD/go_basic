package main

import "fmt"

//포인터 변수 선언

func main() {
	var a int
	var p *int

	p = &a // a의 메모리 주소를 포인터 변수 p에 대입한다.
	fmt.Printf("p: %v\n", p)

	var b int = 500
	var p2 *int

	p2 = &b
	fmt.Printf("p2: %v\n", p2)
	fmt.Printf("p2 value: %v\n", *p2)
	*p2 = 200
	fmt.Printf("p2: %v\n", p2)
	fmt.Printf("p2 value: %v\n", *p2)

}
