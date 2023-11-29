package main

import "fmt"

func main() {
	var testString string = "test2"
	fmt.Println(testString)
	nonstaticvar, secondvar := testFunc() // := will set the var to the type returned by func
	fmt.Println(nonstaticvar, secondvar)
}

func testFunc() (string, string) {
	return "something", "else"
}
