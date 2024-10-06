package main
import (
	"fmt"
	"math/rand"
)

func main() {
	var number, i = 100, 0
	for {
		guess_number := rand.Intn(100) + 1
		i++
		if guess_number == number {
			fmt.Println("Guesses ", i)
			break
		}
	}
}

