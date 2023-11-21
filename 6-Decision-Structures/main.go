package main

import "log"

func main() {
	myVar := "cat"
	switch myVar {
	case "cat":
		log.Println("meow")
	case "dog":
		log.Println("woof")
	default:
		log.Println("none")
	}
}
