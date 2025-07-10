package main

import "fmt"

func HasRichFriend() bool {
	return true
}

func GetFriendsCount() int {
	return 3
}

func main() {
	price := 25000

	if price > 50000 {
		if HasRichFriend() {
			fmt.Println("신발끈이 풀렸다")
		} else {
			fmt.Println("각자 계산")
		}
	} else if price >= 30000 && price <= 50000 {
		if GetFriendsCount() > 3 {
			fmt.Println("신발끈이 풀렸다")
		} else {
			fmt.Println("각자 계산")
		}
	} else {
		fmt.Println("오늘은 내가 계산")
	}

}
