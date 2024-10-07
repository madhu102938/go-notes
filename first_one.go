/*
 Experiment: terraform.go
 Write a program to terraform a slice of strings by prepending each planet with "New ".
Use your program to terraform Mars, Uranus, and Neptune.
 Your first iteration can use a terraform function, but your final implementation should
introduce a Planets type with a terraform method, similar to sort.StringSlice.
*/

package main

import "fmt"

type Planets []string

func (p Planets) terraform() {
	for i := range p {
		switch p[i] {
		case "Mars", "Uranus", "Neptune":
			p[i] = "New " + p[i]
		}
	}
}

func terraform(planets []string) {
	for i := range planets {
		switch planets[i] {
		case "Mars", "Uranus", "Neptune":
			planets[i] = "New " + planets[i]
		}
	}
}

func main() {
	planets := []string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	fmt.Println(planets)
	// terraform(planets)
	Planets(planets).terraform()
	fmt.Println(planets)

}