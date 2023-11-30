package main

import (
	"encoding/json"
	"fmt"
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

	// []byte(mjJson) converts the json object to a slice of bytes that can be fed to json.Unmarshal
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json")
	}

	log.Printf("unmarshalled: %v", unmarshalled)

	// writing json from a struct
	var mySlice []Person

	var m1 Person
	m1.FirstName = "Wally"
	m1.LastName = "West"
	m1.HairColor = "red"
	m1.HasDog = false

	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "Diana"
	m2.LastName = "Prince"
	m2.HairColor = "black"
	m2.HasDog = false

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "    ")
	if err != nil {
		log.Println("error marshalling", err)
	}

	fmt.Println(string(newJson))
}
