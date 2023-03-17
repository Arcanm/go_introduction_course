package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var myIntVar int
	fmt.Println(myIntVar)
	fmt.Printf("Type: %T, Bytes: %d, Bits: %d\n", myIntVar, unsafe.Sizeof(myIntVar), unsafe.Sizeof(myIntVar)*8)

	fmt.Println("----- ARRAYS -----")
	// DEFINE ARRAY -> HAVE STATIC SIZE
	var myArray [6]int
	fmt.Println(myArray)

	// DEFINE ARRAY WITH VALUES
	mySecondArray := [3]string{"1", "2", "3"}
	fmt.Println(mySecondArray)

	// CHANGE ARRAY VALUES
	myArray[1] = 1
	myArray[2] = 2
	myArray[3] = 3
	fmt.Println(myArray)
	fmt.Printf("Array Size: %d", len(myArray))
	fmt.Printf("Type: %T, Bytes: %d, Bits: %d\n", myArray, unsafe.Sizeof(myArray), unsafe.Sizeof(myArray)*8)

	fmt.Println("----- SLICES -----")
	// SLICE -> HAVE DYNAMIC SIZE
	var slice1 []int
	fmt.Printf("Slice Size: %d, slice values %v\n", len(slice1), slice1)

	// ADD VALUES TO SLICE
	slice1 = append(slice1, 10, 20, 30, 40, 50)
	fmt.Printf("Slice Size: %d, slice values %v\n", len(slice1), slice1)

	// OTHER WAY TO INSTANCE SLICE
	slice2 := make([]int, 5)
	fmt.Println(slice2)
	fmt.Printf("Slice Size: %d, slice values %v\n", len(slice2), slice2)

	// Obtain slice with array range -> IT DOESNT TAKE THE LAST VALUE OF RANGE
	// RANGES: YOU CAN USE:
	// myArray[:4] -> From index 0 to 4
	// myArray[1:] -> From index 1 to last

	fmt.Println(myArray)
	fmt.Println(myArray[:4])
	fmt.Println(myArray[2:])
	mySlice := myArray[2:4]
	fmt.Println(mySlice)
	fmt.Println("Size: ", len(mySlice))

	// MEMORY DIRECTIONS
	// SLICES: USE THE SAME MEMORY DIRECTION OF THE PARENT VALUE EXAMPLE
	fmt.Println(&myArray[2])
	fmt.Println(&mySlice[0])

	// IF YOU CHANGE THE SLICE VALUE, IT CHANGE PARENT VALUE SAME, EXAMPLE:
	mySlice[0] = 90
	mySlice[1] = 90
	fmt.Println(myArray)
	fmt.Println(mySlice)

}
