package main

import "fmt"

// 포인터
// 포인터를 사용하면 data변수의 주소값만 복사되기 때문에 메모리 주소값이 8byte만 복사된다.

type Data struct {
	value int
	data  [200]int
}

func ChangeData(arg Data) {
	arg.value = 999
	arg.data[100] = 999
}

func ChangePointData(arg *Data) {
	arg.value = 999
	arg.data[100] = 999
}

func main() {
	var data Data

	ChangeData(data)
	ChangePointData(&data)

	fmt.Printf("data.value: %v\n", data.value)
	fmt.Printf("data.data: %v\n", data.data)
}
