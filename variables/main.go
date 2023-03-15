package main

import (
	"fmt"
	"strconv" // Line 81
	"unsafe"  // Line 48
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

	// PTINTF ADD MESSAGE WITH CUSTOM MESSAGE
	fmt.Println()
	var my8BitsIntVar int8
	fmt.Printf("Int default value: %d\n", my8BitsIntVar)

	// WITH UNSAFE PACKAGE WE CAN SEE MEMORY USED BY VARIABLE
	fmt.Printf("Tpe: %T, bytes: %d, bits: %d\n", my8BitsIntVar, unsafe.Sizeof(my8BitsIntVar), unsafe.Sizeof(my8BitsIntVar)*8)

	var my16BitsIntVar int16
	fmt.Printf("Tpe: %T, bytes: %d, bits: %d\n", my16BitsIntVar, unsafe.Sizeof(my16BitsIntVar), unsafe.Sizeof(my16BitsIntVar)*8)

	var my32BitsIntVar int32
	fmt.Printf("Tpe: %T, bytes: %d, bits: %d\n", my32BitsIntVar, unsafe.Sizeof(my32BitsIntVar), unsafe.Sizeof(my32BitsIntVar)*8)

	var my64BitsIntVar int64
	fmt.Printf("Tpe: %T, bytes: %d, bits: %d\n", my64BitsIntVar, unsafe.Sizeof(my64BitsIntVar), unsafe.Sizeof(my64BitsIntVar)*8)

	// YOU CAN SET STRINGS WITH LINE SPACES AND ITS RESPECTED WITH YOU PRINT
	myStringVar3 := `My string
	is 
	alive`
	fmt.Println(myStringVar3)

	// SCOPES -> ALL THAT YOU SET INSIDE, IS ONLY AVAILABLE INSIDE OF DE { }
	// CHANGE FORM INTEGER OR FLOAT TO STRING
	{
		fmt.Println()
		floatVar := 33.11
		fmt.Printf("Type: %T, value: %.2f\n", floatVar, floatVar)
		floatStrVar := fmt.Sprintf("%.2f", floatVar)
		fmt.Printf("Tpye: %T, value: %s\n", floatStrVar, floatStrVar)
		fmt.Println()
		intVar := 33
		fmt.Printf("Type: %T, value: %d\n", intVar, intVar)
		intStrVar := fmt.Sprintf("%d", intVar)
		fmt.Printf("Tpye: %T, value: %s\n", intStrVar, intStrVar)
	}

	// CHANGE FROM STRING TO INTEGER OR FLOAT -> WE NEED STRCONV PACKAGE
	{
		fmt.Println()
		intVal, err := strconv.ParseInt("133", 0, 64)
		fmt.Println(err)
		fmt.Printf("Tpye: %T, value: %d\n", intVal, intVal)

		// WIH INVALID VALUE
		intVal2, err := strconv.ParseInt("fdsa133", 0, 64)
		fmt.Println(err)
		fmt.Printf("Tpye: %T, value: %d\n", intVal2, intVal2)

		fmt.Println()
		// FOR UNUSED VARIABLES
		floatVar, _ := strconv.ParseInt("133.53", 0, 64)
		fmt.Printf("Tpye: %T, value: %d\n", floatVar, floatVar)
	}

	// BINARY VARIABLES
	{
		var A byte = 'A'
		var a byte = 'a'
		var R byte = 82
		var s byte = 115
		vector := []byte{65, 97, 82, 115}

		fmt.Println()
		fmt.Println("VALUE IN ASCII CODE:")
		fmt.Println(A)
		fmt.Println(a)
		fmt.Println(R)
		fmt.Println(s)
		fmt.Println(vector)

		fmt.Println()
		fmt.Println("VALUE IN STRING:")
		fmt.Println(string(A))
		fmt.Println(string(a))
		fmt.Println(string(R))
		fmt.Println(string(s))
		fmt.Println(string(vector))
	}
}
