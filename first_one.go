// Methods on Structures, constructors
package main

import "fmt"


type coordinate struct {
	d, m, s float64
	h rune
}

// adding a method of coordinate type
func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1.0
	}
	return sign * (c.d + c.m / 60.0 + c.s / 3600.0)
}

// constructor for coordinate type
func newCoordinate(d, m, s float64, h rune) coordinate {
	return coordinate{d, m, s, h}
}


type location struct {
	lat, long float64
}

// constructor for location
func newLocation(lat, long coordinate) location {
	return location{lat.decimal(), long.decimal()}
}

func main() {
	latitude := newCoordinate(4, 35, 22.2, 'S')
	longitude := newCoordinate(137, 26, 30.1, 'E')

	l1 := newLocation(latitude, longitude)
	fmt.Printf("%+v\n%+v\n%+v\n", latitude, longitude, l1)
}