package main

import (
	"fmt"

	"github.com/Arcanm/go_introduction_course/Exercises/07-StructWithMethods/pcpackage"
)

func main() {
	myPc := pcpackage.New(12, 12, "Dell")
	fmt.Println(myPc)
	myPc.PrintPc()
	myPc.DuplicateRam()
	myPc.PrintPc()
	fmt.Println(myPc.GetRam())
}
