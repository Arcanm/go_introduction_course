package main

import "fmt"

func main() {
	var array []int
	var promptValue int
	for {
		fmt.Printf("Ingresa un valor: ")
		fmt.Scanf("%d\n", &promptValue)

		if promptValue == 0 {
			break
		}
		array = append(array, promptValue)
	}
	fmt.Println(array)
}
