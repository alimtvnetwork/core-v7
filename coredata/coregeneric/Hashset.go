package coregeneric

import (
	"fmt"
	"sync"
)

// Hashset is a generic set backed by map[T]bool with embedded mutex.
// It generalizes corestr.Hashset from string-only to any comparable type T.
type Hashset[T comparable] struct {
	items map[T]bool
	sync.RWMutex
}

// EmptyHashset creates a zero-capacity Hashset[T].
func EmptyHashset[T comparable]() *Hashset[T] {
	return NewHashset[T](0)
}

// NewHashset creates a Hashset[T] with pre-allocated capacity.
func NewHashset[T comparable](capacity int) *Hashset[T] {
	return &Hashset[T]{
		items: make(map[T]bool, capacity),
	}
}

// HashsetFrom creates a Hashset[T] from a slice of items.
func HashsetFrom[T comparable](items []T) *Hashset[T] {
	hs := NewHashset[T](len(items))

	for _, item := range items {
		hs.items[item] = true
	}

	return hs
}

// HashsetFromMap creates a Hashset[T] from an existing map.
func HashsetFromMap[T comparable](itemsMap map[T]bool) *Hashset[T] {
	return &Hashset[T]{
		items: itemsMap,
	}
}

// IsEmpty returns true if the hashset has no items.
func (it *Hashset[T]) IsEmpty() bool {
	return it == nil || len(it.items) == 0
}

// HasItems returns true if the hashset has at least one item.
func (it *Hashset[T]) HasItems() bool {
	return !it.IsEmpty()
}

// IsEmptyLock returns IsEmpty with mutex protection.
func (it *Hashset[T]) IsEmptyLock() bool {
	it.RLock()
	defer it.RUnlock()

	return it.IsEmpty()
}

// Length returns the number of items in the set.
func (it *Hashset[T]) Length() int {
	if it == nil || it.items == nil {
		return 0
	}

	return len(it.items)
}

// LengthLock returns the length with mutex protection.
func (it *Hashset[T]) LengthLock() int {
	it.RLock()
	defer it.RUnlock()

	return it.Length()
}

// Add adds an item to the set.
func (it *Hashset[T]) Add(key T) *Hashset[T] {
	it.items[key] = true

	return it
}

// AddBool adds an item and returns true if it already existed.
func (it *Hashset[T]) AddBool(key T) (isExist bool) {
	_, has := it.items[key]
	isNew := !has

	if isNew {
		it.items[key] = true
	}

	return has
}

// AddLock adds an item with mutex protection.
func (it *Hashset[T]) AddLock(key T) *Hashset[T] {
	it.Lock()
	it.items[key] = true
	it.Unlock()

	return it
}

// Adds adds multiple items to the set.
func (it *Hashset[T]) Adds(keys ...T) *Hashset[T] {
	for _, key := range keys {
		it.items[key] = true
	}

	return it
}

// AddSlice adds all items from a slice.
func (it *Hashset[T]) AddSlice(keys []T) *Hashset[T] {
	for _, key := range keys {
		it.items[key] = true
	}

	return it
}

// AddSliceLock adds all items from a slice with mutex protection.
func (it *Hashset[T]) AddSliceLock(keys []T) *Hashset[T] {
	it.Lock()

	for _, key := range keys {
		it.items[key] = true
	}

	it.Unlock()

	return it
}

// AddIf adds the item only if the condition is true.
func (it *Hashset[T]) AddIf(isAdd bool, key T) *Hashset[T] {
	isSkip := !isAdd

	if isSkip {
		return it
	}

	return it.Add(key)
}

// AddIfMany adds items only if the condition is true.
func (it *Hashset[T]) AddIfMany(isAdd bool, keys ...T) *Hashset[T] {
	isSkip := !isAdd

	if isSkip {
		return it
	}

	return it.Adds(keys...)
}

// AddHashsetItems merges another Hashset into this one.
func (it *Hashset[T]) AddHashsetItems(other *Hashset[T]) *Hashset[T] {
	if other == nil || other.IsEmpty() {
		return it
	}

	for k := range other.items {
		it.items[k] = true
	}

	return it
}

// AddItemsMap adds all entries from a map where value is true.
func (it *Hashset[T]) AddItemsMap(itemsMap map[T]bool) *Hashset[T] {
	for k, v := range itemsMap {
		if v {
			it.items[k] = true
		}
	}

	return it
}

// Has returns true if the key exists in the set.
func (it *Hashset[T]) Has(key T) bool {
	_, found := it.items[key]

	return found
}

// Contains is an alias for Has.
func (it *Hashset[T]) Contains(key T) bool {
	return it.Has(key)
}

// ContainsLock returns Contains with mutex protection.
func (it *Hashset[T]) ContainsLock(key T) bool {
	it.RLock()
	_, found := it.items[key]
	it.RUnlock()

	return found
}

// HasAll returns true if all keys exist in the set.
func (it *Hashset[T]) HasAll(keys ...T) bool {
	for _, key := range keys {
		isMissing := !it.Has(key)

		if isMissing {
			return false
		}
	}

	return true
}

// HasAny returns true if any of the keys exist in the set.
func (it *Hashset[T]) HasAny(keys ...T) bool {
	for _, key := range keys {
		if it.Has(key) {
			return true
		}
	}

	return false
}

// Remove removes a key from the set. Returns true if it existed.
func (it *Hashset[T]) Remove(key T) bool {
	_, existed := it.items[key]

	if existed {
		delete(it.items, key)
	}

	return existed
}

// RemoveLock removes a key with mutex protection.
func (it *Hashset[T]) RemoveLock(key T) bool {
	it.Lock()
	defer it.Unlock()

	return it.Remove(key)
}

// List returns all keys as a slice.
func (it *Hashset[T]) List() []T {
	if it.IsEmpty() {
		return []T{}
	}

	list := make([]T, 0, len(it.items))

	for k := range it.items {
		list = append(list, k)
	}

	return list
}

// ListPtr returns a pointer to the list of keys.
func (it *Hashset[T]) ListPtr() *[]T {
	list := it.List()

	return &list
}

// Map returns the underlying map.
func (it *Hashset[T]) Map() map[T]bool {
	return it.items
}

// Resize creates a new internal map with the given capacity, copying existing items.
func (it *Hashset[T]) Resize(capacity int) *Hashset[T] {
	if it.Length() > capacity {
		return it
	}

	newMap := make(map[T]bool, capacity)

	for k := range it.items {
		newMap[k] = true
	}

	it.items = newMap

	return it
}

// Collection converts the Hashset to a Collection[T].
func (it *Hashset[T]) Collection() *Collection[T] {
	return CollectionFrom[T](it.List())
}

// IsEquals compares two Hashsets for equality.
func (it *Hashset[T]) IsEquals(other *Hashset[T]) bool {
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
func (it *Hashset[T]) String() string {
	return fmt.Sprintf("%v", it.List())
}
