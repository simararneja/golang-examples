package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName   string
	lastName    string
	contactInfo contactInfo
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Andreson"}
	// fmt.Println(alex)
	// alex := person{"Alex", "Andreson"}

	// var alex person

	// alex.firstName = "Alex"
	// alex.lastName = "Andreson"
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "party",
		contactInfo: contactInfo{
			email:   "j@g.com",
			zipCode: 678,
		},
	}
	// jimPointer := &jim
	jim.updateName("jimmy")
	jim.print()
	// fmt.Printf("%+v", jim)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
