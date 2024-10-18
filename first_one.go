// Concurrency
// Select statement

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sleepyGopher(id int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(1500)) * (time.Millisecond))
	c <- id
}

func main() {
	timeout := time.After(1 * time.Second)
	c := make(chan int)

	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
	}

	for i := 0; i < 5; i++ {
		select {
		case gopherId := <- c:
			fmt.Println("GopherId :",gopherId)
		case <- timeout:
			fmt.Println("Timeout")
			return
		}
	}
}
/*
`time.After()` starts its timer when its called, then it
sends the current time on the channel, if time expires
by the time we reach the reciever, then we can just immidiately
read, thus it is advisable to use `After()` before the operation
*/