package main

import "fmt"

type contactInfo struct {
	email string
	code  int
}

// This is the definition of what a person is.
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	bruce := person{
		firstName: "Bruce",
		lastName:  "McClure",
		contact: contactInfo{
			email: "bruce@gmail.com",
			code:  2088,
		},
	}
	fmt.Printf("%+v", bruce)
}
