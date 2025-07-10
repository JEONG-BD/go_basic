package main

import "fmt"

func main() {
	a := 3
	var b float64 = 3.14

	var c int = int(b) // float64 -> int
	d := float64(a * int(b))

	fmt.Println(c, d)

	var e int64 = 7
	f := int64(d) * e
	fmt.Println(f)

}
