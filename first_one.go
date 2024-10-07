package main

import (
	"fmt"
)

func main() {
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	terra := planets[:4]
	gasGaint := planets[4:6]
	iceGaint := planets[6:]
	myPlanet := terra[2:3]

	
	fmt.Println(planets)
	fmt.Println(terra, gasGaint, iceGaint)
	fmt.Println(myPlanet)

	myPlanet[0] = "Blue-chan"

	fmt.Println(planets)
	fmt.Println(terra, gasGaint, iceGaint)
	fmt.Println(myPlanet)

	parentString := "full string"
	subString := parentString[5:]

	fmt.Println(parentString,",", subString)

	subString = "another string"
	fmt.Println(parentString,",", subString)

	fmt.Printf("this of type %T\n", planets) // array
	fmt.Printf("this of type %T\n", gasGaint) // slice

	// initializing a slice
	names := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	fmt.Printf("%T, %T\n", names, planets)
}