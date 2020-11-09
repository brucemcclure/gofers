package main

import "fmt"

func main() {
	holder := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, e := range holder {
		if e%2 == 0 {
			fmt.Println(e, "is even")
		} else {
			fmt.Println(e, "is odd")
		}
	}

}
