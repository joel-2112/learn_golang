package main
import (
	"fmt"
)
//global variable
var z int
func main() {

	//local variable
	x := 30
	y := 20
	z = x + y
	fmt.Printf( "x=%d, y=%d, global var z=%d", x, y, z)
}