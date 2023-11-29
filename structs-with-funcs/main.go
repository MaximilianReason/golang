package main

import "log"

type myStruct struct {
	FirstName string
}

// tie the function to the struct by specifying the receiver variable for the function
// e.g. func (receiverVar *myStruc) printFirstName() string
func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {
	var myVar myStruct
	myVar.FirstName = "Max"

	myVar2 := myStruct{
		FirstName: "Tester",
	}

	log.Println("myVar is set to", myVar.printFirstName())
	log.Println("myVar2 is set to", myVar2.printFirstName())
}
