package main

import "log"

func main() {
	cat := "cat"

	// strcmp
	if cat == "cat" {
		log.Println("Cat is cat")
	} else {
		log.Println(("Cat is not cat"))
	}

	myNum := 100
	isTrue := false

	// if-else
	if myNum > 99 && !isTrue {
		log.Println("myNum > 99 & isTrue is set to false")
	} else {
		log.Println("conditions not met")
	}

	// switch
	myVar := "dog"

	// note that go auto breaks out from switch.. no carry over, hence no "break" keyword required
	switch myVar {
	case "dog":
		log.Println("dog is set to dog")
	case "rat":
		log.Println("rat is set to rat")
	default:
		log.Println("value not assigned")
	}

}
