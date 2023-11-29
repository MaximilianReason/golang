package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

// maps are immutable

func main() {
	// Maps
	myMap := make(map[string]User)

	me := User{
		FirstName: "Max",
		LastName:  "Chow",
	}

	myMap["me"] = me

	log.Println(myMap["me"])

	// Slice
	var mySlice []string

	mySlice = append(mySlice, "test1")
	mySlice = append(mySlice, "test2")
	mySlice = append(mySlice, "abc")

	sort.Strings(mySlice)

	log.Println(mySlice)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	log.Println(numbers)
	log.Println(numbers[0:5]) // can specify the range to print
}
