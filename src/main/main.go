package main

import (
	"fmt"
	"person"
)

func main() {
	age := 26
	name := "Ondřej"

	ondrej := person.New(name, age)

	info := ondrej.GetInfo()

	fmt.Printf("Person info: %s\n", info)
}
