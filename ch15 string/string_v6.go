package main

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

func NoStringBuilder() {
	var str string = "Hello"
	stringheader := (*reflect.StringHeader)(unsafe.Pointer(&str))
	addr1 := stringheader.Data

	str += " World"
	addr2 := stringheader.Data

	str += "Welcone"
	addr3 := stringheader.Data

	fmt.Println(str)
	fmt.Printf("addr1 : \t%x\n", addr1)
	fmt.Printf("addr2 : \t%x\n", addr2)
	fmt.Printf("addr3 : \t%x\n", addr3)
}
func StringBuilder() {
	var builder strings.Builder
	builder.WriteString("Hello")
	builder.WriteString("World")
	builder.WriteString("Welcome")
	fmt.Println(builder.String())
}

func main() {

	NoStringBuilder()
	StringBuilder()
}
