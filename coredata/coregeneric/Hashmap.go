package coregeneric

import (
	"fmt"
	"sync"
)

// Hashmap is a generic map wrapper with embedded mutex for thread safety.
// It generalizes corestr.Hashmap from map[string]string to map[K]V.
type Hashmap[K comparable, V any] struct {
	items map[K]V
	sync.RWMutex
}

// EmptyHashmap creates a zero-capacity Hashmap[K, V].
func EmptyHashmap[K comparable, V any]() *Hashmap[K, V] {
	return NewHashmap[K, V](0)
}

// NewHashmap creates a Hashmap[K, V] with pre-allocated capacity.
func NewHashmap[K comparable, V any](capacity int) *Hashmap[K, V] {
	return &Hashmap[K, V]{
		items: make(map[K]V, capacity),
	}
}

// HashmapFrom creates a Hashmap[K, V] from an existing map (no copy).
func HashmapFrom[K comparable, V any](items map[K]V) *Hashmap[K, V] {
	return &Hashmap[K, V]{
		items: items,
	}
}

// HashmapClone creates a Hashmap[K, V] by copying entries from an existing map.
func HashmapClone[K comparable, V any](items map[K]V) *Hashmap[K, V] {
	hm := NewHashmap[K, V](len(items))

	for k, v := range items {
		hm.items[k] = v
	}

	return hm
}

func (it *Hashmap[K, V]) IsEmpty() bool {
	return it == nil || len(it.items) == 0
}

// HasItems returns true if the hashmap has at least one item.
func (it *Hashmap[K, V]) HasItems() bool {
	return it != nil && !it.IsEmpty()
}

// IsEmptyLock returns IsEmpty with mutex protection.
func (it *Hashmap[K, V]) IsEmptyLock() bool {
	it.RLock()
	defer it.RUnlock()

	return it.IsEmpty()
}

// Length returns the number of entries.
func (it *Hashmap[K, V]) Length() int {
	if it == nil || it.items == nil {
		return 0
	}

	return len(it.items)
}

// LengthLock returns the length with mutex protection.
func (it *Hashmap[K, V]) LengthLock() int {
	it.RLock()
	defer it.RUnlock()

	return it.Length()
}

// Set adds or updates a key-value pair. Returns true if the key was newly added.
func (it *Hashmap[K, V]) Set(key K, val V) (isAddedNewly bool) {
	_, isAlreadyExist := it.items[key]
	it.items[key] = val

	return !isAlreadyExist
}

// SetLock adds or updates a key-value pair with mutex protection.
func (it *Hashmap[K, V]) SetLock(key K, val V) *Hashmap[K, V] {
	it.Lock()
	defer it.Unlock()

	it.items[key] = val

	return it
}

// Get returns the value for the given key and whether it was found.
func (it *Hashmap[K, V]) Get(key K) (V, bool) {
	val, found := it.items[key]

	return val, found
}

// GetOrDefault returns the value for the key, or a default value if not found.
func (it *Hashmap[K, V]) GetOrDefault(key K, defaultVal V) V {
	val, found := it.items[key]
	isMissing := !found

	if isMissing {
		return defaultVal
	}

	return val
}

// GetLock returns the value with mutex protection.
func (it *Hashmap[K, V]) GetLock(key K) (V, bool) {
	it.RLock()
	defer it.RUnlock()

	val, found := it.items[key]

	return val, found
}

// Has returns true if the key exists.
func (it *Hashmap[K, V]) Has(key K) bool {
	_, found := it.items[key]

	return found
}

// Contains is an alias for Has.
func (it *Hashmap[K, V]) Contains(key K) bool {
	return it.Has(key)
}

// ContainsLock returns Has with mutex protection.
func (it *Hashmap[K, V]) ContainsLock(key K) bool {
	it.RLock()
	_, found := it.items[key]
	it.RUnlock()

	return found
}

// IsKeyMissing returns true if the key does not exist.
func (it *Hashmap[K, V]) IsKeyMissing(key K) bool {
	return !it.Has(key)
}

// Remove deletes a key from the map. Returns true if it existed.
func (it *Hashmap[K, V]) Remove(key K) bool {
	_, isExisted := it.items[key]

	if isExisted {
		delete(it.items, key)
	}

	return isExisted
}

// RemoveLock removes a key with mutex protection.
func (it *Hashmap[K, V]) RemoveLock(key K) bool {
	it.Lock()
	defer it.Unlock()

	return it.Remove(key)
}

// AddOrUpdateMap merges entries from another map.
func (it *Hashmap[K, V]) AddOrUpdateMap(itemsMap map[K]V) *Hashmap[K, V] {
	if len(itemsMap) == 0 {
		return it
	}

	for k, v := range itemsMap {
		it.items[k] = v
	}

	return it
}

// AddOrUpdateHashmap merges entries from another Hashmap.
func (it *Hashmap[K, V]) AddOrUpdateHashmap(other *Hashmap[K, V]) *Hashmap[K, V] {
	if other == nil || other.IsEmpty() {
		return it
	}

	for k, v := range other.items {
		it.items[k] = v
	}

	return it
}

// Keys returns all keys as a slice.
func (it *Hashmap[K, V]) Keys() []K {
	if it.IsEmpty() {
		return []K{}
	}

	keys := make([]K, 0, len(it.items))

	for k := range it.items {
		keys = append(keys, k)
	}

	return keys
}

// Values returns all values as a slice.
func (it *Hashmap[K, V]) Values() []V {
	if it.IsEmpty() {
		return []V{}
	}

	values := make([]V, 0, len(it.items))

	for _, v := range it.items {
		values = append(values, v)
	}

	return values
}

// Map returns the underlying map.
func (it *Hashmap[K, V]) Map() map[K]V {
	return it.items
}

// ForEach calls fn for each key-value pair.
func (it *Hashmap[K, V]) ForEach(fn func(key K, val V)) {
	for k, v := range it.items {
		fn(k, v)
	}
}

// ForEachBreak calls fn for each key-value pair; stops if fn returns true.
func (it *Hashmap[K, V]) ForEachBreak(fn func(key K, val V) (isBreak bool)) {
	for k, v := range it.items {
		if fn(k, v) {
			return
		}
	}
}

// ConcatNew creates a new hashmap combining this one with others.
func (it *Hashmap[K, V]) ConcatNew(others ...*Hashmap[K, V]) *Hashmap[K, V] {
	totalLen := it.Length()

	for _, other := range others {
		if other != nil {
			totalLen += other.Length()
		}
	}

	result := NewHashmap[K, V](totalLen)
	result.AddOrUpdateMap(it.items)

	for _, other := range others {
		result.AddOrUpdateHashmap(other)
	}

	return result
}

// Clone creates a copy of the Hashmap.
func (it *Hashmap[K, V]) Clone() *Hashmap[K, V] {
	return HashmapClone[K, V](it.items)
}

// IsEquals compares two Hashmaps for key equality (values not compared due to any constraint).
func (it *Hashmap[K, V]) IsEquals(other *Hashmap[K, V]) bool {
	if it == nil && other == nil {
		return true
	}

	if it == nil || other == nil {
		return false
	}

	if it == other {
		return true
	}

	if it.Length() != other.Length() {
		return false
	}

	for k := range it.items {
		isMissing := !other.Has(k)

		if isMissing {
			return false
		}
	}

	return true
}

// String returns a string representation.
func (it *Hashmap[K, V]) String() string {
	return fmt.Sprintf("%v", it.items)
}
