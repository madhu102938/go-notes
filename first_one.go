package main

import "fmt"

func dump(label string, slice []string) {
	fmt.Printf("%-15v length : %v | capacity : %v | %v\n", label, len(slice), cap(slice), slice)
}

func terraform(prefix string, planets ...string) []string {
	// variadic function receive the parameters as slices 

	newplanets := make([]string, len(planets))
	// making a copy, if `planets` were from a slice then editing `planets` can change that slice

	for i := range newplanets {
		newplanets[i] = prefix + " " + planets[i]
	}
	return newplanets
}

func main() {
	numbers := []string{
		"one",
		"two",
		"three",
	}

	// length, capacity
	dump("numbers", numbers)
	numbers = append(numbers, "four", "five")
	dump("numbers+", numbers) // double the capacity
	dump("numbers[1:3]", numbers[1:3])

	nums1 := []int{1, 2, 3, 4, 5}
	nums2 := append(nums1, 6)
	nums3 := append(nums2, 7, 8, 9)

	nums3[3] = 99 // doesn't affect nums1
	fmt.Printf("%v\n%v\n%v\n", nums1, nums2, nums3)

	arr1 := []int{1, 2, 3}
	arr2 := arr1[:1:3] // check this ;)
	arr2 = append(arr2, 4)
	fmt.Printf("%v\n%v\n", arr1, arr2)
	// depending on capacity of `arr2` `arr1` could be overwritten

	// Preallocating some length and capacity using `make`
	array := make([]int, 0, 10)
	array = append(array, 1, 2, 3, 4)
	fmt.Println("array: ", array)

	// variadic functions
	fmt.Println(terraform("New", "Venus", "Mars"))
	
	// passing slice in a variadic function
	planets := []string {
		"Venus",
		"Mars",
	}
	fmt.Println(terraform("Newer", planets...))
}