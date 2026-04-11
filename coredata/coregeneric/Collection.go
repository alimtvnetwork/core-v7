package coregeneric

import (
	"fmt"
	"sort"
	"sync"
)

// Collection is a generic slice-backed collection with embedded mutex for thread safety.
// It generalizes corestr.Collection from string-only to any type T.
type Collection[T any] struct {
	items []T
	sync.RWMutex
}

// EmptyCollection creates a zero-capacity Collection[T].
func EmptyCollection[T any]() *Collection[T] {
	return &Collection[T]{items: []T{}}
}

// NewCollection creates a Collection[T] with pre-allocated capacity.
func NewCollection[T any](capacity int) *Collection[T] {
	return &Collection[T]{items: make([]T, 0, capacity)}
}

// CollectionFrom wraps an existing slice into a Collection[T] (no copy).
func CollectionFrom[T any](items []T) *Collection[T] {
	return &Collection[T]{items: items}
}

// CollectionClone copies items into a new Collection[T].
func CollectionClone[T any](items []T) *Collection[T] {
	cloned := make([]T, len(items))
	copy(cloned, items)

	return &Collection[T]{items: cloned}
}

// CollectionLenCap creates a Collection[T] with specific length and capacity.
func CollectionLenCap[T any](length, capacity int) *Collection[T] {
	return &Collection[T]{items: make([]T, length, capacity)}
}

// HasAnyItem returns true if the collection has at least one item.
func (it *Collection[T]) HasAnyItem() bool {
	return it.Length() > 0
}

// LastIndex returns the index of the last element.
func (it *Collection[T]) LastIndex() int {
	return it.Length() - 1
}

// HasIndex returns true if the given index is valid.
func (it *Collection[T]) HasIndex(index int) bool {
	return index >= 0 && it.LastIndex() >= index
}

// Items returns the underlying slice.
func (it *Collection[T]) Items() []T {
	return it.items
}

// ItemsPtr returns a pointer to the underlying slice.
func (it *Collection[T]) ItemsPtr() *[]T {
	return &it.items
}

// Count returns the number of elements (alias for Length).
func (it *Collection[T]) Count() int {
	return it.Length()
}

// Capacity returns the capacity of the underlying slice.
func (it *Collection[T]) Capacity() int {
	if it.items == nil {
		return 0
	}

	return cap(it.items)
}

// Length returns the number of elements in the collection.
func (it *Collection[T]) Length() int {
	if it == nil || it.items == nil {
		return 0
	}

	return len(it.items)
}

// LengthLock returns the length with mutex protection.
func (it *Collection[T]) LengthLock() int {
	it.RLock()
	defer it.RUnlock()

	return it.Length()
}

// IsEmpty returns true if the collection has no items.
func (it *Collection[T]) IsEmpty() bool {
	return it == nil || len(it.items) == 0
}

// IsEmptyLock returns IsEmpty with mutex protection.
func (it *Collection[T]) IsEmptyLock() bool {
	it.RLock()
	defer it.RUnlock()

	return it.IsEmpty()
}

// HasItems returns true if the collection has at least one item.
func (it *Collection[T]) HasItems() bool {
	return it != nil && len(it.items) > 0
}

// Add appends a single item to the collection.
func (it *Collection[T]) Add(item T) *Collection[T] {
	it.items = append(it.items, item)

	return it
}

// AddLock appends a single item with mutex protection.
func (it *Collection[T]) AddLock(item T) *Collection[T] {
	it.Lock()
	defer it.Unlock()

	it.items = append(it.items, item)

	return it
}

// Adds appends multiple items (variadic).
func (it *Collection[T]) Adds(items ...T) *Collection[T] {
	it.items = append(it.items, items...)

	return it
}

// AddsLock appends multiple items with mutex protection.
func (it *Collection[T]) AddsLock(items ...T) *Collection[T] {
	it.Lock()
	defer it.Unlock()

	it.items = append(it.items, items...)

	return it
}

// AddSlice appends all items from a slice.
func (it *Collection[T]) AddSlice(items []T) *Collection[T] {
	it.items = append(it.items, items...)

	return it
}

// AddIf appends the item only if the condition is true.
func (it *Collection[T]) AddIf(isAdd bool, item T) *Collection[T] {
	isSkip := !isAdd

	if isSkip {
		return it
	}

	it.items = append(it.items, item)

	return it
}

// AddIfMany appends items only if the condition is true.
func (it *Collection[T]) AddIfMany(isAdd bool, items ...T) *Collection[T] {
	isSkip := !isAdd

	if isSkip {
		return it
	}

	it.items = append(it.items, items...)

	return it
}

// AddFunc appends the result of calling the function.
func (it *Collection[T]) AddFunc(f func() T) *Collection[T] {
	it.items = append(it.items, f())

	return it
}

// AddCollection appends all items from another Collection[T].
func (it *Collection[T]) AddCollection(other *Collection[T]) *Collection[T] {
	if other.IsEmpty() {
		return it
	}

	return it.Adds(other.items...)
}

// AddCollections appends items from multiple collections.
func (it *Collection[T]) AddCollections(collections ...*Collection[T]) *Collection[T] {
	for _, col := range collections {
		if col.IsEmpty() {
			continue
		}

		it.AddSlice(col.items)
	}

	return it
}

// RemoveAt removes the item at the given index. Returns true on success.
func (it *Collection[T]) RemoveAt(index int) bool {
	length := it.Length()
	if index < 0 || index >= length {
		return false
	}

	it.items = append(it.items[:index], it.items[index+1:]...)

	return true
}

// First returns the first element. Panics on empty collection.
func (it *Collection[T]) First() T {
	return it.items[0]
}

// Last returns the last element. Panics on empty collection.
func (it *Collection[T]) Last() T {
	return it.items[it.LastIndex()]
}

// FirstOrDefault returns the first element or the zero value of T if empty.
func (it *Collection[T]) FirstOrDefault() T {
	if it.IsEmpty() {
		var zero T

		return zero
	}

	return it.First()
}

// LastOrDefault returns the last element or the zero value of T if empty.
func (it *Collection[T]) LastOrDefault() T {
	if it.IsEmpty() {
		var zero T

		return zero
	}

	return it.Last()
}

// SafeAt returns the element at the given index, or the zero value if out of bounds.
func (it *Collection[T]) SafeAt(index int) T {
	if it.IsEmpty() || !it.HasIndex(index) {
		var zero T

		return zero
	}

	return it.items[index]
}

// Skip returns items after skipping the first n.
func (it *Collection[T]) Skip(count int) []T {
	if count >= it.Length() {
		return []T{}
	}

	return it.items[count:]
}

// Take returns the first n items.
func (it *Collection[T]) Take(count int) []T {
	if count >= it.Length() {
		return it.items
	}

	return it.items[:count]
}

// ForEach calls fn for each item with its index.
func (it *Collection[T]) ForEach(fn func(index int, item T)) {
	for i, item := range it.items {
		fn(i, item)
	}
}

// ForEachBreak calls fn for each item; stops if fn returns true.
func (it *Collection[T]) ForEachBreak(fn func(index int, item T) (isBreak bool)) {
	for i, item := range it.items {
		if fn(i, item) {
			return
		}
	}
}

// Filter returns a new collection containing only items matching the predicate.
func (it *Collection[T]) Filter(predicate func(T) bool) *Collection[T] {
	result := EmptyCollection[T]()

	for _, item := range it.items {
		if predicate(item) {
			result.Add(item)
		}
	}

	return result
}

// CountFunc counts items matching the predicate.
func (it *Collection[T]) CountFunc(predicate func(T) bool) int {
	count := 0

	for _, item := range it.items {
		if predicate(item) {
			count++
		}
	}

	return count
}

// Clone creates a deep copy of the collection's slice.
func (it *Collection[T]) Clone() *Collection[T] {
	if it.IsEmpty() {
		return EmptyCollection[T]()
	}

	return CollectionClone[T](it.items)
}

// SortFunc sorts the collection in-place using the provided less function.
func (it *Collection[T]) SortFunc(less func(a, b T) bool) *Collection[T] {
	sort.Slice(it.items, func(i, j int) bool {
		return less(it.items[i], it.items[j])
	})

	return it
}

// Reverse reverses the collection in-place.
func (it *Collection[T]) Reverse() *Collection[T] {
	for i, j := 0, it.Length()-1; i < j; i, j = i+1, j-1 {
		it.items[i], it.items[j] = it.items[j], it.items[i]
	}

	return it
}

// String returns a string representation using fmt.
func (it *Collection[T]) String() string {
	return fmt.Sprintf("%v", it.items)
}

// ConcatNew creates a new collection by copying this one and appending additional items.
func (it *Collection[T]) ConcatNew(additionalItems ...T) *Collection[T] {
	newLen := it.Length() + len(additionalItems)
	result := NewCollection[T](newLen)
	result.AddSlice(it.items)
	result.Adds(additionalItems...)

	return result
}
