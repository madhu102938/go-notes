// Pointers

package main

import "fmt"

type celsius float64

// pointers with methods
// if pointer not used then it will not be incremented
func (c *celsius) increment() {
	(*c)++
}

type person struct {
	name, superpower	string
	age					int
}

func (p *person) increment() {
	p.age++ // automatic dereference :)
}

func main() {
	answer := 42
	fmt.Println(&answer) // address
	fmt.Println(*&answer) // dereferencing
	fmt.Printf("Pointer is the of the type %T\n", &answer)

	// &10 and &"hello world" can't point to simpler literals

	// *int can only point to int
	var character *string
	eren := "Eren Yeager"
	character = &eren // character points to the address of eren
	mikasa := "Mikasa"
	character = &mikasa // character points to the address of mikasa
	fmt.Println(*character)

	mikasa = "Mikasa Arckamann" // changes to `mikasa` changes `character`
	fmt.Println(*character)

	*character = "Levi" // We can directly change the value that character is pointing to
	fmt.Println(mikasa) // by changing `character` we can change `mikasa`

	character2 := character // `character2` points to same location as `character1`
	*character2 = "Armin"
	fmt.Println(mikasa) // changing `character2` changes `mikasa`

	fmt.Println(character == character2) // true

	armin := *character2 // `armin` is a copy of `*character2`
	armin = "Not Armin" // changing `armin` doesn't change the contents of address of `character2`
	fmt.Printf("%+v\t%+v\n", armin, *character2)

	// points to STRUCTURES
	// We can point to composite literals in structure

	timmy := &person{
		name : "Timmy",
		age : 10,
	}

	fmt.Println(timmy.name, (*timmy).name)
	// structures have automatic dereferencing, thus `timmy.name` is good enough

	timmy.superpower = "flying"
	fmt.Printf("%+v\n", timmy)

	// pointers on ARRAYS
	// As with structures, composite literals for arrays can be prefixed 
	// with the address operator (&) to create a new pointer to an array
	// Arrays also provide automatic dereferencing
	superpowers := &[...]string{"flight", "invisibility", "super strength"}
	fmt.Println(superpowers[0], (*superpowers)[0]) // same results (automatic dereferencing)

	// Composite literals for slices and maps can also be prefixed 
	// with the address operator (&), but there’s no automatic dereferencing.

	// pointers as parameters
	birthday := func(p *person) {
		p.age++
	} // if `p person` were used the changes won't be reflected outside the function

	rebecca := person{
		name:       "Rebecca",
		superpower: "imagination",
		age:        14,
	}

	birthday(&rebecca)
	fmt.Println(rebecca)

	// pointers with methods
	temp1 := celsius(2)
	fmt.Print(temp1, " ")
	temp1.increment()
	// (&temp1).increment() is the proper way, but Go handles things for us, so we can do the normal way :)
	fmt.Print(temp1, "\n")
	
	fmt.Print(rebecca.age, " ")
	rebecca.increment()
	fmt.Print(rebecca.age, "\n")

	// maps and slices are pointers to begin with, so pointer to them
	// would be redundant
	// one would use pointers on slices, only if one wants to edit
	// the slice itself i.e., length, capacity of slice
	planets := []string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	reclassify := func(p *[]string) {
		*p = (*p)[0:4]
	}

	reclassify(&planets)
	fmt.Println(planets)

}