package namevalue

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/alimtvnetwork/core/constants"
)

// Collection is a generic collection of Instance[K, V] items.
// It provides add, prepend, append, join, clone, and serialization operations.
type Collection[K comparable, V any] struct {
	Items        []Instance[K, V] `json:"Items,omitempty"`
	lazyToString *string
}

// NewCollection creates a new Collection[K, V] with the given capacity.
func NewGenericCollection[K comparable, V any](capacity int) *Collection[K, V] {
	slice := make([]Instance[K, V], 0, capacity)

	return &Collection[K, V]{
		Items: slice,
	}
}

// NewGenericCollectionDefault creates a new Collection[K, V] with default capacity.
func NewGenericCollectionDefault[K comparable, V any]() *Collection[K, V] {
	return NewGenericCollection[K, V](constants.Capacity5)
}

// EmptyGenericCollection creates an empty Collection[K, V].
func EmptyGenericCollection[K comparable, V any]() *Collection[K, V] {
	return NewGenericCollection[K, V](0)
}

// NewGenericCollectionUsing creates a Collection[K, V] from existing items.
func NewGenericCollectionUsing[K comparable, V any](
	isClone bool,
	items ...Instance[K, V],
) *Collection[K, V] {
	if items == nil {
		return EmptyGenericCollection[K, V]()
	}

	isSkipClone := !isClone

	if isSkipClone {
		return &Collection[K, V]{
			Items: items,
		}
	}

	slice := NewGenericCollection[K, V](len(items))

	return slice.Adds(items...)
}

func (it *Collection[K, V]) Add(
	item Instance[K, V],
) *Collection[K, V] {
	it.InvalidateLazyString()
	it.Items = append(it.Items, item)

	return it
}

func (it *Collection[K, V]) Adds(
	items ...Instance[K, V],
) *Collection[K, V] {
	if len(items) == 0 {
		return it
	}

	it.InvalidateLazyString()
	it.Items = append(
		it.Items,
		items...)

	return it
}

func (it *Collection[K, V]) Append(
	items ...Instance[K, V],
) *Collection[K, V] {
	if len(items) == 0 {
		return it
	}

	it.InvalidateLazyString()
	it.Items = append(
		it.Items,
		items...)

	return it
}

func (it *Collection[K, V]) AppendIf(
	isAppend bool,
	items ...Instance[K, V],
) *Collection[K, V] {
	if !isAppend || len(items) == 0 {
		return it
	}

	it.InvalidateLazyString()
	it.Items = append(
		it.Items,
		items...)

	return it
}

func (it *Collection[K, V]) Prepend(
	items ...Instance[K, V],
) *Collection[K, V] {
	if len(items) == 0 {
		return it
	}

	it.InvalidateLazyString()
	it.Items = append(
		items,
		it.Items...)

	return it
}

func (it *Collection[K, V]) PrependIf(
	isPrepend bool,
	items ...Instance[K, V],
) *Collection[K, V] {
	if !isPrepend || len(items) == 0 {
		return it
	}

	it.InvalidateLazyString()
	it.Items = append(
		items,
		it.Items...)

	return it
}

func (it *Collection[K, V]) PrependUsingFuncIf(
	isPrepend bool,
	itemsGetterFunc func() []Instance[K, V],
) *Collection[K, V] {
	if !isPrepend || itemsGetterFunc == nil {
		return it
	}

	it.InvalidateLazyString()
	it.Items = append(
		itemsGetterFunc(),
		it.Items...)

	return it
}

func (it *Collection[K, V]) AppendUsingFuncIf(
	isAppend bool,
	itemsGetterFunc func() []Instance[K, V],
) *Collection[K, V] {
	if !isAppend || itemsGetterFunc == nil {
		return it
	}

	it.InvalidateLazyString()
	it.Items = append(
		it.Items,
		itemsGetterFunc()...,
	)

	return it
}

func (it *Collection[K, V]) AppendPrependIf(
	isAppendOrPrepend bool,
	prependItems []Instance[K, V],
	appendItems []Instance[K, V],
) *Collection[K, V] {
	isSkip := !isAppendOrPrepend

	if isSkip {
		return it
	}

	if len(prependItems) > 0 {
		it.InvalidateLazyString()
		it.Items = append(
			prependItems,
			it.Items...)
	}

	if len(appendItems) > 0 {
		it.InvalidateLazyString()
		it.Items = append(
			it.Items,
			appendItems...)
	}

	return it
}

func (it *Collection[K, V]) AddsPtr(
	items ...*Instance[K, V],
) *Collection[K, V] {
	if len(items) == 0 {
		return it
	}

	for _, item := range items {
		if item == nil {
			continue
		}

		it.Items = append(
			it.Items,
			*item)
	}

	return it
}

func (it *Collection[K, V]) HasCompiledString() bool {
	return it != nil && it.lazyToString != nil
}

func (it *Collection[K, V]) InvalidateLazyString() {
	if it == nil {
		return
	}

	it.lazyToString = nil
}

func (it *Collection[K, V]) CompiledLazyString() string {
	if it == nil {
		return constants.EmptyString
	}

	if it.lazyToString != nil {
		return *it.lazyToString
	}

	toString := it.String()
	it.lazyToString = &toString

	return toString
}

func (it *Collection[K, V]) ConcatNew(
	additionalItems ...Instance[K, V],
) *Collection[K, V] {
	cloned := it.Clone()

	return cloned.Adds(additionalItems...)
}

func (it *Collection[K, V]) ConcatNewPtr(
	additionalItems ...*Instance[K, V],
) *Collection[K, V] {
	cloned := it.Clone()

	return cloned.AddsPtr(
		additionalItems...)
}

func (it *Collection[K, V]) AddsIf(
	isAdd bool,
	items ...Instance[K, V],
) *Collection[K, V] {
	isSkip := !isAdd

	if isSkip {
		return it
	}

	it.InvalidateLazyString()

	return it.Adds(items...)
}

func (it *Collection[K, V]) Length() int {
	if it == nil {
		return 0
	}

	return len(it.Items)
}

func (it *Collection[K, V]) Count() int {
	return it.Length()
}

func (it *Collection[K, V]) IsEmpty() bool {
	return it.Length() == 0
}

func (it *Collection[K, V]) HasAnyItem() bool {
	return !it.IsEmpty()
}

func (it *Collection[K, V]) LastIndex() int {
	return it.Length() - 1
}

func (it *Collection[K, V]) HasIndex(index int) bool {
	return it.LastIndex() >= index
}

func (it *Collection[K, V]) Strings() []string {
	list := make([]string, it.Length())

	for i, item := range it.Items {
		list[i] = item.String()
	}

	return list
}

func (it *Collection[K, V]) JsonStrings() []string {
	list := make([]string, it.Length())

	for i, item := range it.Items {
		list[i] = item.JsonString()
	}

	return list
}

func (it *Collection[K, V]) JoinJsonStrings(joiner string) string {
	return strings.Join(it.JsonStrings(), joiner)
}

func (it *Collection[K, V]) Join(joiner string) string {
	return strings.Join(it.Strings(), joiner)
}

func (it *Collection[K, V]) JoinLines() string {
	return strings.Join(it.Strings(), constants.DefaultLine)
}

func (it *Collection[K, V]) JoinCsv() string {
	return strings.Join(it.CsvStrings(), constants.Comma)
}

func (it *Collection[K, V]) JoinCsvLine() string {
	return strings.Join(it.CsvStrings(), constants.CommaUnixNewLine)
}

// IsEqualByString compares two collections by their String() output.
// This is used instead of direct struct comparison because V may not be comparable.
func (it *Collection[K, V]) IsEqualByString(another *Collection[K, V]) bool {
	if it == nil && another == nil {
		return true
	}

	if it == nil || another == nil {
		return false
	}

	if it.Length() != another.Length() {
		return false
	}

	for i, item := range it.Items {
		if item.String() != another.Items[i].String() {
			return false
		}
	}

	return true
}

func (it Collection[K, V]) JsonString() string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	jsonBytes, err := json.Marshal(it)

	if err != nil || len(jsonBytes) == 0 {
		return constants.EmptyString
	}

	return string(jsonBytes)
}

func (it *Collection[K, V]) String() string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	if it.HasCompiledString() {
		return *it.lazyToString
	}

	return it.JoinLines()
}

func (it *Collection[K, V]) Error() error {
	if it.IsEmpty() {
		return nil
	}

	return errors.New(it.String())
}

func (it *Collection[K, V]) ErrorUsingMessage(message string) error {
	if it.IsEmpty() {
		return nil
	}

	toCompiled := message + constants.Space + it.String()

	return errors.New(toCompiled)
}

func (it *Collection[K, V]) CsvStrings() []string {
	if it.IsEmpty() {
		return []string{}
	}

	newSlice := make([]string, it.Length())

	for i, item := range it.Items {
		newSlice[i] = fmt.Sprintf(
			constants.SprintDoubleQuoteFormat,
			item.String())
	}

	return newSlice
}

func (it *Collection[K, V]) Clear() *Collection[K, V] {
	if it == nil {
		return it
	}

	tempItems := it.Items
	clearFunc := func() {
		for _, item := range tempItems {
			item.Dispose()
		}
	}

	go clearFunc()

	it.Items = []Instance[K, V]{}
	it.lazyToString = nil

	return it
}

func (it *Collection[K, V]) Dispose() {
	if it == nil {
		return
	}

	it.Clear()
	it.Items = nil
}

func (it *Collection[K, V]) Clone() *Collection[K, V] {
	list := NewGenericCollection[K, V](it.Length())

	return list.Adds(it.Items...)
}

func (it *Collection[K, V]) ClonePtr() *Collection[K, V] {
	if it == nil {
		return nil
	}

	return it.Clone()
}
