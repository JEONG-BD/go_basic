package main

import "fmt"

type account struct {
	balance   int
	firstName string
	lastName  string
}

func (a1 *account) withdrawPointer(amount int) {
	a1.balance -= amount
}

func (a2 account) withdrawValue(amount int) {
	a2.balance -= amount
}

func (a3 account) withdrawReturnValue(amount int) account {
	a3.balance -= amount
	return a3
}

func main() {
	var mainA *account = &account{100, "Joe", "Park"}
	mainA.withdrawPointer(30)
	fmt.Printf("mainA = %d\n", mainA.balance)

	mainA.withdrawValue(20)
	fmt.Printf("mainA = %d\n", mainA.balance)

	var mainB account = mainA.withdrawReturnValue(20)
	fmt.Printf("mainB = %d\n", mainB.balance)
}
