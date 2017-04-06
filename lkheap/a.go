package main

import "fmt"

type node struct {
	cur   *node
	left  *node
	right *node
}

func NewNode() {
	n := new(node)
	n.left = nil
	n.right = nil
	var nn **node
	nn = &n

}

func main() {
	fmt.Println("vim-go")
}
