package main

import "fmt"

func Divide(a, b int) (result int, flag bool) {
	if b == 0 {
		result = 0
		flag = false
		return
	}
	result = a / b
	flag = true
	return

}

func main() {
	c, success := Divide(9, 3)
	fmt.Println(c, success)
}
