// stacks, queues
// 1. Run `go get github.com/golang-collections/collections`
// Then run the code below

package main

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
	"github.com/golang-collections/collections/queue"
)

func main() {
	s := stack.New()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

	q := queue.New()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
}