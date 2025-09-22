package main
import "fmt"

func main() {
	var a int = 10
	//used for less precise calculations
	var ft32 float32 = 20.5
	//used for more precise calculations
	var ft64 float64 = 20.5
	//default type for floating-point numbers is float64
	var b = 20.5
	// demostrate the float32 and float64 types
	var LP float32 = 34567839023444
	var HP float64 = 34567839023444
	fmt.Println("Low Precision:", LP)
	fmt.Println("High Precision:", HP)
	var c string = "Hello, Go!"
	var d bool = true
	var e complex64 = 1 + 2i
	fmt.Println("Integer:", a)
	fmt.Println("Float32:", ft32)
	fmt.Println("Float64:", ft64)

	fmt.Println("Default Float:", b)
	fmt.Println("String:", c)
	fmt.Println("Boolean:", d)
	fmt.Println("Complex Number:", e)
}