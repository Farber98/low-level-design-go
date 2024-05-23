package cache

import (
	"fmt"

	"github.com/Farber98/low-level-design/cache/eviction"
	"github.com/Farber98/low-level-design/cache/storage"
)

type CacheInterface interface {
	Get(key int) (int, bool)
	Set(key, val int) bool
	Evict()
	PrintCache()
}

type LRUCache struct {
	storage     storage.StorageStrategy
	eviction    eviction.EvictionStrategy
	maxCapacity int
}

func NewLRUCache(maxCapacity int) *LRUCache {
	return &LRUCache{
		storage:     storage.NewHashmapStorageStrategy(maxCapacity),
		eviction:    eviction.NewLRUEvictionStrategy(),
		maxCapacity: maxCapacity,
	}
}

func (c *LRUCache) Get(key int) (int, bool) {
	// Check if exists
	val, ok := c.storage.Get(key)

	// If exists, we need to update it and put it as MRU
	if ok {
		c.eviction.Put(key)
	}

	// Return val and if it existed
	return val, ok
}

func (c *LRUCache) Set(key, val int) bool {
	// If it was present, just update val and MRU
	if _, ok := c.Get(key); ok {
		c.storage.Set(key, val)
		c.eviction.Put(key)
		return true
	}

	// If it was not present, we need to check if we are at capacity or not.
	if c.storage.Length() == c.maxCapacity {
		// If we are at capacity, evict LRU
		key := c.eviction.Evict()
		c.storage.Delete(key)
	}

	// Now we insert new node, put it as LRU and as it not existed we need to return false
	c.storage.Set(key, val)
	c.eviction.Put(key)
	return false
}

func (c *LRUCache) PrintCache(operation string) {
	fmt.Print("OP: " + operation + " | ")
	c.eviction.PrintEviction()
	fmt.Println()
}
