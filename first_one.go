package main

import (
	"fmt"
	"math"
)

func main() {
	var mp map[string]int	// map[key type]value type
	fmt.Println(mp["name"])	// keys which don't exist are initialized with zero values

	mp1 := map[string]int{
		"Earth" : 16,
		"Venus" : 464,
	}
	fmt.Println(mp1)

	// Checkiing if a key actually exists in the map
	if val, ok := mp["Mercury"]; ok {
		fmt.Printf("%d Value exists", val)
	} else {
		fmt.Println("Value doesn't exist")
	}

	slice1 := []int{1}
	slice2 := slice1
	slice2[0] = 2
	fmt.Println(slice1, slice2) // slices are passed by references

	mp2 := mp1
	fmt.Println(mp1["Earth"])
	mp1["Earth"] = 20
	fmt.Println(mp2["Earth"]) // maps also share are same underlying data

	// deleting elements from the map
	delete(mp1, "Venus")
	fmt.Println(mp1)

	// printing map using for loop
	for key, value := range mp1 {
		fmt.Println(key, value)
	}

	//  Instead of determining the frequency of temperatures, 
	// let’s group temperatures together in divisions of 10 each
	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}
	groups := map[float64][]float64{}
	for _, temps := range temperatures {
		division := math.Trunc(temps / 10.0) * 10
		groups[division] = append(groups[division], temps)
	}
	fmt.Println(groups)


	// using maps for set
	arr := []int{1,1,2,2,3,3,3}
	set := make(map[int]bool)
	for _, i := range(arr) {
		set[i] = true
	}

	fmt.Println(set[1], set[0])
}