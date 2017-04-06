package heap

import (
	"fmt"
	"testing"
)

func TestUpheap(t *testing.T) {
	fmt.Println("vim-go")
	q := NewQueue(5)
	q.Insert(2)
	q.Insert(1)
	q.Insert(5)
	q.Insert(4)
	q.Insert(3)
	q.Insert(3)
	q.Show()
	fmt.Println(q.Remove())
	fmt.Println(q.Remove())
	fmt.Println(q.Remove())
	fmt.Println(q.Remove())
	fmt.Println(q.Remove())
	fmt.Println(q.Remove())
}

func TestDownheap(t *testing.T) {
	fmt.Println("vim-go")
}
