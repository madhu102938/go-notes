// interfaces
// interface embedding and struct embedding

package main

import (
	"fmt"
)

type first interface {
	printString()
}

type second interface {
	returnInt() int
}

// interface embedding
type final interface {
	first
	second
}

type struct1 struct {}

func (s struct1) printString() {
	fmt.Println("printing from PrintString")
}

type struct2 struct {}

func (s struct2) returnInt() int {
	return 2
}

// struct embedding
type finalStruct struct {
	struct1
	struct2
}

func main() {
	var i1 first
	var i2 second
	var i3 final
	var s1 struct1
	var s2 struct2
	var s3 finalStruct

	i1 = s1
	i2 = s2
	i3 = s3

	i1.printString()
	fmt.Println(i2.returnInt())
	
	i3.printString()
	fmt.Println(i3.returnInt())
}

/*
what is interface embedding?
- interface embedding is a way to compose interfaces by including other interfaces as part of the definition of a new interface.
only the type which satisfies all the embedded interfaces is said to satisfy the final interface.
*/
