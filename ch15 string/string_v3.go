package main

import "fmt"

func main() {
	str := "Hello 월드"
	runes := []rune(str)

	fmt.Printf("str: %v\n", str)
	fmt.Printf("runes: %v\n", runes)

	for i := 0; i < len(str); i++ {
		fmt.Printf("T : %T i: %v\n %c", str[i], str[i], str[i])
	}

	for _, v := range str {
		fmt.Printf("%T %v %c\n", v, v, v)
	}

	str1 := "Hello"
	str2 := "World"

	str3 := str1 + str2
	fmt.Printf("str3: %v\n", str3)
}
