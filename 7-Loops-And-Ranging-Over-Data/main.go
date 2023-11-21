package main

import "log"

func main() {
	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
		Gender    string
		isAdmin   bool
	}
	var users []User
	users = append(users, User{"John", "Smith", "a@example.com", 30, "fluid", true})
	users = append(users, User{"Mary", "Jones", "b@example.com", 25, "male", false})
	users = append(users, User{"Sally", "Brown", "c@example.com", 35, "female", false})
	users = append(users, User{"Alex", "Blood", "d@example.com", 37, "umom", false})

	for _, l := range users {
		log.Println(l)
	}
}
