package main
import (
	"fmt"
)
func main() {
//arrays in go 
//arrays are fixed slice, in size and type
//arrays are value types
//arrays are not used much in go, slices are used more often
//arrays are used when you know the size of the array at compile time
//arrays are used when you want to avoid the overhead of slices
//arrays are used when you want to avoid the overhead of pointers
//arrays are used when you want to avoid the overhead of garbage collection

//declare an array
var arr [5]int
fmt.Println("default array: ",arr)//[0 0 0 0 0] - default value of int is 0
fmt.Println("length",len(arr)) //length of array
//declare and initialize an array
arr1 := [5]int{1,2,3,4,5} //or var arr1 = [5]int{1,2,3,4,5}
//accessing array elements
fmt.Println("the first array element",arr1[0])
fmt.Println("the whole array ",arr1)

//initialize by infering the size
arr2 := [...]int{1,2,3,4,5,6}// or var arr2 = []int{1,2,3,4,5,6}
fmt.Println("size infering array",arr2)
fmt.Println("length",len(arr2))
//initialize by array elements

var n [10]int
for i:=1; i<10; i++ {
	n[i] = i*10
	fmt.Println("the ",i,"th element is ",n[i])
}
}