/*
Experiment: capacity.go
Write a program that uses a loop to continuously append an element to a slice. Print out 
the capacity of the slice whenever it changes. Does append always double the capacity 
when the underlying array runs out of room?
*/
package main

import "fmt"

func main() {
	var slice []string
	sliceCap := cap(slice)
	fmt.Println(sliceCap)

	for i := 0; i < 1000; i++ {
		slice = append(slice, "hehe")
		if sliceCap != cap(slice) {
			sliceCap = cap(slice)
			fmt.Println(sliceCap)
		}
	}
}