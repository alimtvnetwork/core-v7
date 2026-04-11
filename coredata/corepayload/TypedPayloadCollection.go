package corepayload

import (
	"fmt"
	"sync"

	"github.com/alimtvnetwork/core/defaulterr"
)

// TypedPayloadCollection is a generic, thread-safe collection of TypedPayloadWrapper[T].
//
// It mirrors PayloadsCollection but provides compile-time type safety for the
// deserialized payload data. The embedded sync.RWMutex supports concurrent access
// via *Lock methods.
//
// Usage:
//
//	col := corepayload.NewTypedPayloadCollection[User](10)
//	col.Add(typedWrapper)
//	col.ForEach(func(index int, item *TypedPayloadWrapper[User]) {
//	    fmt.Println(item.Data().Name)
//	})
type TypedPayloadCollection[T any] struct {
	items []*TypedPayloadWrapper[T]
	sync.RWMutex
}

// EmptyTypedPayloadCollection creates a zero-capacity TypedPayloadCollection[T].
func EmptyTypedPayloadCollection[T any]() *TypedPayloadCollection[T] {
	return &TypedPayloadCollection[T]{
		items: []*TypedPayloadWrapper[T]{},
	}
}

// NewTypedPayloadCollection creates a TypedPayloadCollection[T] with pre-allocated capacity.
func NewTypedPayloadCollection[T any](capacity int) *TypedPayloadCollection[T] {
	return &TypedPayloadCollection[T]{
		items: make([]*TypedPayloadWrapper[T], 0, capacity),
	}
}

// TypedPayloadCollectionFrom wraps existing typed wrappers into a collection (no copy).
func TypedPayloadCollectionFrom[T any](items []*TypedPayloadWrapper[T]) *TypedPayloadCollection[T] {
	return &TypedPayloadCollection[T]{items: items}
}

// TypedPayloadCollectionFromPayloads converts a PayloadsCollection into a
// TypedPayloadCollection[T] by deserializing each wrapper.
//
// Items that fail deserialization are skipped.
func TypedPayloadCollectionFromPayloads[T any](
	payloads *PayloadsCollection,
) *TypedPayloadCollection[T] {
	if payloads == nil || payloads.IsEmpty() {
		return EmptyTypedPayloadCollection[T]()
	}

	collection := NewTypedPayloadCollection[T](payloads.Length())

	for _, wrapper := range payloads.Items {
		typed, err := NewTypedPayloadWrapper[T](wrapper)

		if err != nil {
			continue
		}

		collection.items = append(collection.items, typed)
	}

	return collection
}

// =============================================================================
// Core accessors
// =============================================================================

// Items returns the underlying slice of typed wrappers.
func (it *TypedPayloadCollection[T]) Items() []*TypedPayloadWrapper[T] {
	if it == nil {
		return nil
	}

	return it.items
}

// Length returns the number of items in the collection.
func (it *TypedPayloadCollection[T]) Length() int {
	if it == nil || it.items == nil {
		return 0
	}

	return len(it.items)
}

// LengthLock returns the length with mutex protection.
func (it *TypedPayloadCollection[T]) LengthLock() int {
	it.Lock()
	defer it.Unlock()

	return it.Length()
}

// Count returns the number of items (alias for Length).
func (it *TypedPayloadCollection[T]) Count() int {
	return it.Length()
}

// IsEmpty returns true if the collection has no items.
func (it *TypedPayloadCollection[T]) IsEmpty() bool {
	return it == nil || len(it.items) == 0
}

// IsEmptyLock returns IsEmpty with mutex protection.
func (it *TypedPayloadCollection[T]) IsEmptyLock() bool {
	it.Lock()
	defer it.Unlock()

	return it.IsEmpty()
}

// HasItems returns true if the collection has at least one item.
func (it *TypedPayloadCollection[T]) HasItems() bool {
	return it != nil && len(it.items) > 0
}

// HasAnyItem returns true if the collection has at least one item.
func (it *TypedPayloadCollection[T]) HasAnyItem() bool {
	return it.HasItems()
}

// LastIndex returns the index of the last element.
func (it *TypedPayloadCollection[T]) LastIndex() int {
	return it.Length() - 1
}

// HasIndex returns true if the given index is valid.
func (it *TypedPayloadCollection[T]) HasIndex(index int) bool {
	return index >= 0 && it.LastIndex() >= index
}

// =============================================================================
// Element access
// =============================================================================

// First returns the first element. Panics on empty collection.
func (it *TypedPayloadCollection[T]) First() *TypedPayloadWrapper[T] {
	return it.items[0]
}

// Last returns the last element. Panics on empty collection.
func (it *TypedPayloadCollection[T]) Last() *TypedPayloadWrapper[T] {
	return it.items[it.LastIndex()]
}

// FirstOrDefault returns the first element or nil if empty.
func (it *TypedPayloadCollection[T]) FirstOrDefault() *TypedPayloadWrapper[T] {
	if it.IsEmpty() {
		return nil
	}

	return it.First()
}

// LastOrDefault returns the last element or nil if empty.
func (it *TypedPayloadCollection[T]) LastOrDefault() *TypedPayloadWrapper[T] {
	if it.IsEmpty() {
		return nil
	}

	return it.Last()
}

// SafeAt returns the element at the given index, or nil if out of bounds.
func (it *TypedPayloadCollection[T]) SafeAt(index int) *TypedPayloadWrapper[T] {
	if it.IsEmpty() || !it.HasIndex(index) {
		return nil
	}

	return it.items[index]
}

// =============================================================================
// Mutation
// =============================================================================

// Add appends a single typed wrapper to the collection.
func (it *TypedPayloadCollection[T]) Add(item *TypedPayloadWrapper[T]) *TypedPayloadCollection[T] {
	it.items = append(it.items, item)

	return it
}

// AddLock appends a single typed wrapper with mutex protection.
func (it *TypedPayloadCollection[T]) AddLock(item *TypedPayloadWrapper[T]) *TypedPayloadCollection[T] {
	it.Lock()
	defer it.Unlock()

	it.items = append(it.items, item)

	return it
}

// Adds appends multiple typed wrappers.
func (it *TypedPayloadCollection[T]) Adds(items ...*TypedPayloadWrapper[T]) *TypedPayloadCollection[T] {
	it.items = append(it.items, items...)

	return it
}

// AddCollection appends all items from another TypedPayloadCollection[T].
func (it *TypedPayloadCollection[T]) AddCollection(
	other *TypedPayloadCollection[T],
) *TypedPayloadCollection[T] {
	if other.IsEmpty() {
		return it
	}

	it.items = append(it.items, other.items...)

	return it
}

// RemoveAt removes the item at the given index. Returns true on success.
func (it *TypedPayloadCollection[T]) RemoveAt(index int) bool {
	length := it.Length()

	if index < 0 || index >= length {
		return false
	}

	it.items = append(it.items[:index], it.items[index+1:]...)

	return true
}

// =============================================================================
// Iteration
// =============================================================================

// ForEach calls fn for each typed wrapper with its index.
func (it *TypedPayloadCollection[T]) ForEach(
	fn func(index int, item *TypedPayloadWrapper[T]),
) {
	for i, item := range it.items {
		fn(i, item)
	}
}

// ForEachData calls fn for each deserialized data item with its index.
func (it *TypedPayloadCollection[T]) ForEachData(
	fn func(index int, data T),
) {
	for i, item := range it.items {
		fn(i, item.Data())
	}
}

// ForEachBreak calls fn for each item; stops if fn returns true.
func (it *TypedPayloadCollection[T]) ForEachBreak(
	fn func(index int, item *TypedPayloadWrapper[T]) (isBreak bool),
) {
	for i, item := range it.items {
		if fn(i, item) {
			return
		}
	}
}

// =============================================================================
// Filter
// =============================================================================

// Filter returns a new collection containing only items matching the predicate.
func (it *TypedPayloadCollection[T]) Filter(
	predicate func(*TypedPayloadWrapper[T]) bool,
) *TypedPayloadCollection[T] {
	result := EmptyTypedPayloadCollection[T]()

	for _, item := range it.items {
		if predicate(item) {
			result.Add(item)
		}
	}

	return result
}

// FilterByData returns a new collection containing only items
// whose typed data matches the predicate.
func (it *TypedPayloadCollection[T]) FilterByData(
	predicate func(T) bool,
) *TypedPayloadCollection[T] {
	result := EmptyTypedPayloadCollection[T]()

	for _, item := range it.items {
		if predicate(item.Data()) {
			result.Add(item)
		}
	}

	return result
}

// FirstByFilter returns the first item matching the predicate, or nil.
func (it *TypedPayloadCollection[T]) FirstByFilter(
	predicate func(*TypedPayloadWrapper[T]) bool,
) *TypedPayloadWrapper[T] {
	for _, item := range it.items {
		if predicate(item) {
			return item
		}
	}

	return nil
}

// FirstByData returns the first item whose data matches the predicate, or nil.
func (it *TypedPayloadCollection[T]) FirstByData(
	predicate func(T) bool,
) *TypedPayloadWrapper[T] {
	for _, item := range it.items {
		if predicate(item.Data()) {
			return item
		}
	}

	return nil
}

// FirstByName returns the first item matching the given name, or nil.
func (it *TypedPayloadCollection[T]) FirstByName(name string) *TypedPayloadWrapper[T] {
	return it.FirstByFilter(func(item *TypedPayloadWrapper[T]) bool {
		return item.Name() == name
	})
}

// FirstById returns the first item matching the given identifier, or nil.
func (it *TypedPayloadCollection[T]) FirstById(id string) *TypedPayloadWrapper[T] {
	return it.FirstByFilter(func(item *TypedPayloadWrapper[T]) bool {
		return item.Identifier() == id
	})
}

// CountFunc counts items matching the predicate.
func (it *TypedPayloadCollection[T]) CountFunc(
	predicate func(*TypedPayloadWrapper[T]) bool,
) int {
	count := 0

	for _, item := range it.items {
		if predicate(item) {
			count++
		}
	}

	return count
}

// =============================================================================
// Slice operations
// =============================================================================

// Skip returns items after skipping the first n.
func (it *TypedPayloadCollection[T]) Skip(count int) []*TypedPayloadWrapper[T] {
	if count >= it.Length() {
		return []*TypedPayloadWrapper[T]{}
	}

	return it.items[count:]
}

// Take returns the first n items.
func (it *TypedPayloadCollection[T]) Take(count int) []*TypedPayloadWrapper[T] {
	if count >= it.Length() {
		return it.items
	}

	return it.items[:count]
}

// =============================================================================
// Extraction
// =============================================================================

// AllData extracts typed data from all wrappers into a slice.
func (it *TypedPayloadCollection[T]) AllData() []T {
	if it.IsEmpty() {
		return []T{}
	}

	result := make([]T, it.Length())

	for i, item := range it.items {
		result[i] = item.Data()
	}

	return result
}

// AllNames extracts names from all wrappers.
func (it *TypedPayloadCollection[T]) AllNames() []string {
	if it.IsEmpty() {
		return []string{}
	}

	result := make([]string, it.Length())

	for i, item := range it.items {
		result[i] = item.Name()
	}

	return result
}

// AllIdentifiers extracts identifiers from all wrappers.
func (it *TypedPayloadCollection[T]) AllIdentifiers() []string {
	if it.IsEmpty() {
		return []string{}
	}

	result := make([]string, it.Length())

	for i, item := range it.items {
		result[i] = item.Identifier()
	}

	return result
}

// =============================================================================
// Conversion
// =============================================================================

// ToPayloadsCollection converts back to a non-generic PayloadsCollection.
func (it *TypedPayloadCollection[T]) ToPayloadsCollection() *PayloadsCollection {
	if it.IsEmpty() {
		return &PayloadsCollection{Items: []*PayloadWrapper{}}
	}

	wrappers := make([]*PayloadWrapper, it.Length())

	for i, item := range it.items {
		wrappers[i] = item.ToPayloadWrapper()
	}

	return &PayloadsCollection{Items: wrappers}
}

// Clone creates a deep copy of the collection by cloning each wrapper.
func (it *TypedPayloadCollection[T]) Clone() (*TypedPayloadCollection[T], error) {
	if it.IsEmpty() {
		return EmptyTypedPayloadCollection[T](), nil
	}

	cloned := NewTypedPayloadCollection[T](it.Length())

	for _, item := range it.items {
		clonedItem, err := item.ClonePtr(true)

		if err != nil {
			return nil, err
		}

		cloned.items = append(cloned.items, clonedItem)
	}

	return cloned, nil
}

// CloneMust creates a deep copy or panics on error.
func (it *TypedPayloadCollection[T]) CloneMust() *TypedPayloadCollection[T] {
	cloned, err := it.Clone()

	if err != nil {
		panic(err)
	}

	return cloned
}

// ConcatNew creates a new collection by cloning this one and appending additional items.
func (it *TypedPayloadCollection[T]) ConcatNew(
	additionalItems ...*TypedPayloadWrapper[T],
) (*TypedPayloadCollection[T], error) {
	cloned, err := it.Clone()

	if err != nil {
		return nil, err
	}

	cloned.Adds(additionalItems...)

	return cloned, nil
}

// Clear removes all items and resets the collection.
func (it *TypedPayloadCollection[T]) Clear() {
	if it == nil {
		return
	}

	it.items = []*TypedPayloadWrapper[T]{}
}

// Dispose clears and nils the underlying slice.
func (it *TypedPayloadCollection[T]) Dispose() {
	if it == nil {
		return
	}

	it.items = nil
}

// =============================================================================
// Deserialization
// =============================================================================

// TypedPayloadCollectionDeserialize deserializes raw JSON bytes containing
// an array of PayloadWrappers into a TypedPayloadCollection[T].
func TypedPayloadCollectionDeserialize[T any](
	rawBytes []byte,
) (*TypedPayloadCollection[T], error) {
	typedWrappers, err := TypedPayloadWrapperDeserializeToMany[T](rawBytes)

	if err != nil {
		return nil, err
	}

	return TypedPayloadCollectionFrom[T](typedWrappers), nil
}

// TypedPayloadCollectionDeserializeMust deserializes or panics.
func TypedPayloadCollectionDeserializeMust[T any](
	rawBytes []byte,
) *TypedPayloadCollection[T] {
	collection, err := TypedPayloadCollectionDeserialize[T](rawBytes)

	if err != nil {
		panic(err)
	}

	return collection
}

// NewTypedPayloadCollectionSingle creates a collection containing a single item.
func NewTypedPayloadCollectionSingle[T any](
	item *TypedPayloadWrapper[T],
) *TypedPayloadCollection[T] {
	if item == nil {
		return EmptyTypedPayloadCollection[T]()
	}

	return &TypedPayloadCollection[T]{
		items: []*TypedPayloadWrapper[T]{item},
	}
}

// NewTypedPayloadCollectionFromData creates a collection from typed data values,
// wrapping each into a TypedPayloadWrapper[T] with auto-detected entity type.
func NewTypedPayloadCollectionFromData[T any](
	name string,
	dataItems []T,
) (*TypedPayloadCollection[T], error) {
	if len(dataItems) == 0 {
		return EmptyTypedPayloadCollection[T](), nil
	}

	collection := NewTypedPayloadCollection[T](len(dataItems))

	for i, data := range dataItems {
		identifier := fmt.Sprintf("%s-%d", name, i)
		typed, err := TypedPayloadWrapperNameIdRecord[T](name, identifier, data)

		if err != nil {
			return nil, err
		}

		collection.items = append(collection.items, typed)
	}

	return collection, nil
}

// NewTypedPayloadCollectionFromDataMust creates a collection from data or panics.
func NewTypedPayloadCollectionFromDataMust[T any](
	name string,
	dataItems []T,
) *TypedPayloadCollection[T] {
	collection, err := NewTypedPayloadCollectionFromData[T](name, dataItems)

	if err != nil {
		panic(err)
	}

	return collection
}

// IsValid returns true if all items are parsed and non-nil.
func (it *TypedPayloadCollection[T]) IsValid() bool {
	if it.IsEmpty() {
		return true
	}

	for _, item := range it.items {
		if item == nil || !item.IsParsed() {
			return false
		}
	}

	return true
}

// HasErrors returns true if any item has an error.
func (it *TypedPayloadCollection[T]) HasErrors() bool {
	for _, item := range it.items {
		if item != nil && item.HasError() {
			return true
		}
	}

	return false
}

// Errors returns all errors from items that have them.
func (it *TypedPayloadCollection[T]) Errors() []error {
	if it.IsEmpty() {
		return nil
	}

	var errs []error

	for _, item := range it.items {
		if item != nil && item.HasError() {
			itemErr := item.Error()

			if itemErr != nil {
				errs = append(errs, itemErr)
			}
		}
	}

	return errs
}

// FirstError returns the first error found, or nil.
func (it *TypedPayloadCollection[T]) FirstError() error {
	for _, item := range it.items {
		if item != nil && item.HasError() {
			return item.Error()
		}
	}

	return nil
}

// MergedError returns a single merged error from all item errors.
func (it *TypedPayloadCollection[T]) MergedError() error {
	errs := it.Errors()

	if len(errs) == 0 {
		return nil
	}

	if len(errs) == 1 {
		return errs[0]
	}

	return defaulterr.Marshalling // placeholder: use errcore.MergeErrors in production
}
