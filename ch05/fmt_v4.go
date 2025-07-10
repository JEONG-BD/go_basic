package main

import "fmt"

/*
* 표준 입출력 기능은 Go 언어 패키지인 fmt 에셔 제공한다.
* Go 언어는 자주 사용하는 기능을 묶어서 패키지로 제공한다.
 */

func main() {
	str := "Hello \tGo\t\tworld\n \"Go is Awesome!\n"
	fmt.Print(str)
	fmt.Printf("%s", str)
	fmt.Printf("%q", str)
}
