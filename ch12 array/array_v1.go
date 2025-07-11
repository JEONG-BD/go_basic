package main

import "fmt"

var scores [5]float64

func main() {
	var t [5]float64 = [5]float64{24.0, 25.9, 26.9, 29.2, 22.3}
	for i := 0; i < 5; i++ {
		fmt.Println(t[i])
	}

	fmt.Println(len(t))
	fmt.Printf("t: %v\n", t)

}
