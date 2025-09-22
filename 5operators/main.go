package main

import (
	"fmt"
)

// operators are special symbols that perform operations on variables and values
// Arithmentic operators
// +, -, *, /, %
// Relational operators
// ==, !=, >, <, >=, <=
// Logical operators
// &&, ||, !
// Assignment operators
// =, +=, -=, *=, /=, %=
// Bitwise operators
// &, |, ^, <<, >>
// Other operators
// &, *, <-, ...
func main() {
	//Arithmentic operators
	// +, -, *, /, %
	fmt.Println("Arithmentic Operators:")
	a := 10
	b := 5
	fmt.Println("a + b =", a+b)
	fmt.Println("a - b =", a-b)
	fmt.Println("a * b =", a*b)
	fmt.Println("a / b =", a/b)
	fmt.Println("a % b =", a%b)
	a++
	fmt.Println("a++ =", a)
	a--
	fmt.Println("a-- =", a)
	//Relational operators
	fmt.Println("Relational Operators:")
	fmt.Println("a == b is", a == b)
	fmt.Println("a != b is", a != b)
	fmt.Println("a > b is", a > b)
	fmt.Println("a < b is", a < b)
	fmt.Println("a >= b is", a >= b)
	fmt.Println("a <= b is", a <= b)
	//Logical operators
	fmt.Println("Logical Operators:")
	fmt.Println("a > 0 && b > 0 is", a > 0 && b > 0)
	fmt.Println("a > 0 || b < 0 is", a > 0 || b < 0)
	fmt.Println("!(a > b) is", !(a > b))
	//Assignment operators
	fmt.Println("Assignment Operators:")
	c := 10
	fmt.Println("c =", c)
	c += 5
	fmt.Println("c += 5 is", c)
	c -= 3
	fmt.Println("c -= 3 is", c)
	c *= 2
	fmt.Println("c *= 2 is", c)
	c /= 4
	fmt.Println("c /= 4 is", c)
	c %= 3
	fmt.Println("c %= 3 is", c)
	//Bitwise operators
	fmt.Println("Bitwise Operators:")
	x := 4 // 0100 in binary
	y := 5 // 0101 in binary
	fmt.Println("x & y is", x&y)
	fmt.Println("x |  y is", x|y)
	fmt.Println("x ^ y is", x^y)
	fmt.Println("x << 1 is", x<<1)
	fmt.Println("y >> 1 is", y>>1)

}
