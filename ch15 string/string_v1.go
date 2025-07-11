package main

import "fmt"

func main() {
	str1 := "Hello\t 'World'"
	fmt.Printf("str1: %v\n", str1)

	str2 := `Go is "awesome"! \n Go is Simple and \t'powerful'`
	fmt.Printf("str2: %v\n", str2)

	poit1 := "죽는 날까지 하늘을 우러러\n 한 점 부끄럼이 없기를 \n 잎새에 이는 바람에도 \n나는 괴로웠다"
	poit2 := `죽는 날까지 하늘을 우러러 
	한 점 부끄럼이 없기를 
	잎새에 이는 바람에도 
	나는 괴로웠다. `
	fmt.Printf("poit1: %v\n", poit1)
	fmt.Printf("poit2: %v\n", poit2)
}
