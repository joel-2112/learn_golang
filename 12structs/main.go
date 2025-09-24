package main
import (
	"fmt"
)
//structs
//defining the str
//keyword type and struct
func main() {
type Person struct {
	personId int
	name string
	age  int
	height float64
}
//accessing struct fields
//using dot notation
person1:= Person{
	personId:111,
	name: "Eyuel", 
	age: 26, 
	height: 5.5,
}

	fmt.Println("Person ID:", person1.personId)
	fmt.Println("Name:", person1.name)
	fmt.Println("Age:", person1.age)
	fmt.Println("Height:", person1.height)

}