package function

import (
	"errors"
	"fmt"
)

// FOR MUBLIC FUNCTIONS NEEDS TO INITIALIZE WITH UPPER CASE LETTER - Add
// FOR PRIVATE FUNCTIONS NEEDS TO INITIALIZE WITH DOWN CASE LETTE - add

// INITIALIZE MODULE - SIMLAR TO GEMFILE
// go mod init

// IT NEEDS TO BE PUBLIC BECAUSE WE CALL THESE FUNCTION IN MAIN
func Add(first_value int, second_value int) int {
	return first_value + second_value
}

func RepeatString(repeated int, value string) {
	for i := 0; i < repeated; i++ {
		fmt.Println(value)
	}
}

// SIMILAR TO ENUM, ASSING VALUES, 1, 2, 3..... TO CONSTANTS
type Operation int

const (
	SUM Operation = iota
	SUB
	DIV
	MUL
)

// YOU CAN DECLARE ONLY ONE TYPE IS IS THE SAME
// EXAMPLE 1: first_value float32, second_value float32
// EXAMPLE 2: first_value, second_value float32
func Calculator(first_value, second_value float32, operation Operation) (float32, error) {
	switch operation {
	case SUM:
		return first_value + second_value, nil
	case SUB:
		return first_value - second_value, nil
	case DIV:
		if second_value == 0 {
			return 0, errors.New("Second value musn't be zero")
		}
		return first_value / second_value, nil
	case MUL:
		return first_value * second_value, nil
	default:
		return 0, errors.New("Invalid Operation Type")
	}
}

// RETURN AND INSTANCE VARIABLES IN FUNCTION DECLARATION
// EXAMPLE 1:
// func Split(v int) (int, int) {
// 	x := v * 4 / 9
// 	y := v - x
// 	return x, y
// }

// EXAMPLE 2: DECLARATION OF VARIABLES IN FUNCTION DECLARATION
func Split(v int) (x, y int) {
	x = v * 4 / 9
	y = v - x
	return
}

// FUNCTION WITH DYNAMIC VALUES
// GOLANG TRANSFORM FUNCTION PARAMS INTO ARRAY
// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// IF WE USE ..., IT NEEDS TO BE THE LAST PARAMETER OF THE FUNCTION AND WE CANT HAVE TWO PARAMS WITH ..., EXAMPLES:
// func op(value ...float32)                 				 -> GOOD
// func op(op Operation, values ...float64)  				 -> GOOD

// func op(values ...float32, op Operation)  				 -> BAD
// func op(values1 ...float32, values2 ...string)    -> BAD
// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
func SumValues(values ...float32) float32 {
	var sum float32
	for _, value := range values {
		sum += value
	}
	return sum
}

func ArrayOperations(operation Operation, values ...float32) (float32, error) {
	if len(values) == 0 {
		return 0, errors.New("Minimum one value required")
	}
	result := values[0]
	for _, value := range values[1:] {
		switch operation {
		case SUM:
			result += value
		case SUB:
			result -= value
		case DIV:
			if value == 0 {
				return 0, errors.New("Musn't be zero")
			}
			result /= value
		case MUL:
			result *= value
		default:
			return 0, errors.New("Invalid Operation Type")
		}
	}
	return result, nil
}
