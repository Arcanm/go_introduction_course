package mypackage

import (
	"fmt"
)

// PublicCar - A car with public acces
// Must be UpperCase
type PublicCar struct {
	Year  int
	Brand string
}

// privateCar - Private struct, i cant use in main.go
type privaterCar struct {
	year  int
	brand string
}

// PublicPrintMessage - Public function
func PublicPrintMessage(message string) {
	fmt.Println(message)
}

// privatePrintMessage - Public function
func privatePrintMessage(message string) {
	fmt.Println(message)
}
