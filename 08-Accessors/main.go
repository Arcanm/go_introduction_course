package main

import (
	"fmt"

	// YOU CAN ADD ALIAS TO IMPORTS, ITS NOT COMMON, EXAMPLE:
	pk "github.com/Arcanm/go_introduction_course/08-Accessors/mypackage"
)

func main() {
	// Public Struct
	// If i dont use alias in line 7, i need to import with mypackage.PublicCar
	var myCar pk.PublicCar
	myCar.Brand = "Honda"

	// If i declare in PublicCar struct "Year" like "year" it means that the attribute is private and I wouldn't be able to modify it here.
	myCar.Year = 2025

	fmt.Println(myCar)

	// Private struct
	// I cant use next lines because is a private struct
	// var privaterCar pk.privaterCar
	// fmt.Println(privaterCar)

	// Public func
	pk.PublicPrintMessage("Hello World")

	// Private func
	// I cant use next lines because is a private func
	// pk.privatePrintMessage("Hello World")
}
