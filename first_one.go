package main

import (
	"fmt"
)

func random(a [5]string) {
	for i, j := range a {
		fmt.Println(i, j)
	}
}

func main() {
	var planets [8]string
	planets[0] = "mercury"
	planets[1] = "venus"
	planets[2] = "earth"

	fmt.Println(planets[2])
	fmt.Println(len(planets))
	fmt.Println(planets[3] == "") // initialized with zero values

	// composite literals
	names := [5]string{"madhu", "somename", "othername", "even", "odd"}
	for i := 0; i < len(names); i++ {
		// fmt.Println(names[i])
	}

	planets2 := [...]string{ 
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",// last comma is required
	} 
	fmt.Println(len(planets2))
	
	random(names, // comma required
	)

	// random(planets) // error, size doesn't match

	// 2d arrays
	var nums [2][2]int
	for i := range nums[0] {
		nums[0][i] = 5
	}

	for i := range nums {
		for j := range nums[i] {
			fmt.Print(nums[i][j], " ")
		}
		fmt.Println()
	}


}