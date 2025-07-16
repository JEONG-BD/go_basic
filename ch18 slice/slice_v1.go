package main

import "fmt"

func main() {
	//var array[]int
	var slice []int

	if len(slice) == 0 {
		fmt.Println("slice is empty", slice)
	}

	//slice[1] = 10
	//fmt.Println(slice) //panic

	var slice1 = []int{1, 2, 3}
	var slice2 = []int{1, 5: 2, 10: 3}
	fmt.Println(slice1)
	fmt.Println(slice2)

	// var array = [...]int{1, 2, 3} // 배열
	// var slice = []int{1, 2, 3}    // 슬라이스

	var slice3 = make([]int, 3)
	fmt.Println(slice3)

	// 슬라이스 요소 접근
	var slice4 = make([]int, 4)
	slice4[1] = 5
	fmt.Println(slice4)

	// 슬라이스 순회
	var slice5 = []int{1, 2, 3}

	for i := 0; i < len(slice5); i++ {
		slice5[i] += 10
	}
	fmt.Println(slice5)

	for i, v := range slice5 {
		slice5[i] = v * 2
	}
	fmt.Println(slice5)

	//슬라이스 요소 추가
	slice6 := append(slice5, 4)
	fmt.Println(slice6)
	fmt.Println(slice6)

	//슬라이스 요소 여러 요소 추가
	slice7 := append(slice5, 1, 2, 3, 4)
	fmt.Println(slice7)

}
