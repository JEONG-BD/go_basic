package main

import "fmt"

const (
	Red   int = iota
	Blue  int = iota
	Green int = iota
)

const (
	C1 uint = iota + 1
	C2
	C3
)

func main() {
	fmt.Println("Colors:")
	fmt.Println("Red:", Red)
	fmt.Println("Blue:", Blue)
	fmt.Println("Green:", Green)

	fmt.Println("\nConstants C:")
	fmt.Println("C1:", C1)
	fmt.Println("C2:", C2)
	fmt.Println("C3:", C3)
}
