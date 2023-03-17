package main

import "fmt"

func main() {
	sum := 0

	// FOR
	// for [INITIAL VALUE] ; [VALIDATION] ; [INCREMENT OR DECREMENT]
	fmt.Println("----- FOR -----")

	// CLASIC FOR DECLARATION
	for i := 0; i < 10; i++ {
		sum++
	}
	fmt.Println(sum)

	// FOR ONLY WITH VALIDATION
	for sum < 1000 {
		sum++
	}
	fmt.Println(sum)

	// INFINITE FOR WITH A BREAK
	for {
		if sum > 1000 {
			break
		}
		sum++
	}
	fmt.Println(sum)

	// FOR WITH ARRAYS
	fmt.Println("----- FOR WITH ARRAYS-----")
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for index := range arr {
		fmt.Println("INDEX: ", index, " - VALUE: ", arr[index])
	}

	// WITH INDEX AND VALUE
	fmt.Println()
	for index, value := range arr {
		fmt.Println("INDEX: ", index, " - VALUE: ", value)
	}

	// ONLY WITH VALUE
	fmt.Println()
	for _, value := range arr {
		fmt.Println("VALUE: ", value)
	}

	fmt.Println("\n----- FOR WITH MAPS-----")
	mapOne := map[string]float32{
		"A": 1.10,
		"B": 1.20,
		"C": 1.30,
	}

	// ONLY WITH KEY
	fmt.Println()
	for key := range mapOne {
		fmt.Println("INDEX: ", key, " - VALUE: ", mapOne[key])
	}

	// WITH KEY AND VALUE
	fmt.Println()
	for key, value := range mapOne {
		fmt.Println("INDEX: ", key, " - VALUE: ", value)
	}

	// ONLY WITH VALUE
	fmt.Println()
	for _, value := range mapOne {
		fmt.Println("VALUE: ", value)
	}

	// FOR WITH MAP OF ARRAYS
	fmt.Println("\n----- FOR WITH MAP OF ARRAYS-----")
	mapTwo := map[string][]int{
		"A": nil,
		"B": {1, 2, 3},
		"C": {1, 2, 3},
	}

	for key, value := range mapTwo {
		fmt.Println("MAP KEY: ", key)
		for _, valueArr := range value {
			fmt.Println("VALUE: ", valueArr)
		}
		fmt.Println()
	}

}
