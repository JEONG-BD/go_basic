package main

import "fmt"

/*
* Go 에서는 if 문을 검사하기 전에 초기문을 넣을 수 있다
* 초기문은 검사에 사용할 변수를 초기화할 때 사용한다.
 */

func getMyAge() (int, bool) {
	return 33, true
}

func main() {
	if age, ok := getMyAge(); ok && age < 20 {
		fmt.Println("You are young", age)
	} else if ok && age < 30 {
		fmt.Println("Nice age", age)
	} else if ok {
		fmt.Println("You are beautiful", age)
	} else {
		fmt.Println("Error")
	}
	//Error
	//fmt.Println("You age is ", age)
}
