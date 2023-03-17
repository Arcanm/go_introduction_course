package main

import (
	"fmt"
)

func main() {
	array := [7]int{4, 2, 5, 6, 7, 12, 34}

	for index, value := range array {
		array[index] = value + 20
	}
	fmt.Println("Los valores del array son: ", array)
}
