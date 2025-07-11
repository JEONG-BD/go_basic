package main

import "fmt"

// 구조체 선언

type House struct {
	Address string
	Size    int
	Price   float64
	Type    string
}

func main() {

	var house House = House{"seoul", 20, 20.1, "apartment"}
	fmt.Printf("house: %v\n", house)
	var house1 House = House{Size: 10}
	fmt.Printf("house1: %v\n", house1)
}
