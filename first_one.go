// composition and forwarding
// composition : a structure can have a field that is another structure
// forwarding : a method of a structure can call a method of another structure

package main

import "fmt"

type report0 struct {
	sol int
	high, low int
	lat, long int
} // this is not very good, instead we can make a structure of structures

type celsius float64

type temperature struct {
	high, low celsius
}

type location struct {
	lat, long float64
}

type report1 struct {
	sol int
	temperature temperature
	location location
}

// a method of temperatures
func (t temperature) average() celsius {
	return (t.high + t.low) / 2.0
}

// updated report accounting for method forwarding
type report struct {
	sol int
	temperature // automatically named temperature
	location // automatically named
}

// we forward all types methods, not just the ones for a structure
type sol int
type report2 struct {
	sol
	temperature
	location
}

// a method for sol
func (s sol) days() int {
	return 1
}

// a method for location
func (l location) days() int {
	return 2
}

// resolving the name conflict
func (r report2) days() int {
	return r.sol.days()
}

func main() {
	loc1 := location{-4.5895, 137.4417}
	t1 := temperature{high:-1.0, low:-78.0}
	repo1 := report1{sol:15, temperature:t1, location:loc1}

	fmt.Printf("%+v\n", repo1)

	fmt.Println("Average :=", repo1.temperature.average())
	// to call average we need to repo1.temperature.average(), this is inconvienent
	// it would be nice if we can call average just by repo1.average()
	// we can do this by forwarding the call to the temperature type

	repo := report{sol:15, temperature:t1, location:loc1}
	fmt.Println("Average :=", repo.average()) // successfully forwarded

	// name conflicts
	repo2 := report2{sol:15, temperature:t1, location:loc1}
	
	// now we have two days methods, one from `sol` and one from `location`
	// good news is we get the error only we if call the method
	// fmt.Println(repo2.days()) // this will get ambigous selector error
	
	// parent structure method takes precedence over child structure method
	fmt.Println(repo2.days())

	// Inheritance vs Composition
	/*
	 Inheritance is a different way of thinking about designing software. With inheritance, a
	 rover is a type of vehicle and thereby inherits the functionality that all vehicles share.

	 With composition, a rover has an engine and wheels and various other parts that pro
	 vide the functionality a rover needs. A truck may reuse several of those parts, but there
	 is no vehicle type or hierarchy descending from it.
	*/
}