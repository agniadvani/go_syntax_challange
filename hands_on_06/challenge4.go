//create a type CIRCLE
//create a type SQUARE
//attach a method to each that calculates AREA and returns it
//circle area= π r 2
//square area = L * W
//create a type SHAPE that defines an interface as anything that has the AREA method
//create a func INFO which takes type shape and then prints the area
//create a value of type square
//create a value of type circle
//use func info to print the area of square
//use func info to print the area of circle

package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}
type square struct {
	side float64
}
type shape interface {
	area() float64
}

func main() {
	s := square{
		side: 10,
	}
	c := circle{
		radius: 22,
	}
	info(s)
	info(c)

}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}
func (b square) area() float64 {
	return b.side * b.side
}
func info(a shape) {
	fmt.Println(a.area())
}
