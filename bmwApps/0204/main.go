package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirtDate    time.Time
}

func main() {
	user := User{
		FirstName:   "Besto",
		LastName:    "Testo",
		PhoneNumber: "11 555 1212",
	}

	log.Println(user.FirstName, user.LastName, user.PhoneNumber, "Birthdate:", user.BirtDate)

}
