package main

import "fmt"

func main() {
	var a int = 10
	//default type for floating-point numbers is float64
	var b = 34567839023444
	// demostrate the float32 and float64 types
	var LP float32 = 34567839023444
	var HP float64 = 34567839023444
	fmt.Println("Low Precision:", LP)
	fmt.Println("High Precision:", HP)
	fmt.Println("Default Float:", b)
	//default type for string is string
	var name string = "Eyuel"
	var d bool = true
	var e complex64 = 1 + 2i
	//demonstrater complex numbers
	var cn1 complex128 = complex(5, 7)
	var cn2 complex128 = 3 + 4i
	var cnSum complex128 = cn1 + cn2
	fmt.Println("Complex Number 1:", cn1)
	fmt.Println("Complex Number 2:", cn2)
	fmt.Println("Complex Number:", e)
	fmt.Println("Sum of Complex Numbers:", cnSum)

	fmt.Println("Integer:", a)

	fmt.Println("My name is:", name)
	fmt.Println("Boolean:", d)
}
