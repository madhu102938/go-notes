// New errors, custom errors, panics and recover

package main

import (
	"errors"
	"fmt"
)

const rows, columns = 9, 9
type grid [rows][columns]int8

func inBound(row, column int) bool {
	if row < 0 || row >= rows || column < 0 || column >= columns {
		return false
	}
	return true
}

func (g *grid) set1(row, column int, digit int8) error {
	if !inBound(row, column) {
		return errors.New("not in bounds")
	}

	g[row][column] = digit
	return nil
}

// best practices
var (
	ErrBounds = errors.New("not in bounds")
	ErrDigit = errors.New("invalid digit")
)

func validDigit(digit int8) bool {
	if digit > rows || digit < 0 {
		return false
	}
	return true
}

func (g *grid) set2(row, column int, digit int8) error {
	if !inBound(row, column) {
		return ErrBounds
	} else if !validDigit(digit) {
		return ErrDigit
	}

	g[row][column] = digit
	return nil
}

// custom error
type SudokuError []error
// to qualify as an `error` interface it needs to have an `Error() string` method
func (s SudokuError) Error() string {
	totalError := s[0].Error()
	for i := 1; i < len(s); i++ {
		totalError += ", " + s[i].Error()
	}
	return totalError
}

func (g *grid) set3(row, column int, digit int8) error {
	var s SudokuError

	if !inBound(row, column) || !validDigit(digit) {
		if !inBound(row, column) {
			s = append(s, ErrBounds)
		}
		if !validDigit(digit) {
			s = append(s, ErrDigit)
		}
		return s
	}

	g[row][column] = digit

	// return s (if you return s, then we cannot equate it to `nil` as its an interface, even if the underlying slice is `nil`, interface will not be `nil`)
	return nil
}


func main() {
	var sudoku grid
	err := sudoku.set1(10, 0, 3)
	if err != nil {
		fmt.Println(err)
	}

	err = sudoku.set2(3, 3, -3)
	if err != nil {
		fmt.Println(err)
	}

	err = sudoku.set3(4, 4, 4)
	if err != nil {
		fmt.Println(err)
	}

	// type assertions
	err = sudoku.set3(10, 10, 10)
	if err != nil {
		if errs, ok := err.(SudokuError); ok {
			fmt.Printf("%v error(s) occured:\n", len(errs))
			for index, error_i := range errs {
				fmt.Println(index+1, ":", error_i)
			}
		}
	}

	// panic and recover
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e) // prints "Program crashed T_T"
		}
	}() // only deferred functions can make use of `recover`
	// `panic` is often better than `os.Exit` in that `panic` will run any deferred functions, whereas `os.Exit` does not.

	panic("Program crashed T_T")
}