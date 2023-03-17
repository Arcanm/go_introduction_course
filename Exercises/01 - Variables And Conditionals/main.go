package main

import (
	"fmt"
)

func main() {
	var license bool
	var age int8

	fmt.Println("Do you have drivers licence?")
	fmt.Scanf("%t\n", &license)
	fmt.Println("What is your age?")
	fmt.Scanf("%d\n", &age)

	switch {
	case license == false:
		fmt.Println("You can't keep going")
	case age <= 15:
		fmt.Println("You can't keep going")
	default:
		fmt.Println("Welcome, you can keep going")
	}
}
