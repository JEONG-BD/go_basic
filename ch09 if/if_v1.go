package main

import (
	"fmt"
)

func main() {
	light := "red"
	fmt.Println(light)
	if light == "red" {
		fmt.Println("wait")
	} else {
		fmt.Println("Walk")
	}
}
