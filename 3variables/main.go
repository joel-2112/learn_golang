package main

import (
	"fmt"
)

func main() {
	//declare and then assign in 2 steps
	var mango string
	mango = "mango"

	//declare and assign in 1 step
	var apple string = "apple"
	//declare and assign using := shorthand
	country := "Ethiopia"
	city := "Addis Ababa"
	//output
	fmt.Println("declare and assign in 2 line mango:", mango)
	fmt.Println("declare and assign in 1 line apple:", apple)
	fmt.Println("the capital city of ", country," is ", city)
	//declare multiple variables in 1 line
	var fruit1, fruit2, fruit3 string = "mango", "apple", "banana"
	fmt.Println("declare multiple variables in 1 line",fruit1, fruit2, fruit3)
	//declare multiple variables with type inference
	var fruit4, fruit5, fruit6 = "mango", "apple", "banana"
	fmt.Println("declare multiple variables with type inference",fruit4, fruit5, fruit6)
	//declare multiple variables using :=
	fruit7, fruit8, fruit9 := "mango", "apple", "banana"
	fmt.Println("multiple variables using := ",fruit7, fruit8, fruit9)
	//declare multiple variables in 1 line with different types
	var fruit10, price, isGood = "mango", 20, true
	fmt.Println("multiple variables in 1 line multiple type",fruit10, price, isGood)
	//declare multiple variables using := with different types
	fruit11, price1, isGood1 := "mango", 20, true
	fmt.Println("multiple variables using := with different type",fruit11, price1, isGood1)
	//swap values of 2 variables
	mango, apple = apple, mango
	fmt.Println("after swapping", mango, apple)
	// use the println not fmt.Println when use \n
	var x, y int = 5, 10
	print("the value of x is: ", x, "\nthe value of y is: ", y)

	//checking the type of a variable
	fmt.Printf("\nthe type of mango is: %T", mango)
	fmt.Printf("\nthe type of price is: %T", price)
	fmt.Printf("\nthe type of isGood is: %T", isGood)


}
