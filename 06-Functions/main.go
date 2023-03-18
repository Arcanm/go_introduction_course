package main

import (
	"fmt"

	"github.com/Arcanm/go_introduction_course/06-Functions/function"
)

func main() {
	fmt.Println(function.Add(3, 4))
	function.RepeatString(10, "Hola")
	v, err := function.Calculator(3, 6, function.SUM)
	fmt.Println("VALUE: ", v, "- ERROR: ", err)

	// PRINT ERROR AS STRING -> err.Error()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("VALUE: ", v, "- ERROR: ", err)
	}

	fmt.Println()
	v, err = function.Calculator(3, 0, function.DIV)
	fmt.Println("VALUE: ", v, "- ERROR: ", err)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("VALUE: ", v, "- ERROR: ", err)
	}

	fmt.Println()
	value1, value2 := function.Split(20)
	fmt.Println("VALUE: ", value1, "- VALUE 2: ", value2)

	fmt.Println()
	sum := function.SumValues(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("SUM: ", sum)

	fmt.Println()
	v, err = function.ArrayOperations(function.SUM, 1, 2, 3, 4, 5)
	fmt.Println("RESULT SUM: ", v, "- ERRORS: ", err)

	fmt.Println()
	v, err = function.ArrayOperations(function.DIV, 1, 2, 0, 4, 5)
	fmt.Println("RESULT DIV: ", v, "- ERRORS: ", err)

	fmt.Println()
	v, err = function.ArrayOperations(function.SUB, 1, 2, 0, 4, 5)
	fmt.Println("RESULT SUB: ", v, "- ERRORS: ", err)

	fmt.Println()
	factory := function.FactoryPatternOperation(function.SUB)
	v = factory(2, 3)
	fmt.Println("RESULT SUB: ", v)

	fmt.Println()
	factory = function.FactoryPatternOperation(function.MUL)
	v = factory(2, 3)
	fmt.Println("RESULT MUL: ", v)

}
