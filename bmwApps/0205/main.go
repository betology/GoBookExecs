package main

import "log"

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {
	var myVar myStruct
	myVar.FirstName = "Jon"

	myVar2 := myStruct{
		FirstName: "Mary",
	}

	log.Println("myVar is set to", (&myVar).printFirstName())
	log.Println("myVar is set to", myVar2.printFirstName())
}
