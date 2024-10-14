// Interfaces and Pointers

package main

import (
	"fmt"
	"strings"
)

type talker interface {
	talk() string
}

type martian struct {}

func (m martian) talk() string {
	return "Pew Pew"
}

type celsius float64

func (c *celsius) talk() string {
	return "It's Hot"
}

func shout(t talker) {
	fmt.Println(strings.ToUpper(t.talk()))
}

func main() {
	shout(martian{})
	shout(&martian{})
	/*
	Interface is satisfied by both martian and *martian
	*/

	first := celsius(1)
	shout(first)	// gives error T_T (should match the exact method definition)
	shout(&first)
	/*
	Interface is satisfied by *celsius but not celsius (as the method is defined for *celsius)
	*/
}