package pcpackage

import "fmt"

type Pc struct {
	ram   int
	disk  int
	brand string
}

func New(ram, disk int, brand string) Pc {
	myPc := Pc{ram: ram, disk: disk, brand: brand}
	return myPc
}

func (myPc Pc) PrintPc() {
	fmt.Printf("Pc brand %s has memory RAM of %dGB and disk of %dGB.\n", myPc.brand, myPc.ram, myPc.disk)
}

func (myPc *Pc) DuplicateRam() {
	myPc.ram = myPc.ram * 2
}

func (myPc Pc) GetDisk() int {
	return myPc.disk
}

func (myPc Pc) GetRam() int {
	return myPc.ram
}

func (myPc Pc) GetBrand() string {
	return myPc.brand
}
