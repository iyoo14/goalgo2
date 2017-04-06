package lkheap

import (
	"fmt"
	"math"
)

type Queue struct {
	cur   *Queue
	left  *Queue
	right *Queue
}

func NewQueue(max int) *Queue {
	q := new(Queue)
	q.cur = q
	q.left = nil
	q.right = nil
	return q
}

func (q *Queue) upheap() {
	v := q.node[q.cur]
	k := q.cur
	for v > q.node[k/2] {
		q.node[k] = q.node[k/2]
		k = k / 2
	}
	q.node[k] = v
}

func (q *Queue) downheap() {
	v := q.node[q.cur]
	q.node[1] = q.node[q.cur]
	q.cur--
	k := 1
	for k <= q.cur/2 {
		j := k + k
		if q.node[j] < q.node[j+1] {
			j = j + 1
		}
		if v >= q.node[j] {
			break
		}
		q.node[k] = q.node[j]
		k = j
	}
	q.node[k] = v
}

func (q *Queue) Insert(v int64) error {
	fmt.Println(q.cur, q.max)
	if q.cur >= q.max {
		return fmt.Errorf("queue is full.")
	}
	q.cur++
	q.node[q.cur] = v
	q.upheap()
	return nil
}

func (q *Queue) Remove() (int64, error) {
	if q.cur < 1 {
		return 0, fmt.Errorf("queue is empty.")
	}
	v := q.node[1]
	q.downheap()
	return v, nil
}

func (q *Queue) Show() {
	fmt.Println(q.node)
}
