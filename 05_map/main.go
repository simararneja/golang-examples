package main

import (
	"fmt"
)

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	printMap(colors)
	// fmt.Println(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

// func main() {
// 	// intialize with zero value
// 	// var colors map[string]string

// 	colors := make(map[string]string)

// 	colors["white"] = "#ffffff"

// 	// Delete a key from a map
// 	// delete(colors, "white")

// 	// colors := map[string]string{
// 	// 	"red":   "#ff0000",
// 	// 	"green": "#4bf745",
// 	// }

// 	fmt.Println(colors)
// }
