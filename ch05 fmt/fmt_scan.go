package main

import "fmt"

func main() {
	var a int // 1 값을 저장할 변수
	var b int
	n, err := fmt.Scan(&a, &b) // 2 입력 두 개 받기
	if err != nil {            // 3 에러가 발생하면 에러 코드 출력
		fmt.Println(n, err)
	} else { // 4 정상 입력되면 입력값 출력
		fmt.Println(n, a, b)
	}
}
