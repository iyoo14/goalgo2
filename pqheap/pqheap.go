package pqheap

import (
	"fmt"
)

type Queue struct {
	max  int
	cur  int
	node []string
	p    []int64
}

func NewQueue(max int, node []string) *Queue {
	q := new(Queue)
	q.max = max + 1
	q.cur = 0
	q.node = node
	q.p = make([]int64, q.max)
	q.p[0] = -1
	for i, _ := range node {
		q.cur++
		q.p[q.cur] = int64(i)
		q.upheap()
	}
	return q
}

func (q *Queue) upheap() {
	nv := q.node[q.p[q.cur]]
	v := q.p[q.cur]
	k := q.cur
	//fmt.Println(q.node)
	//fmt.Println(k)
	for {
		if k <= 1 {
			break
		}
		if nv > q.node[q.p[k/2]] {
			break
		}
		q.p[k] = q.p[k/2]
		k = k / 2
	}
	q.p[k] = v
}

func (q *Queue) downheap() {
	nv := q.p[q.cur]
	q.p[1] = q.p[q.cur]
	q.cur--
	k := 1
	for k <= q.cur/2 {
		j := k + k
		if q.node[q.p[j]] >= q.node[q.p[j+1]] {
			j = j + 1
		}
		if q.node[nv] < q.node[q.p[j]] {
			break
		}
		q.p[k] = q.p[j]
		k = j
	}
	q.p[k] = nv
}

func (q *Queue) Remove() {
	q.downheap()
}

func pre() {

	str := "ZA SIMPLE STRING TO BE ENCODED USING A MINIMAL NUMBER OF BITS"
	//kidx := make(map[int32]int32)
	//count := make([]int32,26)
	aidx := make(map[string]int, 27)
	aidx[" "] = 0
	for n := 0; n < 26; n++ {
		fmt.Println(string(n + 65))
		aidx[string(n+65)] = n + 1
	}
	fmt.Println(aidx)
	fq := make(map[int]int)
	for _, s := range str {
		fq[aidx[string(s)]]++
	}
	fmt.Println(fq)
}
