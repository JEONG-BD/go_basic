package main

import "fmt"

type ColorType int

const (
	Red ColorType = iota
	Blue
	Green
	Yellow
)

func colorToString(color ColorType) (result string) {
	switch color {
	case Red:
		result = "Red"
		fmt.Println("Red")
	case Blue:
		result = "Blue"
		fmt.Println("Blue")
	case Green:
		result = "Green"
		fmt.Println("Green")
	default:
		result = "Undefined"
		fmt.Println("Undefined")
	}
	return

}

func getMyFavoriteColor() ColorType {
	return Blue
}

func main() {
	fmt.Println("My Favorite", colorToString(getMyFavoriteColor()))
}
