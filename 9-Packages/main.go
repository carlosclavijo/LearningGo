package main

import (
	"log"

	"github.com/carlosclavijo/LearningGo/helpers"
)

const numPool = 100000

func main() {
	intChan := make(chan int)
	defer close(intChan)
	go CalculateValue(intChan)
	num := <-intChan
	log.Println(num)
}

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}
