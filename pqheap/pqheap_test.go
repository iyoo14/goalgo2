package pqheap

import (
	"fmt"
	"testing"
)

func TestUpheap(t *testing.T) {
	fmt.Println("vim-go")
	n := []string{"f", "b", "a", "d", "e", "c"}
	q := NewQueue(6, n)

	fmt.Println(n)
	fmt.Println(q.p)

	q.Remove()
	fmt.Println(q.p)
}
