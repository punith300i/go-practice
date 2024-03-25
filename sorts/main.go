package main

import (
	"fmt"
	"sort"
)

func main() {
	var tempSlice = [][]int{{5, 1}, {1, 2}, {3, 4}, {4, 8}, {4, 4}}
	// tempSlice2 := make([][]int, 5)
	sort.Slice(tempSlice, func(i, j int) bool {
		fmt.Println(tempSlice[i], tempSlice[j])
		return tempSlice[i][1] < tempSlice[j][1]
	})
	fmt.Println(tempSlice)
}
