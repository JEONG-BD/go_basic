package main

import "fmt"

func CheckAge(age int) (result bool) {

	if 10 <= age && age <= 15 {
		fmt.Println("You are young")
		result = false
	} else if 30 <= age || age <= 20 {
		fmt.Println("You are not 20s")
		result = false
	} else {
		fmt.Println("Best age of your life")
		result = true
	}
	return
}

func main() {
	age := 20
	result := CheckAge(age)
	fmt.Println(result)
}
