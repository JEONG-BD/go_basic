package main

import "fmt"

/*
* 하나의 case는 히니 이상의 값을 비교 가능하다
 */

func MultiSwitch(day string) (result bool) {
	switch day {
	case "monday", "tuesday":
		fmt.Println("학교 가는 날")
		result = false
	case "wednesday", "thursday", "friday":
		fmt.Println("실습 가는 날")
		result = false
	default:
		fmt.Println("쉬는 날")
		result = true
	}
	return

}

func main() {
	day := "thurday"
	result := MultiSwitch(day)
	fmt.Println("result", result)
}
