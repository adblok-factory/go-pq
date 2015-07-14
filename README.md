pq
===

[![Build Status](https://travis-ci.org/hideo55/go-pq.svg?branch=master)](https://travis-ci.org/hideo55/go-pq)
[![Godoc](https://godoc.org/github.com/hideo55/go-pq?status.png)](https://godoc.org/github.com/hideo55/go-pq)
[![Coverage Status](https://coveralls.io/repos/hideo55/go-pq/badge.svg?branch=master)](https://coveralls.io/r/hideo55/go-pq?branch=master)

Description
------------

Priority Queue implementation for Go.

Usage
-----

```go
package main
import (
    "github.com/hideo55/go-pq"
)

type Node struct {
    depth     int
    beginNode int
    endNode   int
}
func main() {
    q := pq.NewPriorityQueue(func (a, b interface{}) bool {
        aVal := a.(*Node)
        bVal := b.(*Node)
        if aVal.depth != bVal.depth {
            return aVal.depth < bVal.depth
        } else {
            return aVal.beginNode < bVal.beginNode
        }
    })
    q.Push(&Node{0, 0, 1})
    q.Push(&Node{0, 1, 2})
        
    ...
    item := q.Pop().(*Node)
}
```

Supported version
-----------------

Go 1.4 or later

License
--------

MIT License
