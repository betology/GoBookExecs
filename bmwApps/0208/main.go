package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	var mySlice []User

	u1 := User{
		FirstName: "Trevor",
	}
	u2 := User{
		FirstName: "Sam",
	}

	mySlice = append(mySlice, u1)
	mySlice = append(mySlice, u2)

	for i, x := range mySlice {
		log.Println(i, x.FirstName)
	}
}
