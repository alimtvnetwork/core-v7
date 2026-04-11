package coregeneric

import (
	"fmt"
	"sync"
)

// LinkedList is a generic singly-linked list with head/tail pointers and embedded mutex.
// It generalizes corestr.LinkedList from string to any type T.
type LinkedList[T any] struct {
	head, tail *LinkedListNode[T]
	length     int
	sync.RWMutex
}

// EmptyLinkedList creates an empty LinkedList[T].
func EmptyLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// LinkedListFrom creates a LinkedList[T] from a slice of items.
func LinkedListFrom[T any](items []T) *LinkedList[T] {
	ll := EmptyLinkedList[T]()

	for _, item := range items {
		ll.Add(item)
	}

	return ll
}

// Head returns the head node.
func (it *LinkedList[T]) Head() *LinkedListNode[T] {
	return it.head
}

// Tail returns the tail node.
func (it *LinkedList[T]) Tail() *LinkedListNode[T] {
	return it.tail
}

// Length returns the number of nodes.
func (it *LinkedList[T]) Length() int {
	return it.length
}

// LengthLock returns the length with mutex protection.
func (it *LinkedList[T]) LengthLock() int {
	it.RLock()
	defer it.RUnlock()

	return it.length
}

// IsEmpty returns true if the linked list has no nodes.
func (it *LinkedList[T]) IsEmpty() bool {
	return it == nil || it.head == nil || it.length == 0
}

// IsEmptyLock returns IsEmpty with mutex protection.
func (it *LinkedList[T]) IsEmptyLock() bool {
	it.RLock()
	defer it.RUnlock()

	return it.IsEmpty()
}

// HasItems returns true if the linked list has at least one node.
func (it *LinkedList[T]) HasItems() bool {
	return it.head != nil && it.length > 0
}

// Add appends an item to the back of the list.
func (it *LinkedList[T]) Add(item T) *LinkedList[T] {
	node := &LinkedListNode[T]{
		Element: item,
		next:    nil,
	}

	if it.IsEmpty() {
		it.head = node
		it.tail = it.head
		it.length++

		return it
	}

	it.tail.next = node
	it.tail = it.tail.next
	it.length++

	return it
}

// AddLock appends an item with mutex protection.
func (it *LinkedList[T]) AddLock(item T) *LinkedList[T] {
	it.Lock()
	defer it.Unlock()

	return it.Add(item)
}

// Adds appends multiple items.
func (it *LinkedList[T]) Adds(items ...T) *LinkedList[T] {
	for _, item := range items {
		it.Add(item)
	}

	return it
}

// AddSlice appends all items from a slice.
func (it *LinkedList[T]) AddSlice(items []T) *LinkedList[T] {
	for _, item := range items {
		it.Add(item)
	}

	return it
}

// AddIf appends the item only if the condition is true.
func (it *LinkedList[T]) AddIf(isAdd bool, item T) *LinkedList[T] {
	isSkip := !isAdd

	if isSkip {
		return it
	}

	return it.Add(item)
}

// AddsIf appends items only if the condition is true.
func (it *LinkedList[T]) AddsIf(isAdd bool, items ...T) *LinkedList[T] {
	isSkip := !isAdd

	if isSkip {
		return it
	}

	return it.Adds(items...)
}

// AddFunc appends the result of calling the function.
func (it *LinkedList[T]) AddFunc(f func() T) *LinkedList[T] {
	return it.Add(f())
}

// AddFront prepends an item to the front of the list.
func (it *LinkedList[T]) AddFront(item T) *LinkedList[T] {
	if it.IsEmpty() {
		return it.Add(item)
	}

	node := &LinkedListNode[T]{
		Element: item,
		next:    it.head,
	}

	it.head = node
	it.length++

	return it
}

// PushBack is an alias for Add.
func (it *LinkedList[T]) PushBack(item T) *LinkedList[T] {
	return it.Add(item)
}

// PushFront is an alias for AddFront.
func (it *LinkedList[T]) PushFront(item T) *LinkedList[T] {
	return it.AddFront(item)
}

// Push is an alias for Add.
func (it *LinkedList[T]) Push(item T) *LinkedList[T] {
	return it.Add(item)
}

// AppendNode appends an existing node to the back.
func (it *LinkedList[T]) AppendNode(node *LinkedListNode[T]) *LinkedList[T] {
	if it.IsEmpty() {
		it.head = node
		it.tail = it.head
		it.length++

		return it
	}

	it.tail.next = node
	it.tail = it.tail.next
	it.length++

	return it
}

// AppendChainOfNodes appends a chain of linked nodes.
func (it *LinkedList[T]) AppendChainOfNodes(nodeHead *LinkedListNode[T]) *LinkedList[T] {
	endOfChain, chainLen := nodeHead.EndOfChain()

	if it.IsEmpty() {
		it.head = nodeHead
	} else {
		it.tail.next = nodeHead
	}

	it.tail = endOfChain
	it.length += chainLen

	return it
}

// First returns the first element. Panics on empty list.
func (it *LinkedList[T]) First() T {
	return it.head.Element
}

// Last returns the last element. Panics on empty list.
func (it *LinkedList[T]) Last() T {
	return it.tail.Element
}

// FirstOrDefault returns the first element or zero value if empty.
func (it *LinkedList[T]) FirstOrDefault() T {
	if it.IsEmpty() {
		var zero T

		return zero
	}

	return it.head.Element
}

// LastOrDefault returns the last element or zero value if empty.
func (it *LinkedList[T]) LastOrDefault() T {
	if it.IsEmpty() {
		var zero T

		return zero
	}

	return it.tail.Element
}

// Items collects all elements into a slice.
func (it *LinkedList[T]) Items() []T {
	if it.IsEmpty() {
		return []T{}
	}

	items := make([]T, 0, it.length)
	node := it.head
	items = append(items, node.Element)

	for node.HasNext() {
		node = node.Next()
		items = append(items, node.Element)
	}

	return items
}

// Collection converts the LinkedList to a Collection[T].
func (it *LinkedList[T]) Collection() *Collection[T] {
	return CollectionFrom[T](it.Items())
}

// ForEach calls fn for each element with its index.
func (it *LinkedList[T]) ForEach(fn func(index int, item T)) {
	if it.IsEmpty() {
		return
	}

	node := it.head
	fn(0, node.Element)

	i := 1
	for node.HasNext() {
		node = node.Next()
		fn(i, node.Element)
		i++
	}
}

// ForEachBreak calls fn for each element; stops if fn returns true.
func (it *LinkedList[T]) ForEachBreak(fn func(index int, item T) (isBreak bool)) {
	if it.IsEmpty() {
		return
	}

	node := it.head
	if fn(0, node.Element) {
		return
	}

	i := 1
	for node.HasNext() {
		node = node.Next()

		if fn(i, node.Element) {
			return
		}

		i++
	}
}

// IndexAt returns the node at the given index. BigO(n).
func (it *LinkedList[T]) IndexAt(index int) *LinkedListNode[T] {
	if it.IsEmpty() || index < 0 || index >= it.length {
		return nil
	}

	node := it.head

	for i := 0; i < index; i++ {
		isEndOfList := !node.HasNext()

		if isEndOfList {
			return nil
		}

		node = node.Next()
	}

	return node
}

// String returns a string representation.
func (it *LinkedList[T]) String() string {
	return fmt.Sprintf("%v", it.Items())
}
