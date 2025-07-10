package main

import "fmt"

/*
* 표준 입출력 기능은 Go 언어 패키지인 fmt 에셔 제공한다.
* Go 언어는 자주 사용하는 기능을 묶어서 패키지로 제공한다.
 */

func main() {
	var a int = 10
	var b int = 20
	var f float64 = 32799438743.8297

	fmt.Print(a, b, f)
	fmt.Println("a", a, "b", b, "f", f)
	fmt.Printf("a : $d b: %d f:%f\n", a, b, f)
}
