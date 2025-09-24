package main
import (
	"fmt"
)
func main() {
	//pointers
	//pointers are variables that store the address of other variables
	fmt.Println("pointers")
	//basic operations in pointer
	//store the address: &
	//dereferencing: * access the values from  the address pointed
	var a =20
	var ip *int
	ip=&a
	fmt.Println("the address of a",&a)
	fmt.Println("the value of a",a)
	fmt.Println("the address of a using pointer",ip)
	fmt.Println("the value of a using pointer",*ip)
	//nil pointer
	var p *int
	fmt.Println("the value of p is",p)
}