package pq

import (
	"testing"
)

type Node struct {
	key string
	priority int
}

func TestMain(t *testing.T) {
	pq := NewPriorityQueue(func (a, b interface{}) bool {
		return a.(*Node).priority < b.(*Node).priority
	})
	qSize := 100
	for i := 0; i < qSize; i++ {
		pq.Push(&Node{"key", i})
	}

	if pq.Len() != qSize {
		t.Error("Expected", qSize, "Got", pq.Len())
	}

	for i := qSize - 1; i >= 0; i-- {
		head := pq.Head().(*Node)
		item := pq.Pop().(*Node)
		if head.priority != item.priority {
			t.Error("Not match return value form Head() and Pop()")
		}
		if item.priority != i {
			t.Error("Expected", i, "Got", item.priority)
		}
	}
}
