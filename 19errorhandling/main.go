package main

import (
	"errors"
	"fmt"
	"math"
)
type error interface {
	Error() string
}
//swuare root of the function

func Sqrt(value float64) (float64, error) {
	if value < 0 {
		return 0, errors.New("negative number is passed to sqrt")
	}
	return math.Sqrt(value), nil
}

func main() {
	fmt.Println("error handling in go")
	result, err := Sqrt(16)
	//!nill means if there is error
	//not nill means if there is error
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}
