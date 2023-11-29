package main

import (
	"log"
	"time"
)

// Note Cap letter => Public
// Note small letter => Private
// global var is shadowed in funcs, if funcs declare with same name, it'll use the var from the func
var s = "seven" // auto detects as string

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	// Global-shadow example
	var s2 = "six"
	log.Println("s is", s)
	log.Println("s2 is", s2)
	saySomething("xxx")

	user := User{
		FirstName: "Max",
		LastName:  "Chow",
	}

	log.Println(user.FirstName, user.LastName, "Birthdate:", user.BirthDate)

}

// if func name is declared with small first letter, it's private, and can only be used within its declaration package
func saySomething(s string) (string, string) {
	log.Println("s from saySomething func is", s)
	return s, "World"
}

// func name with caps for first letter is public and can be used by other programs calling this package
func PublicFunc() {

}
