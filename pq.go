package pq

import (
	"sync"
)

type CmpFunc func(a, b interface{}) bool

type PQ struct {
	sync.RWMutex
	items []interface{}
	size  int
	comparator CmpFunc
}

func NewPriorityQueue(cmp CmpFunc) *PQ {
	pq := &PQ{size: 0, comparator: cmp}
	pq.items = make([]interface{}, 1)
	pq.items[0] = nil
	return pq
}

func (pq  *PQ) Len() int {
	return pq.size
}

func (pq *PQ) Push(item interface{}) {
	pq.Lock()
	pq.items = append(pq.items, item)
	pq.size++
	pq.reorder(pq.Len())
	pq.Unlock()
}

func (pq *PQ) Pop() interface{} {
	pq.Lock()
	defer pq.Unlock()
	if pq.Len() < 1 {
		return nil
	}
	head := pq.items[1]
	pq.swap(1, pq.Len())
	pq.items = pq.items[0:pq.Len()]
	pq.size--
	pq.sink(1)
	return head
}

func (pq *PQ) Head() interface{} {
	pq.RLock()
	defer pq.RUnlock()
	if pq.Len() < 1 {
		return nil
	}
	return pq.items[1]
}

func (pq *PQ) less(i, j int) bool {
	return pq.comparator(pq.items[i], pq.items[j])
}

func (pq *PQ) swap(i, j int) {
	tmp := pq.items[i]
	pq.items[i] = pq.items[j]
	pq.items[j] = tmp
}

func (pq *PQ) reorder(i int) {
	for i > 1 && pq.less(i/2, i) {
		pq.swap(i/2, i)
		i = i /2
	}
}

func (pq *PQ) sink(i int) {
	for 2 * i <= pq.Len() {
		j := 2 * i
		if j < pq.Len() && pq.less(j, j + 1) {
			j++
		}
		if !pq.less(i, j) {
			break
		}
		pq.swap(i, j)
		i = j
	}
}
