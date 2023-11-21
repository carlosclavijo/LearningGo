package main

import (
	helpers "command-line-argumentsC:\\Users\\CARLOS\\Desktop\\Aprendizaje\\Go\\ModernGolang\\learningGo\\9-Packages\\helpers\\helpers.go"
	"log"
)

func main() {
	var myVar helpers.SomeType
	myVar.TypeName = "some name"
	log.Println(myVar.TypeName)
}
