package main

import "fmt"

func main() {

	var char rune = '한'
	fmt.Printf("char: %v\n", char)
	fmt.Printf("%c\n", char)
	fmt.Printf("%T\n", char)

	//len()

	str1 := "가나다라마바사"
	str2 := "abcde"

	fmt.Printf("len(str1): %v\n", len(str1))
	fmt.Printf("len(str2): %v\n", len(str2))

	//
	str3 := "Hellow World"
	runes := []rune{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100}

	fmt.Printf("str3: %v\n", str3)
	fmt.Printf("runes: %v\n", runes)
	fmt.Printf("runes: %v\n", string(runes))
}
