package main

import (
	"fmt"
	"math"
)

//interface
type Shape interface {
	area() float64
}
type Circle struct{
	radius float64
}
type Rectangle struct{
	width, height float64
}

func getArea(s Shape) float64 {
	return s.area()

}

func (c Circle) area() float64{
	return math.Pi*c.radius * c.radius
}
func (r Rectangle) area()float64 {
	return r.height * r.width 
}
func main() {
	fmt.Println("interfaces")

	circle := Circle{radius:5}
	rectangle := Rectangle{width: 5, height: 6}
	fmt.Printf("the area of circle is %.2f\n", getArea(circle))
	fmt.Printf("the area of rectangle is%.2f",getArea(rectangle))
}