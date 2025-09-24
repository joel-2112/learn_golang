package main

import (
	"fmt"
	"strings"
)
func main() {
	//strings
	var name = "hello go"
	fmt.Println("normal string",name)
	fmt.Printf("formated string : %s\n", name)
	//displaying hexadecimal type of string
	fmt.Println("hexbytes...")
	for i:=0; i<len(name); i++ {
		fmt.Printf("0x%x", name[i])
	}
	fmt.Println()

	//string concatination
	//create slice of string to concatinate
	var names = []string{"hello", "go", "lang"}
	//using join method to concatinate
	fmt.Println(strings.Join(names, " "))

}