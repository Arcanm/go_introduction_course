package main

import (
	"fmt"

	"github.com/Arcanm/go_introduction_course/08-Accessors/mypackage"
)

func main() {
	var myCar mypackage.PublicCar
	myCar.Brand = "Honda"

	fmt.Println(myCar)
}
