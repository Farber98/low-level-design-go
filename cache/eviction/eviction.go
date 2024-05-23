package eviction

import "fmt"

type EvictionStrategy interface {
	Put(key int)
	Evict() int
	PrintEviction()
}

// We'll need a doubly linked list to track our LRU and be able to:
// Evict, put new ones and refresh LRU nodes
type LRUNode struct {
	key        int
	next, prev *LRUNode
}

func newLRUNode(key int) *LRUNode {
	return &LRUNode{key, nil, nil}
}

type LRUEvictionStrategy struct {
	// Map will let us map key to nodes for operate quickly with them
	mp         map[int]*LRUNode
	head, tail *LRUNode
	// We will do the following: <HEAD> <LRU> ... <MRU> <TAIL>
	// We will add to tail
	// We will remove from head.
}

func NewLRUEvictionStrategy() *LRUEvictionStrategy {
	// Initialize map, head and tail.
	mp := make(map[int]*LRUNode, 0)
	head, tail := newLRUNode(0), newLRUNode(0)

	// Wire head and tail
	head.next, tail.prev = tail, head

	// Return initialized strategy
	return &LRUEvictionStrategy{
		mp,
		head,
		tail,
	}

}

func (lru *LRUEvictionStrategy) Put(key int) {
	// If it's the first one, we'll need to point head and tail
	if lru.isEmpty() {
		// Create node
		node := newLRUNode(key)

		// Point head and tail to node
		lru.head.next, lru.tail.prev = node, node

		// Point node to head and tail
		node.prev, node.next = lru.head, lru.tail

		// put in our map
		lru.mp[key] = node

		return
	}

	// If LRU wasn't empty...
	// If key already existed, remove because we'll need to reinsert
	if _, ok := lru.mp[key]; ok {
		lru.evictNode(key)
	}

	// Now, we'll just need to insert (grabbing both cases, new and old node refreshed)
	newMRU := newLRUNode(key)
	oldMRU := lru.tail.prev
	// locate our node between prev MRU and tail
	newMRU.prev, newMRU.next = oldMRU, lru.tail
	// point prev MRU and tail to our node
	oldMRU.next, lru.tail.prev = newMRU, newMRU

	// put in our map
	lru.mp[key] = newMRU
}

func (lru *LRUEvictionStrategy) evictNode(key int) {
	// Evict only if not empty
	if !lru.isEmpty() {
		// Grab direction from mp
		node, ok := lru.mp[key]
		if !ok {
			return
		}

		prevNode, nextNode := node.prev, node.next
		// fix prev node
		prevNode.next = nextNode
		// fix following node
		nextNode.prev = prevNode

		// Remove from mp
		delete(lru.mp, key)
	}
}

func (lru *LRUEvictionStrategy) Evict() int {
	var evictedKey int
	// Evicts LRU
	if !lru.isEmpty() {
		nodeToEvict := lru.head.next

		newLRU := nodeToEvict.next

		// point new LRU to head and head to new LRU
		newLRU.prev = lru.head
		lru.head.next = newLRU

		// Remove from mp
		delete(lru.mp, nodeToEvict.key)
		evictedKey = nodeToEvict.key
	}

	// So we remove from storage
	return evictedKey
}

func (lru *LRUEvictionStrategy) isEmpty() bool {
	return len(lru.mp) == 0
}

func (lru *LRUEvictionStrategy) PrintEviction() {
	fmt.Print("QUEUE: ", printQueue(lru.head, lru.tail))
	fmt.Print(" | ")
	fmt.Print("MAP: ", lru.mp)
}

func printQueue(head, tail *LRUNode) string {
	q := "(HEAD)<->"
	for node := head.next; node != tail; node = node.next {
		q += fmt.Sprintf("(%v)<->", node.key)
	}
	q += "(TAIL)"
	return q
}
