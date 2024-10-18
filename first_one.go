// concurrency
// go routines, channels

package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go sleepyGopher(i) // goroutines don't run in order
	}
	time.Sleep(1 * time.Second) 
	// without this line, the program will exit before the goroutines finish


	// channels
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go sleepyGopher2(i, c)
	}
	for i := 0; i < 5; i++ {
		fmt.Println(<- c) // it will wait until some is written
	}
}

// Could be some function with a long running time.
func sleepyGopher(id int) {
	fmt.Println(id, "...snore...")
}

func sleepyGopher2(id int, c chan int) {
	c <- id // it will wait until it is read
}