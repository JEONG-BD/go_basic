package main

import "fmt"

func main() {
	nums := [...]int{10}

	nums[0] = 300
	for i := 0; i < len(nums); i++ {
		fmt.Printf("nums: %v\n", nums[i])
	}
	fmt.Printf("nums: %v\n", len(nums))

	var t [5]float64 = [5]float64{25.0, 25.3, 25.2, 25.1, 25.5}
	for i, v := range t {
		fmt.Println(i, v)
	}
}
