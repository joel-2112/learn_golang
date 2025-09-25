package main
import ("fmt")

func main(){
	fmt.Println("the ranges")
	//using range with slice and arrays
	numbers := []int {0,1,2,3,4,5,6,7,8,9}
	for i, number := range numbers{
		fmt.Printf("Slice %d is %d\n", i, number)
	}
	n := []int {1,2,3}
	for i := range n {
       fmt.Printf("Array item %d is %d\n", i, n[i])
	}
}
