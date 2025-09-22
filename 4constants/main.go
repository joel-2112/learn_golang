package main

import (
	"fmt"
)

//constants are immutable values, and cannot be changed once declared
//by convention constants are named using uppercase letters

func main() {

	const PI = 3.14
	const LANGUAGE = "GoLang"
	const VERSION = 1.20
	fmt.Println("the value of PI is:", PI)
	fmt.Println("the value of LANGUAGE is:", LANGUAGE)
	fmt.Println("the value of VERSION is:", VERSION)
}
