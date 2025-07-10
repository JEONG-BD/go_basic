package main

import "fmt"

func getMyAge() int {
	return 22
}

func InitSwitch() {
	switch age := getMyAge(); true {
	case age < 10:
		fmt.Println("Child")
	case age < 20:
		fmt.Println("Teenager")
	case age < 30:
		fmt.Println("20s")
	default:
		fmt.Println("따뜻한 날씨")
	}
}

func main() {
	InitSwitch()
}
