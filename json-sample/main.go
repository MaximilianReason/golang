package main

import (
	"encoding/json"
	"log"
)

// Able to use this struct to unmarshall the json obejct received
type Person struct {
	FirstName string `json:"first_name"` // this syntax tells go that when you receive a json, you shld look for the key "first_name" to assign the values
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `
[
	{
		"first_name": "Clark",
		"last_name": "Kent",
		"hair_color": "black",
		"has_dog": true
	},
	{
		"first_name": "Bruce",
		"last_name": "Wayne",
		"hair_color": "black",
		"has_dog": false
	}
]`

	// json will return slice in Go
	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json")
	}

	log.Printf("unmarshalled: %v", unmarshalled)
}
