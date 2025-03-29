package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

// Add function to pc struct
func (testPc pc) ping() {
	fmt.Println(testPc.brand, "Pong")
}

// Update attributes for a struct from memory address
func (testPc *pc) duplicateRam() {
	testPc.ram = testPc.ram * 2
}

func main() {
	a := 50
	// Save in b memory address of a variable
	b := &a

	// Print memory address
	fmt.Println(b)

	// Print value of memory address
	fmt.Println(*b)

	myPc := pc{ram: 16, disk: 18, brand: "dell"}
	fmt.Println(myPc)

	myPc.ping()

	fmt.Println(myPc)
	myPc.duplicateRam()
	fmt.Println(myPc)
	myPc.duplicateRam()
	fmt.Println(myPc)
}
