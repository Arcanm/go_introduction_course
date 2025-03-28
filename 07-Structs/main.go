package main

import "fmt"

// DATA STRUCT, YOU CAN DEFINE ANY ATTRIBYTES OF STRUCT INSIDE
type Car struct {
	brand string
	year  int
}

func main() {
	// FIRST WAY OF INITIALIZE A STRUCT
	myFirstCar := Car{brand: "Kia", year: 2025}
	fmt.Println(myFirstCar)

	// SECOND WAY TO INITIALIZE A STRUCT,
	// IF YOU DONT ASSIGN A VALUE FOR ATTRIBUTE IT TAKES THE ZERO VALUE OF THE DATA TYPE
	var otherCar Car
	otherCar.brand = "Ford"
	fmt.Println(otherCar)
}
