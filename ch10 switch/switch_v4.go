package main

import "fmt"

func CompareSwitch(temp int) {
	switch true {
	case temp < 10, temp > 30:
		fmt.Println("바같 활동하기 좋지 않은 날씨")
	case temp >= 10 && temp < 20:
		fmt.Println("약간 추울 수 있으니 가벼운 겉옷")
	case temp >= 15 && temp < 25:
		fmt.Println("야외 활동하기 좋은 날씨")
	default:
		fmt.Println("따뜻한 날씨")
	}

}

func main() {

	temp := 10
	CompareSwitch(temp)

}
