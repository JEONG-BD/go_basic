package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// 문자열 대입하기

func stringInformation(str1 string, str2 string) {
	var p1 *string = &str1
	var p2 *string = &str2
	fmt.Printf("str1: %v\n", p1)
	fmt.Printf("str2: %v\n", p2)
	fmt.Println(*p1)
	fmt.Println(*p2)
}

func stringHeader(str1 string, str2 string) {
	stringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str1))
	stringHeader2 := (*reflect.StringHeader)(unsafe.Pointer(&str2))

	fmt.Printf("stringHeader1: %v\n", stringHeader1)
	fmt.Printf("stringHeader2: %v\n", stringHeader2)
}

func main() {
	str1 := "안녕하세요 한글 문자열입니다"
	str2 := str1
	stringInformation(str1, str2)
	stringHeader(str1, str2)
}
