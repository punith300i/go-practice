package main

import "fmt"

func main() {
	colours := map[string]string{
		"red":   "hex1",
		"green": "hex2",
	}

	// two ways to create empty map
	// using var
	// var colours1 map[string]string
	// if colours1["game"] == nil {
	// 	colours1["game"] = "red"
	// }

	// using make
	colours2 := make(map[string]string)
	colours2["game"] = "red"

	delete(colours2, "game")

	fmt.Println(colours)
	// fmt.Println(colours1)
	fmt.Println(colours2)
	printMap(colours)
}

func printMap(m map[string]string) {
	for k, v := range m {
		fmt.Println(k, v)
	}
}
