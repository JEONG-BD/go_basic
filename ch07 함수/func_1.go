package main

import "fmt"

// 첫글자가 대문자인 함수는 패키지 외부로 공개
func Add(a int, b int) int {

	return a + b
}
func PrintAverage(name string, math int, eng int, history int) {
	total := math + eng + history
	avg := total / 3
	fmt.Println(name, " average score 는", avg, " 입니다")
}
func main() {
	c := Add(1, 3)
	fmt.Print(c)
	fmt.Println(c)
	PrintAverage("user1", 3, 3, 4)
}
