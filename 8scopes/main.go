package main

import (
	"fmt"
)

// global variable
var z int
var a = 100

func main() {

	//local variable
	x := 30
	y := 20
	z = x + y
	var a = 200
	fmt.Printf("local var a=%d\n", a)
	fmt.Printf("x=%d, y=%d, global var z=%d\n", x, y, z)
}
