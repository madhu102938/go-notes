package main
import (
	"fmt"
)
func main() {
	var days_to_reach int = 28
	var distance int = 56_000_000
	fmt.Println("To reach in ", days_to_reach, " days, we need to travel at ", (float32(distance)) / (float32(days_to_reach * 24)), " kmph")
}