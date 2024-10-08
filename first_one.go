// Structs, Json
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var first struct {
		one int
		two int	// remember no commas,
	}
	
	first.one, first.two = 1, 2
	fmt.Printf("%v\n", first)
	fmt.Printf("%+v\n", first) // focus on that +

	// Reusing structures with types
	type location struct{
		x int
		y int
	}

	var second location
	second.x, second.y = 3, 4

	// Initializing structures with composite literals
	third := location{x:5} // way 1
	// undeclared attributes are defaulted to their zero values

	forth := location{7, 8} // way 2
	// will not work if all attributes are not declared

	fmt.Println(third, forth)

	fifth := forth
	fifth.x = 5
	fmt.Println(forth, fifth) // structures are copied (pass by value)

	sliceStruct := []location{
		{x:1, y:2},
		{x:2, y:3},
		{x:3, y:4},
	}

	fmt.Println(sliceStruct)

	type coord struct {
		Lat float64
		Long float64
	}
	sixth := coord{2.23, 54.3}
	
	bytes, e := json.Marshal(sixth)
	
	if e != nil {
		os.Exit(1)
	} else {
		fmt.Println(string(bytes))
	}
	/*
	As we know that only those with capital letters are exported, so if 
	attributes of struct were lower case, they would not be exported,
	but this would restrict our naming conventions, so we can use tags
	to change the name of the attribute in the json file
	*/

	type coord2 struct {
		Lat float64 `json:"latitude"`
		Long float64 `json:"longitude"`
	} // `json:"latitude" xml:"latitude"` for more than one format

	seventh := coord2{34.24, 89.23}

	bytes, e = json.Marshal(seventh)

	if e != nil {
		os.Exit(1)
	} else {
		fmt.Println(string(bytes))
	}
}