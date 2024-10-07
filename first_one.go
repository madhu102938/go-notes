package main

import (
	"fmt"
	"strings"
	"sort"
)

func changingSlice(names []string) {
	for i, name := range names {
		if name == "Pluto" {
			names[i] = "Hehe smol rock"
		}
	}
}

func main() {
	// initializing a slice
	names := []string{" Ceres  ", "Pluto", " Haumea ", " Makemake ", " Eris "}

	fmt.Println(names)
	changingSlice(names)
	fmt.Println(names)

	// trim and join
	for i := range(names) {
		names[i] = strings.Trim(names[i], " ")
	}
	fmt.Println(strings.Join(names, ","))

	sort.StringSlice(names).Sort()
	// 1. converting to StringSlice (provided by the sort package)
	// 2. sorting the slice by called the Sort() method on StringSlice type

	fmt.Println(names)
	fmt.Printf("%T\n", names) // []string
}