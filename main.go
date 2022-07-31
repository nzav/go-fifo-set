package main

import (
	"carted/pkg/queue/fifoSet"
	"log"
)

// Implement FIFO set, e.g.
// 5 -> 1,2,3,4 => 1,2,3,4,5
// 2 -> 1,2,3,4 => 1,3,4,2

func main() {
	lst := fifoSet.New[string]()

	lst.Push("a")
	lst.Push("B")
	lst.Push("abc")
	lst.Push("a")
	lst.Push("B")

	popped, _ := lst.Pop()
	log.Printf("%v", popped)
}
