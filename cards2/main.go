package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	pk := person{
		firstName: "Punith",
		lastName:  "Krishna",
		contact: contactInfo{
			email:   "test@gmail.com",
			zipCode: 90000,
		},
	}

	// we can either create a pointer or directly make the call with the struct pk
	// if the struct object is used directly and we have a functions defined with a pointer reciever
	// Go will directly pass the address of struct object to that handler.
	var pkPointer *person = &pk
	pkPointer.updateName("P Punith")
	pk.print()

}

func (pkPtr *person) updateName(newName string) {
	(*pkPtr).firstName = newName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
