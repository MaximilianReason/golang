package main

import "log"

func main() {
	var myString string
	myString = "green"

	log.Println("myString is set to", myString)
	changeUsingPointer(&myString)
	log.Println("after func call, myString is set to", myString)
}

func changeUsingPointer(s *string) {
	newValue := "red"
	*s = newValue
}
