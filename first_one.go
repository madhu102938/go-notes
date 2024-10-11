// Interfaces

package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type martial struct {}
type laser int

func (l laser)	talk() string {
	return strings.Repeat("pew ", int(l))
}

func (m martial) talk() string {
	return "nack nack"
}

// types for reusability
type talker interface {
	talk() string
}

// an interface is satisfied if the type implements the necessary methods
// else we will get compile time errors
func shout(t talker) { 
	fmt.Println(strings.ToUpper(t.talk()))
}

// Stringer interface, need to implement the String() method
func (m martial) String() string {
	return fmt.Sprintf("nack nack from stringer interface")
}

// Marshaler interface, need to implement the MarshalJSON() method
// func (m martial) MarshalJSON() ([]byte, error)
func (m martial) MarshalJSON() ([]byte, error) {
	str := `"knock knock"` // if not enclosed in `` then getting error T_T
	ans := []byte(str)
	fmt.Println(ans)
	return ans, nil
}

func main() {
	var t interface {
		talk() string
	}

	t = martial{}
	fmt.Println(t.talk())
	shout(t)
	
	t = laser(2)
	fmt.Println(t.talk())
	shout(t)
	// Polymorphism achieved :)

	// Using interfaces with struct embedding
	var starship struct{ laser }
	starship.laser = 3 
	//  Now the starship also satisfies the talker interface, 
	// allowing it to be used with shout
	shout(starship)

	// Any code can implement interfaces, even code that already exists (check obsidian)

	// Interfaces provided by the standard library
	// io.Reader, io.Writer, fmt.Stringer, json.Marshaler, json.Unmarshaler
	// fmt.Stringer is used to print the value of a type
	// json.Marshaler is used to convert a type to JSON
	// json.Unmarshaler is used to convert JSON to a type
	fmt.Println(martial{})

	bytes, err := json.Marshal(martial{})
	if err == nil {
		fmt.Println(string(bytes))
	} else {
		fmt.Println(err)
	}
}