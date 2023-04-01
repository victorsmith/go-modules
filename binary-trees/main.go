package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t != nil {
		Walk(t.Left, ch)
		ch <- t.Value
		fmt.Println(t.Value)
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int, 10)
	c2 := make(chan int, 10)

	go Walk(t1, c1)
	go Walk(t2, c2)

	for i := 0; i < 10; i++ {
		x, y := <-c1, <-c2

		fmt.Printf("x: %v, y: %v\n", x, y)
		if x != y {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("New(1) || New(1)", Same(tree.New(1), tree.New(1)))
	fmt.Println("New(1) || New(2)", Same(tree.New(1), tree.New(2)))
}
