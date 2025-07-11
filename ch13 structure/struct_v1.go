package main

import "fmt"

// 구조체 선언
type Student struct {
	Name  string
	Class int
	No    int
	Score float64
}

type House struct {
	Address string
	Size    int
	Price   float64
	Type    string
}

func main() {
	var student Student
	student.Name = "mike"
	student.Class = 1
	student.No = 1
	student.Score = 90.2

	fmt.Printf("student: %v\n", student)
}
