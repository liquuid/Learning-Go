package main

import "fmt"

func main() {
	//var colors map[string] string
	//colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#4BF745",
	}

	colors["white"] = "#FFFFFF"
	//delete(colors, "white")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
