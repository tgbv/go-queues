package main

import (
	"fmt"

	q "github.com/tgbv/go-queues/pkg"
)

func main() {
	w := q.Queue{}

	// Addition
	w.Add(22).Add(33).Add("meme")

	// traverse
	w.Traverse(func(n *q.Node) {
		fmt.Println(n)
	})

	// pop
	n := w.Pop()
	fmt.Println(n)

	// lookup
	fmt.Println(w.Find(func(n *q.Node) bool {
		if n.Value == 33 {
			return true
		} else {
			return false
		}
	}))

	//
	fmt.Println(w)
}
