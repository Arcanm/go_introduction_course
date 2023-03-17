package main

import (
	"fmt"
)

func main() {
	yearsOld := 27

	fmt.Println("----- Operators -----")
	fmt.Println(yearsOld > 30)  // false
	fmt.Println(yearsOld < 30)  // true
	fmt.Println(yearsOld >= 28) // false
	fmt.Println(yearsOld <= 27) // true
	fmt.Println(yearsOld == 27) // false

	// OR
	fmt.Println()
	fmt.Println("--- OR ---")
	fmt.Println(yearsOld >= 27 || yearsOld == 53) // (true or false) = true
	fmt.Println(yearsOld > 32 || yearsOld < 27)   // (false or false) = false
	fmt.Println(yearsOld > 32 || yearsOld == 27)  // (false or true) = true

	// AND
	fmt.Println()
	fmt.Println("--- AND--- ")
	fmt.Println(yearsOld >= 27 && yearsOld == 53) // (true or false) = false
	fmt.Println(yearsOld > 32 && yearsOld < 27)   // (false or false) = false
	fmt.Println(yearsOld > 32 && yearsOld == 27)  // (false or true) = false
	fmt.Println(yearsOld >= 20 && yearsOld == 27) // (ture or true) = true

	// NOT
	fmt.Println()
	fmt.Println("--- NOT ---")
	fmt.Println(true)
	fmt.Println(!true)

	// IF / ELSE
	fmt.Println()
	fmt.Println("----- CONDITIONALS -----")
	yearsOld = 19

	fmt.Println("--- IF/ELSE ---")
	if yearsOld > 20 {
		fmt.Printf("%d is higher than 20\n", yearsOld)
	} else if yearsOld == 19 {
		fmt.Println("Tienes 19")
	} else {
		fmt.Printf("%d is lower than 20\n", yearsOld)
	}

	// ASSIGN AND VALIDATE VARIABLE EXAMPLES -> VARIABLE ONLY EXISTS INSIDE IF
	fmt.Println("-- Assign and validate varibles --")
	if value := "Hola"; value == "Hola" {
		fmt.Println("Your word is Hola")
	}

	// WITH BOOLEANS
	if value := true; value {
		fmt.Println("Is true")
	}

	// SWITCH
	fmt.Println()
	fmt.Println("--- SWITCH ---")

	{
		myInt := 3
		switch myInt {
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("One")
		case 3:
			fmt.Println("One")
		default:
			fmt.Println("Invalid")
		}
	}

	// Declaring variable in switch instance
	{
		switch myInt := 5; myInt {
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("One")
		case 3:
			fmt.Println("One")
		default:
			fmt.Println("Invalid")
		}
	}

	// SWITH WITH CONDITIONALS
	{
		myInt := 10
		switch {
		case myInt < 0:
			fmt.Println("Is lower than 0")
		case myInt > 0:
			fmt.Println("Is higher than 0")
		default:
			fmt.Println("Invalid")
		}
	}
}
