package main

import (
	"log"

	"github.com/MaximilianReason/myniceprogram/helpers"
)

const numPool = 10

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}

func main() {
	// var myVar helpers.SomeType
	// myVar.TypeName = "Some name"

	// log.Println(myVar.TypeName)

	// a channel that can only handle int
	intChan := make(chan int)
	// defer is a keyword for doing the call at the end of the function
	// in this case closes the channel at end of function
	defer close(intChan)

	// a go routine - a concurrent operation
	go CalculateValue(intChan)

	// listening for the result from the channel
	num := <-intChan
	log.Println(num)
}
