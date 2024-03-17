package main

import "fmt"

func main() {
	var n int = 20
	var p *int = &n

	struc := myStruct{X: 15, Y: 20}

	var myStructPointer *myStruct = &struc

	myStructPointer.Y = 30

	fmt.Println(*p)
	fmt.Println(struc)
	fmt.Println(*myStructPointer)
}

type myStruct struct {
	X int
	Y int
}
