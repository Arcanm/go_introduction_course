package main

import (
	"fmt"
)

func main() {

	// DECLARE MAP
	// map[KEY TYPE OF VALUE]TYPE_OF_VALUE
	fmt.Println("----- MAPS -----")
	mapOne := make(map[int]string)
	mapOne[1] = "A"
	mapOne[2] = "B"
	mapOne[3] = "C"

	fmt.Println(mapOne)
	fmt.Println(mapOne[1])
	fmt.Println(mapOne[2])

	// OTHER WAY TO CREATE MAP
	mapTwo := map[int]string{
		1: "A",
		2: "B",
		3: "C",
	}

	fmt.Println(mapTwo)
	fmt.Println(mapTwo[1])
	fmt.Println(mapTwo[2])

	// Remove value from map
	delete(mapOne, 2)
	fmt.Println(mapOne)
	fmt.Println(mapOne[2])

	mapOne[1] = "A2"
	fmt.Println(mapOne)

	// WITH EMPTY OR INVALID KEY -> RETURN DEFAULT VALUE TYPE (0,"",0.0) ETC
	mapOne[5] = ""
	fmt.Println(mapOne[5])
	fmt.Println(mapOne[66])

	// WITH EMPTY BUT VALID KEY
	v, isPresent := mapOne[5]
	fmt.Printf("VALUE: %s, Present?: %t\n", v, isPresent)

	// WITH EMPTY BUT INVALID KEY
	v, isPresent = mapOne[66]
	fmt.Printf("VALUE: %s, Present?: %t", v, isPresent)

}
