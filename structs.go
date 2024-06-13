package main

import (
	"fmt"
)

type Address struct {
	house_number int
	street_name  string
	city         string
	state        string
	zip          string
}

type Person struct {
	name    string
	address Address
}

func (p Person) printPerson() {
	fmt.Println(p.name)
	fmt.Println(p.address)
}

func printPerson2(p Person) {
	fmt.Println(p.name)
	fmt.Println(p.address)
}

func main() {

	var p = new(Person)

	p.name = "person"
	p.address.house_number = 101
	p.address.street_name = "this street"
	p.address.city = "city"
	p.address.state = "state"
	p.address.zip = "11111"

	p.printPerson()

	printPerson2(*p)

	p2 := Person{
		name: "person2",
		address: Address{
			house_number: 102,
			street_name:  "new street",
			city:         "new city",
			state:        "New state",
			zip:          "new zip",
		},
	}

	p2.printPerson()

}
