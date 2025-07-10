package main

import "fmt"

func NoSwitch(day int) {
	if day == 1 {
		fmt.Println("첫째날 입니다.")
		fmt.Println("개발 회의 입니다.")
	} else if day == 2 {
		fmt.Println("둘째날 입니다.")
		fmt.Println("디자인 회의 입니다.")
	} else if day == 3 {
		fmt.Println("셋쨰날 입니다.")
		fmt.Println("설계 회의 입니다.")
	} else if day == 4 {
		fmt.Println("넷째날 입니다.")
		fmt.Println("고객 회의 입니다.")
	} else if day == 5 {
		fmt.Println("다섯째 입니다.")
		fmt.Println("테스트 회의 입니다.")
	} else {
		fmt.Println("프로젝트 완료.")
	}
}
func Switch(day int) {

	switch day {
	case 1:
		fmt.Println("첫째날 입니다.")
		fmt.Println("개발 회의 입니다.")
	case 2:
		fmt.Println("둘째날 입니다.")
		fmt.Println("디자인 회의 입니다.")
	case 3:
		fmt.Println("셋쨰날 입니다.")
		fmt.Println("설계 회의 입니다.")
	case 4:
		fmt.Println("넷째날 입니다.")
		fmt.Println("고객 회의 입니다.")
	case 5:
		fmt.Println("다섯째 입니다.")
		fmt.Println("테스트 회의 입니다.")
	default:
		fmt.Println("프로젝트 완료.")
	}
}

func main() {
	var a int
	fmt.Print("요일 번호를 입력하세요 (1~7): ")
	_, err := fmt.Scanln(&a)
	if err != nil {
		fmt.Println("입력 오류:", err)
		return
	}
	NoSwitch(a)
	Switch(a)
}
