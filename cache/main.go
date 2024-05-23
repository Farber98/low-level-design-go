package main

import (
	"fmt"

	"github.com/Farber98/low-level-design/cache/cache"
)

func main() {
	// Create a new LRUCache with a maximum capacity of 3
	cache := cache.NewLRUCache(3)
	cache.PrintCache("LRUCache(3)")

	// Set some key-value pairs
	cache.Set(1, 1)
	cache.PrintCache("Set(1,1)")
	cache.Set(2, 2)
	cache.PrintCache("Set(2,2)")
	cache.Set(3, 3)
	cache.PrintCache("Set(3,3)")

	// Get a value from the cache
	val, ok := cache.Get(1)
	if ok {
		fmt.Println("Value for key 1:", val) // Output: Value for key 1: 1
	}
	cache.PrintCache("Get(1)")

	// Set a new key-value pair, which should evict the least recently used key
	cache.Set(4, 4)
	cache.PrintCache("Set(4,4)")

	// Try to get the evicted key
	_, ok = cache.Get(2)
	cache.PrintCache("Get(2)")
	if !ok {
		fmt.Println("Key 2 has been evicted") // Output: Key 2 has been evicted
	}
}

/*
<>

INFORMATION:
- Build a cache

QUESTIONS & ASSUMPTIONS
- What capacity will we handle?
- We'll need just get, set, delete operations?
- We'll need more than one eviction policy?
	- LRU preferred?
	- LIFO, FIFO LFU others worth mentioning
- We'll need more than one storage class possibly?
- What we'll be storing?
	- key to vals as ints works good for example?

CONSTRAINTS:

COMPONENTS:
- Cache
	- Will be composed of our storage strategy, our eviction strategy and a size

- StorageStrategy
	- Will handle our underlying storage structure
	- Interface: get, set, delete

- EvictionStrategy
	- Will handle our eviction policies
	- Interface: put, evict



<>
*/
