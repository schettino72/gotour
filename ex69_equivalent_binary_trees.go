package main

import "fmt"
import "code.google.com/p/go-tour/tree"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walk_sub(t, ch)
	close(ch)
}
func walk_sub(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		walk_sub(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		walk_sub(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		v1, open1 := <-ch1
		v2, open2 := <-ch2
		if v1 != v2 {
			return false
		}
		if !open1 || !open2 {
			return !open1 && !open2
		}
	}
	return true
}

func main() {
	t1 := tree.New(1)
	/*
	   ch := make(chan int, 10)
	   go Walk(t1, ch)
	   for val := range(ch){
	       fmt.Println(val)
	   }
	*/
	fmt.Println(Same(t1, tree.New(1)))
	fmt.Println(Same(t1, tree.New(2)))
}
