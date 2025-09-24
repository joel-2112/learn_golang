package main
import ("fmt")
//slices -> dynamic type and size not specified

func main(){
	fmt.Println("slices")
	//declare slices
	numbers := make([]int,5)
	fmt.Println(numbers)
	//subslicing the last is not included
	numbers2 := []int {0,1,2,3,4,5,6,7,8,9}
	slice1 := numbers2[1:5]//1,2,3,4
	slice2 := numbers2[:5]//0,1,2,3,4
	slice3 := numbers2[5:9]//5,6,7,8
	fmt.Println("sub slicing")
	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)
	fmt.Println("slice3:", slice3)
	// append and copy
	num := []int{}
	num = append(num,1,2,3,4,5,6,7,8,9)
	num1 := make([]int, len(num), cap(num)*2)
	copy(num1,num)
	fmt.Println("num slice ",num)
	fmt.Println("num1 after copy", num1)
	//nil slice
	var x []int
	fmt.Println("nil slice: ",x)
	if x == nil {
		fmt.Println("this is nil slice",x)
	}
}