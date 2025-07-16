package main

import "fmt"

func main() {
	var slice = make([]int, 3)
	var slice2 = make([]int, 3, 5)
	fmt.Println(slice)
	fmt.Println(slice2)
}
