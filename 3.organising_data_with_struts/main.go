package main

import "fmt"

// This is the definition of what a person is.
type person struct {
	firstName string
	lastName  string
}

func main() {
	// bruce := person{"Bruce", "McClure"} // This relys 100% on the order for assignemnt
	bruce := person{firstName: "Bruce", lastName: "McClure"} // This doesn't rely on order
	fmt.Println(bruce)
}
