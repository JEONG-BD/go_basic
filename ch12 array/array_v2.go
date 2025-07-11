package main

import "fmt"

//var nums [5] int

func main() {
	nums := [5]int{1, 2, 3, 4, 5}
	var temps [5]float64 = [5]float64{11.1, 11.2}
	fmt.Printf("nums: %v\n", nums)
	fmt.Printf("temps: %v\n", temps)

	//배열 선언시 개수는 항상 상수

	const Y int = 3
	//x := 5
	//a := [x]int{1, 2, 3, 4, 5}
	b := [Y]int{1, 2, 3}
	fmt.Println(b)
}
