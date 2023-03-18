package function

// YOU NEED TO SPECIFY FUNCION PARAMS AND RETURN VALUE OF FUNCTION
func FactoryPatternOperation(op Operation) func(float32, float32) float32 {
	switch op {
	case SUM:
		return sum
	case SUB:
		return sub
	case DIV:
		return div
	case MUL:
		return mul
	default:
		return nil
	}
}

func sum(x, y float32) float32 {
	return x + y
}

func sub(x, y float32) float32 {
	return x - y
}

func div(x, y float32) float32 {
	if y == 0 {
		return 0
	}
	return x / y
}

func mul(x, y float32) float32 {
	return x * y
}
