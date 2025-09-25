package main

import (
	"fmt"
	"strconv"
)
func main() {
	fmt.Println("type casting in go")
	var total int = 20
	var items int = 3
	var average float64 
	average = float64(total)/float64(items)
	fmt.Printf("the average casted to float is %.2f\n",average)

	//converting strings to number -if convertible
	str := "123"
	num, err := strconv.Atoi(str)
	if err !=nil{
		fmt.Println("error converting")
		return 
	}else{

		fmt.Println("converting str to num",num)
	}
}