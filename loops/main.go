package main

import "log"

func main() {
	// simple loop
	// for i := 0; i <= 10; i++ {
	// 	log.Println(i)
	// }

	// slice loop
	// animals := []string{"dog", "cat", "fish"}
	// // _ implies ignore the looping variable, and loops the range
	// for _, animal := range animals {
	// 	log.Println(animal)
	// }

	// map loop
	// animals := make(map[string]string)
	// animals["dog"] = "Fido"
	// animals["cat"] = "Fluffy"

	// // prints out key & value of map
	// for animalType, animal := range animals {
	// 	log.Println(animalType, animal)
	// }

	// string loop
	// var firstLine = "this is a test line"
	// for i, l := range firstLine {
	// 	log.Println(i, ":", l)
	// }

	// custom type loop
	type User struct {
		FirstName string
		LastName  string
		Age       int
	}

	var users []User
	users = append(users, User{"John", "Smith", 30})
	users = append(users, User{"X", "Smith", 40})
	users = append(users, User{"Y", "Smith", 50})
	users = append(users, User{"Z", "Smith", 60})

	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Age)
	}

}
