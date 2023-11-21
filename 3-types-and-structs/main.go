package main

import (
	"log"
	"time"
)

var s = "seven"

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName:   "Trevor",
		LastName:    "Sawler",
		PhoneNumber: "123-456-7890",
		Age:         25,
		BirthDate:   time.Now(),
	}
	log.Println(user.FirstName, user.LastName, user.PhoneNumber)
}

func whatever() {
	log.Println("XD")
}
