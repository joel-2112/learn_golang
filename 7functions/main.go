package main

import (
	"fmt"
)

func saycountry(s string) {
	fmt.Printf("%s, is a country\n", s)
}
func saycity(s string) {
	fmt.Printf("%s, is a city\n", s)
}
func maxoftwo(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func swap(firstName, lastName string) (string, string) {
	return lastName, firstName
}
func main() {
	fmt.Println("funtions")
	saycountry("Ethiopia")
	saycity("Addis Ababa")
	a, b := 300, 500
	result := maxoftwo(a, b)
	fmt.Println("max is ", result)
	fmt.Printf("max of %d and %d is %d\n", a, b, result)
	//when you declare multiple returns as string use long name not single letter
	firstName, lastName  := swap("eyuel", "kassahun")
	fmt.Printf("swapped : %s %s\n", firstName, lastName)
}
 