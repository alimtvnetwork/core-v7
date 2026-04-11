package coredynamic

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/defaultcapacity"
	"github.com/alimtvnetwork/core/pagingutil"
)

// Collection is a generic, type-safe collection that replaces
// DynamicCollection and AnyCollection with compile-time type safety.
//
// Deprecated alternatives:
//   - DynamicCollection — use Collection[Dynamic] instead
//   - AnyCollection — use Collection[any] instead
type Collection[T any] struct {
	items []T
	sync.RWMutex
}

// NewCollection creates a new Collection with the given initial capacity.
func NewCollection[T any](capacity int) *Collection[T] {
	return &Collection[T]{
		items: make([]T, 0, capacity),
	}
}

// EmptyCollection creates a new empty Collection with zero capacity.
func EmptyCollection[T any]() *Collection[T] {
	return NewCollection[T](0)
}

// CollectionFrom creates a Collection wrapping an existing slice (no copy).
func CollectionFrom[T any](items []T) *Collection[T] {
	if items == nil {
		items = []T{}
	}
	return &Collection[T]{items: items}
}

// CollectionClone creates a Collection by copying the given slice.
func CollectionClone[T any](items []T) *Collection[T] {
	length := len(items)
	cloned := make([]T, length, length+constants.Capacity4)
	copy(cloned, items)
	return &Collection[T]{items: cloned}
}

// --- Accessors ---

// At returns the item at the given index.
func (it *Collection[T]) At(index int) T {
	return it.items[index]
}

// First returns the first item. Panics if empty.
func (it *Collection[T]) First() T {
	return it.items[0]
}

// Last returns the last item. Panics if empty.
func (it *Collection[T]) Last() T {
	return it.items[it.LastIndex()]
}

// FirstOrDefault returns a pointer to the first item, or nil if empty.
func (it *Collection[T]) FirstOrDefault() (*T, bool) {
	if it.IsEmpty() {
		return nil, false
	}
	first := it.items[0]
	return &first, true
}

// LastOrDefault returns a pointer to the last item, or nil if empty.
func (it *Collection[T]) LastOrDefault() (*T, bool) {
	if it.IsEmpty() {
		return nil, false
	}
	last := it.items[it.LastIndex()]
	return &last, true
}

// Items returns the underlying slice.
func (it *Collection[T]) Items() []T {
	if it == nil || it.items == nil {
		return []T{}
	}
	return it.items
}

// --- Size ---

// Length returns the number of items.
func (it *Collection[T]) Length() int {
	if it == nil {
		return 0
	}
	return len(it.items)
}

// Count is an alias for Length.
func (it *Collection[T]) Count() int {
	return it.Length()
}

// IsEmpty returns true if the collection has no items.
func (it *Collection[T]) IsEmpty() bool {
	if it == nil {
		return true
	}
	return len(it.items) == 0
}

// HasAnyItem returns true if the collection has at least one item.
func (it *Collection[T]) HasAnyItem() bool {
	return !it.IsEmpty()
}

// LastIndex returns the index of the last item.
func (it *Collection[T]) LastIndex() int {
	return it.Length() - 1
}

// HasIndex returns true if the given index is valid.
func (it *Collection[T]) HasIndex(index int) bool {
	return index >= 0 && index <= it.LastIndex()
}

// --- Slicing ---

// Skip returns items after skipping the first n.
func (it *Collection[T]) Skip(n int) []T {
	return it.items[n:]
}

// Take returns the first n items.
func (it *Collection[T]) Take(n int) []T {
	return it.items[:n]
}

// Limit is an alias for Take.
func (it *Collection[T]) Limit(n int) []T {
	return it.Take(n)
}

// SkipCollection returns a new Collection after skipping the first n items.
func (it *Collection[T]) SkipCollection(n int) *Collection[T] {
	return &Collection[T]{items: it.items[n:]}
}

// TakeCollection returns a new Collection with the first n items.
func (it *Collection[T]) TakeCollection(n int) *Collection[T] {
	return &Collection[T]{items: it.items[:n]}
}

// LimitCollection is an alias for TakeCollection.
func (it *Collection[T]) LimitCollection(n int) *Collection[T] {
	return it.TakeCollection(n)
}

// SafeLimitCollection returns a new Collection limited to at most n items.
func (it *Collection[T]) SafeLimitCollection(limit int) *Collection[T] {
	limit = defaultcapacity.MaxLimit(it.Length(), limit)
	return &Collection[T]{items: it.items[:limit]}
}

// --- Mutators ---

// Add appends an item and returns the collection for chaining.
func (it *Collection[T]) Add(item T) *Collection[T] {
	it.items = append(it.items, item)
	return it
}

// AddMany appends multiple items.
func (it *Collection[T]) AddMany(items ...T) *Collection[T] {
	it.items = append(it.items, items...)
	return it
}

// AddNonNil appends a pointer-dereferenced item only if the pointer is non-nil.
func (it *Collection[T]) AddNonNil(item *T) *Collection[T] {
	if item == nil {
		return it
	}
	it.items = append(it.items, *item)
	return it
}

// RemoveAt removes the item at the given index. Returns false if index is invalid.
func (it *Collection[T]) RemoveAt(index int) bool {
	isInvalidIndex := !it.HasIndex(index)

	if isInvalidIndex {
		return false
	}
	it.items = append(it.items[:index], it.items[index+1:]...)
	return true
}

// Clear removes all items, keeping allocated capacity.
func (it *Collection[T]) Clear() {
	it.items = it.items[:0]
}

// Dispose removes all items and releases memory.
func (it *Collection[T]) Dispose() {
	it.items = nil
}

// --- Iteration ---

// Loop iterates over items. Return true from the callback to break.
func (it *Collection[T]) Loop(
	fn func(index int, item T) (isBreak bool),
) {
	if it.IsEmpty() {
		return
	}
	for i, item := range it.items {
		if fn(i, item) {
			return
		}
	}
}

// LoopAsync iterates over items concurrently. Break is not supported.
func (it *Collection[T]) LoopAsync(
	fn func(index int, item T),
) {
	if it.IsEmpty() {
		return
	}

	length := it.Length()
	wg := sync.WaitGroup{}
	wg.Add(length)

	for i := 0; i < length; i++ {
		go func(idx int) {
			fn(idx, it.items[idx])
			wg.Done()
		}(i)
	}

	wg.Wait()
}

// Filter returns a new Collection containing only items that match the predicate.
func (it *Collection[T]) Filter(
	predicate func(T) bool,
) *Collection[T] {
	if it.IsEmpty() {
		return EmptyCollection[T]()
	}

	result := NewCollection[T](it.Length() / 2)
	for _, item := range it.items {
		if predicate(item) {
			result.items = append(result.items, item)
		}
	}
	return result
}

// --- Paging ---

// GetPagesSize returns the number of pages for the given page size.
// Returns 0 if eachPageSize is zero or negative.
func (it *Collection[T]) GetPagesSize(eachPageSize int) int {
	if eachPageSize <= 0 {
		return 0
	}

	return int(math.Ceil(float64(it.Length()) / float64(eachPageSize)))
}

// GetPagingInfo returns paging metadata.
func (it *Collection[T]) GetPagingInfo(
	eachPageSize int,
	pageIndex int,
) pagingutil.PagingInfo {
	return pagingutil.GetPagingInfo(
		pagingutil.PagingRequest{
			Length:       it.Length(),
			PageIndex:    pageIndex,
			EachPageSize: eachPageSize,
		},
	)
}

// GetSinglePageCollection returns a single page. pageIndex is 1-based.
func (it *Collection[T]) GetSinglePageCollection(
	eachPageSize int,
	pageIndex int,
) *Collection[T] {
	length := it.Length()
	if length < eachPageSize {
		return it
	}

	pageInfo := it.GetPagingInfo(eachPageSize, pageIndex)
	return &Collection[T]{
		items: it.items[pageInfo.SkipItems:pageInfo.EndingLength],
	}
}

// GetPagedCollection splits the collection into pages of the given size.
func (it *Collection[T]) GetPagedCollection(
	eachPageSize int,
) []*Collection[T] {
	length := it.Length()
	if length < eachPageSize {
		return []*Collection[T]{it}
	}

	pageCount := it.GetPagesSize(eachPageSize)
	pages := make([]*Collection[T], pageCount)

	wg := sync.WaitGroup{}
	wg.Add(pageCount)

	for i := 1; i <= pageCount; i++ {
		go func(pageIndex int) {
			pages[pageIndex-1] = it.GetSinglePageCollection(eachPageSize, pageIndex)
			wg.Done()
		}(i)
	}

	wg.Wait()
	return pages
}

// --- Serialization ---

// MarshalJSON implements json.Marshaler.
func (it *Collection[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.items)
}

// UnmarshalJSON implements json.Unmarshaler.
func (it *Collection[T]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &it.items)
}

// JsonString returns the JSON string representation.
func (it *Collection[T]) JsonString() (string, error) {
	bytes, err := json.Marshal(it.items)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// JsonStringMust returns the JSON string or panics on error.
func (it *Collection[T]) JsonStringMust() string {
	s, err := it.JsonString()
	if err != nil {
		panic(err)
	}
	return s
}

// --- Stringer ---

// Strings returns a string representation of each item using fmt.Sprintf.
func (it *Collection[T]) Strings() []string {
	slice := make([]string, it.Length())
	for i, item := range it.items {
		slice[i] = fmt.Sprintf(constants.SprintValueFormat, item)
	}
	return slice
}

// String returns all items joined by newline.
func (it *Collection[T]) String() string {
	return strings.Join(it.Strings(), constants.NewLineUnix)
}
