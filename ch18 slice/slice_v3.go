package main

import "fmt"

func changeArray(array2 [5]int) {
	array2[2] = 200
}

func changeSlice(slice2 []int) {
	slice2[2] = 200
}

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}

	fmt.Println(array)
	fmt.Println(slice)

	changeArray(array)
	changeSlice(slice)

	fmt.Println(array)
	fmt.Println(slice)
}
