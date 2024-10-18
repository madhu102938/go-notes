// Concurrency
// Pipelines

package main

import (
	"fmt"
	"strings"
)

func sourceGopher(downstream chan string) {
	for _, v := range []string{"hello world", "a bad apple", "goodbye all"} {
		downstream <- v
		fmt.Println(v, "sent on c0",)
	}
	close(downstream)
	fmt.Println("c0 Closed")
}

func filterGopherv1(upstream, downstream chan string) {
	for {
		item, ok := <- upstream
		if !ok {
			close(downstream)
			fmt.Println("c1 closed")
			return
		}
		fmt.Println(item, "received on c0")

		if !strings.Contains(item, "bad") {
			downstream <- item
			fmt.Println(item, "sent on c1")
		}
	}
}
// `filerGopher` can be refactored

func filterGopherv2(upstream, downstream chan string) {
	for item := range upstream {
		fmt.Println(item, "received on c0")
		if !strings.Contains(item, "bad") {
			downstream <- item
			fmt.Println(item, "sent on c1")
		}
	}

	close(downstream)
	fmt.Println("c1 closed")
}

func printGopher(upstream chan string) {
	for item := range upstream {
		fmt.Println(item)
	}
}

func main() {
	c0, c1 := make(chan string), make(chan string)
	go sourceGopher(c0)
	go filterGopherv1(c0, c1)
	printGopher(c1)
}