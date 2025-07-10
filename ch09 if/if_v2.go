package main

import "fmt"

func ControlTemperature(temp int) {

	if temp > 28 {
		fmt.Println("에어컨을 켠다")
	} else if temp < 3 {
		fmt.Println("히터를 켠다")
	} else {
		fmt.Println("아무 것도 하지 않는다")
	}

}

func main() {
	temp := 33

	ControlTemperature(temp)

}
