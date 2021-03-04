package main

import "fmt"

type pc struct {
	ram int
	disk int
	brand string
}

func (myPc pc) ping() {
	fmt.Println(myPc.brand, "Pong")
}

func (myPc *pc) duplicateRAM() {
	myPc.ram = myPc.ram * 2
}

func (myPc pc) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d DB Disco y es una %s", myPc.ram, myPc.disk, myPc.brand)
}

func main() {
	a := 50
	b := &a

	fmt.Println(b)
	fmt.Println(*b)

	*b = 100
	fmt.Println(a)

	myPc := pc{ram:16, disk: 200, brand: "msi"}
	fmt.Println(myPc)

	myPc.ping()

	fmt.Println(myPc)
	myPc.duplicateRAM()
	
	fmt.Println(myPc)
	myPc.duplicateRAM()
	
	fmt.Println(myPc)
}