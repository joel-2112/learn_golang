package main
import ("fmt")

func main()  {
	fmt.Println("recursions in go")
	fmt.Println("the factorial of given positive number")
	x:=factorial(5)
	fmt.Println(x)
}

//base case : stops the recursion by preventing the loop
//recursion case : function calls itself

//factorial function

func factorial(n int)int  {
	if n==0 {
		return 1
	}else{
		return  n*factorial(n-1)
	}
}