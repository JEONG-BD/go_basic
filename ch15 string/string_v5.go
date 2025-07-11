package main

import "fmt"

func main() {
	var str string = "Hello world"
	//var slice []byte = []byte(str)

	var p1 *string = &str
	slice := []byte(str)

	var p2 *byte = &slice[2]
	fmt.Println(*p2)

	fmt.Printf("p1: %v\n", p1)
	fmt.Printf("p2: %v\n", p2)

	slice[2] = 'a'
	fmt.Println(str)
	fmt.Printf("%s\n", slice)

}
