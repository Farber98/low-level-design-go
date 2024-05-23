package storage

type StorageStrategy interface {
	Get(key int) (int, bool)
	Set(key, val int) bool
	Delete(key int) bool
	Length() int
}

// Let's define a hashmap stragey so we can make quick set and get operations in O(1) time
type HashmapStorageStrategy struct {
	mp map[int]int
}

func NewHashmapStorageStrategy(capacity int) *HashmapStorageStrategy {
	return &HashmapStorageStrategy{make(map[int]int, capacity)}
}

func (h *HashmapStorageStrategy) Get(key int) (int, bool) {
	// If exists, return the val and true as it was returned
	if val, ok := h.mp[key]; ok {
		return val, ok
	}

	// Else, return default val and false as it wasn't present
	return 0, false
}

func (h *HashmapStorageStrategy) Set(key, val int) bool {
	// If exists, wil return true as it existed
	// If it didn't existed, will return false
	var exists bool
	if _, ok := h.mp[key]; ok {
		exists = true
	}

	// Update or insert the key
	h.mp[key] = val
	return exists

}

func (h *HashmapStorageStrategy) Delete(key int) bool {
	// If exists, delete it
	_, exists := h.Get(key)
	if exists {
		delete(h.mp, key)
	}

	// Return if existed
	return exists
}

func (h *HashmapStorageStrategy) Length() int {
	return len(h.mp)
}
