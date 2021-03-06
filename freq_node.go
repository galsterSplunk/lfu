package lfu

import (
	set "github.com/deckarep/golang-set"
)

type freqNode struct {
	value int
	items set.Set
	prev  *freqNode
	next  *freqNode
}

func newFreqNode() *freqNode {
	n := &freqNode{
		value: 0,
		items: set.NewSet(),
	}
	n.prev = n
	n.next = n
	return n
}

func (fn *freqNode) remove() {
	fn.prev.next = fn.next
	fn.next.prev = fn.prev
}

func getNewNode(value int, prev, next *freqNode) *freqNode {
	newNode := newFreqNode()
	newNode.value = value
	newNode.prev = prev
	newNode.next = next

	prev.next = newNode
	next.prev = newNode

	return newNode
}
