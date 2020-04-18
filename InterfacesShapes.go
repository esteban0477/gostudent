package main

import (
	"fmt"
	"math"
)

type square struct {
	large float64
	width float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.large * s.width
}

func (s circle) area() float64 {
	return math.Pi * s.radius * s.radius
}

type shape interface {
	area() float64
}

func main() {
	firstSquare := square{
		large: 30.0,
		width: 30.0,
	}

	firstCirle := circle{
		radius: 10.0,
	}

	info(firstSquare)
	info(firstCirle)
}

func info(form shape) {
	switch form.(type) {
	case square:
		fmt.Println("This is a square")
		fmt.Println("Large:", form.(square).large, "& the Width:", form.(square).width)
		fmt.Println("The area of this shape is:", form.(square).area())
	case circle:
		fmt.Println("This is a circle")
		fmt.Println("Radius:", form.(circle).radius)
		fmt.Println("The area of this shape is:", form.(circle).area())
	}
}
