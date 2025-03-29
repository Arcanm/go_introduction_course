package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

// Stringer
func (myPc pc) String() string {
	return fmt.Sprintf("I have %dGB of ram, %dGB of disk and is a %s", myPc.ram, myPc.disk, myPc.brand)
}

func main() {
	myPc := pc{ram: 16, disk: 18, brand: "dell"}
	fmt.Println(myPc)
}
