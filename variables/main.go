package main

import (
	"fmt"
)

func main() {
	var myIntVar int
	myIntVar = -12
	fmt.Println("Numero: ", myIntVar)

	// Solo Enteros positivos
	var myUintVar uint
	myUintVar = 12
	fmt.Println("Positive integer: ", myUintVar)

	var myStringVar string
	myStringVar = "Hola Mundo"
	fmt.Println("My variable is: ", myStringVar)

	var myBoolVar bool
	myBoolVar = true
	fmt.Println("My variable is: ", myBoolVar)

	// Obtain memory address of variable
	fmt.Println("My variable memory address is: ", &myStringVar)

	// Set automatically variable type
	myIntVar2 := 12
	fmt.Println("My variable is: ", myIntVar2)

	myStringVar2 := "Testing"
	fmt.Println("My variable is: ", myStringVar2)

	//  CONSTANTS, YOU CANNOT CHANGE CONSTANT VALUES
	const myFirstConst string = "First constant"
	const mySecondConst = 12

	fmt.Println("My constants are: ", myFirstConst, mySecondConst)
}
