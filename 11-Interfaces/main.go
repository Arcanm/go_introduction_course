package main

import "fmt"

// INTERFACE
type figures2D interface {
	area() float64
}

type square struct {
	width float64
}

type rectangle struct {
	width  float64
	height float64
}

func (square square) area() float64 {
	return square.width * square.width
}

func (rectangle rectangle) area() float64 {
	return rectangle.height * rectangle.width
}

func calculate(figure figures2D) {
	fmt.Println("Area: ", figure.area())
}

func main() {
	mySquare := square{width: 2}
	myRectange := rectangle{height: 2, width: 4}

	calculate(mySquare)
	calculate(myRectange)

	// Interfaces List - Simulate Array of different data types
	myInterface := []interface{}{"Hola", 12, 4.1, myRectange, mySquare}

	fmt.Println(myInterface...)
}
