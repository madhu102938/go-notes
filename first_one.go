// nil
// A pointer with nowhere to point has the value nil.
// And the nil identifier is the zero value for slices, maps, and interfaces too.

package main

import (
	"fmt"
	"sort"
	// "log"
)

type person struct {
	age int
}

func (p *person) birthday() {
	if p == nil {
		// log.Fatal("nil pointer dereference") // is eqivalent to print followed by exit
		fmt.Println("nil pointer dereference")
		return
	}
	p.age++ // dereferencing a nil pointer will cause a panic
}

func sortStrings(s []string, less func(i, j int) bool) {
	if less == nil {
		less = func(i, j int) bool {
			return s[i] > s[j]
		}
	}
	sort.Slice(s, less)
}

func main() {
	var nobody *person
	fmt.Printf("%T %[1]v\n", nobody)
	nobody.birthday()

	// nil function values
	// zero value for a function type is nil
	var f func(a, b int) int
	fmt.Printf("%T %[1]v\n", f)
	// f(2, 3) // panic: call of nil function

	food := []string{"onion", "carrot", "celery"}
	sortStrings(food, nil)
	fmt.Println(food)

	less := func(i, j int) bool {
		return len(food[i]) > len(food[j])
	}

	sortStrings(food, less)
	fmt.Println(food)

	// nil slices
	// var first []string // nil slice
	// var second []string = make([]string, 0) // empty slice
	// var second = []string{} // empty slice

	var soup []string
	fmt.Println(soup == nil) // Prints true

	for _, ingredient := range soup {
		fmt.Println(ingredient)
	}
	
	fmt.Println(len(soup), cap(soup)) // Prints 0 0
	
	soup = append(soup, "onion", "carrot", "celery")
	fmt.Println(soup) // [onion carrot celery]
	// Range, len, cap and append work with nil slices

	// nil maps
	var m map[string]int // nil map
	// var m2 = make(map[string]int) // empty map

	fmt.Println(m == nil) // Prints true
	
	if value, ok := m["first"]; ok {
		fmt.Println(value)
	}
	// Writing to a nil map (soup["onion"] = 1) will panic with: assignment to entry in nil map

	// nil interfaces
	var v interface{}
	fmt.Println(v == nil)

	var p *person
	v = p
	fmt.Printf("%T %[1]v %v\n", v, v == nil) // *main.person <nil> false
	fmt.Printf("%#v\n", v) // (*main.person)(nil)
}