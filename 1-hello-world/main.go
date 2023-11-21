package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	var whatToSay string
	var i int
	whatToSay = "Goodbye, cruel world"
	fmt.Println(whatToSay)
	i = 1
	fmt.Println("is set to", i)
	whatWasSaid, theOtherThingThatWasSaid := saySomething()
	fmt.Println(whatWasSaid, theOtherThingThatWasSaid)
}

func saySomething() (string, string) {
	return "something", "else"
}
