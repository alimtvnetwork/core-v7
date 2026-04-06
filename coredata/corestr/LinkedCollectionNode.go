package corestr

import (
	"fmt"
	"strings"

	"github.com/alimtvnetwork/core/constants"
)

type LinkedCollectionNode struct {
	Element *Collection
	next    *LinkedCollectionNode
}

func (it *LinkedCollectionNode) IsEmpty() bool {
	return it == nil || it.Element == nil
}

func (it *LinkedCollectionNode) HasElement() bool {
	return it.Element != nil
}

func (it *LinkedCollectionNode) HasNext() bool {
	return it.next != nil
}

func (it *LinkedCollectionNode) Next() *LinkedCollectionNode {
	return it.next
}

func (it *LinkedCollectionNode) AddNext(
	linkedCollection *LinkedCollections,
	collection *Collection,
) *LinkedCollectionNode {
	newNode := &LinkedCollectionNode{
		Element: collection,
		next:    it.Next(),
	}

	it.next = newNode

	linkedCollection.incrementLength()

	return newNode
}

func (it *LinkedCollectionNode) AddStringsToNode(
	linkedCollection *LinkedCollections,
	isSkipOnNull bool,
	items []string,
	isMakeClone bool,
) *LinkedCollections {
	collection := New.
		Collection.
		StringsOptions(isMakeClone, items)

	return linkedCollection.AddCollectionToNode(
		isSkipOnNull,
		it,
		collection,
	)
}

func (it *LinkedCollectionNode) AddCollectionToNode(
	linkedCollection *LinkedCollections,
	isSkipOnNull bool,
	collection *Collection,
) *LinkedCollections {
	return linkedCollection.AddCollectionToNode(
		isSkipOnNull,
		it,
		collection,
	)
}

func (it *LinkedCollectionNode) AddNextNode(
	linkedCollection *LinkedCollections,
	nextNode *LinkedCollectionNode,
) *LinkedCollectionNode {
	nextNode.next = it.Next()
	it.next = nextNode

	linkedCollection.incrementLength()

	return nextNode
}

func (it *LinkedCollectionNode) IsChainEqual(another *LinkedCollectionNode) bool {
	if it == another {
		return true
	}

	if another == nil && it == nil {
		return true
	}

	if another == nil || it == nil {
		return false
	}

	return it.IsEqual(another) &&
		it.isNextChainEqual(another)
}

func (it *LinkedCollectionNode) IsEqual(another *LinkedCollectionNode) bool {
	if it == nil && nil == another {
		return true
	}

	if it == nil || nil == another {
		return false
	}

	if it == another {
		return true
	}

	//goland:noinspection GoNilness

	elem1 := it.Element
	elem2 := another.Element

	//goland:noinspection GoNilness
	if elem1 == nil && nil == elem2 {
		return true
	}

	if elem1 == nil || nil == elem2 {
		return false
	}

	if elem1 == elem2 {
		return true
	}

	isElementSame := elem1.IsEquals(elem2)

	return isElementSame &&
		it.isNextEqual(another)
}

func (it *LinkedCollectionNode) isNextEqual(
	another *LinkedCollectionNode,
) bool {
	next1 := it.Next()
	next2 := another.Next()

	if next1 == nil && nil == next2 {
		return true
	}

	if next1 == nil || nil == next2 {
		return false
	}

	if next1 == next2 {
		return true
	}

	return next1.
		Element.
		IsEquals(
			next2.Element,
		)
}

func (it *LinkedCollectionNode) isNextChainEqual(
	another *LinkedCollectionNode,
) bool {
	next1 := it.Next()
	next2 := another.Next()

	if next1 == nil && nil == next2 {
		return true
	}

	if next1 == nil || nil == next2 {
		return false
	}

	return next1.IsChainEqual(next2)
}

func (it *LinkedCollectionNode) IsEqualValue(collection *Collection) bool {
	elem1 := it.Element

	//goland:noinspection GoNilness
	if elem1 == nil && nil == collection {
		return true
	}

	if elem1 == nil || nil == collection {
		return false
	}

	if elem1 == collection {
		return true
	}

	return elem1.IsEquals(collection)
}

func (it *LinkedCollectionNode) EndOfChain() (
	endOfChain *LinkedCollectionNode,
	length int,
) {
	node := it
	length++

	for node.HasNext() {
		node = node.Next()
		length++
	}

	return node, length
}

func (it *LinkedCollectionNode) LoopEndOfChain(
	processor LinkedCollectionSimpleProcessor,
) (endOfLoop *LinkedCollectionNode, length int) {
	node := it
	arg := &LinkedCollectionProcessorParameter{
		Index:         0,
		CurrentNode:   node,
		PrevNode:      nil,
		IsFirstIndex:  true,
		IsEndingIndex: false,
	}

	isBreak := processor(arg)

	length++

	if isBreak {
		return node, length
	}

	i := 1

	for node.HasNext() {
		prev := node
		node = node.Next()
		isEndingIndex := !node.HasNext()
		arg2 := &LinkedCollectionProcessorParameter{
			Index:         i,
			CurrentNode:   node,
			PrevNode:      prev,
			IsFirstIndex:  false,
			IsEndingIndex: isEndingIndex,
		}

		isBreak = processor(arg2)
		length++
		i++

		if isBreak {
			return node, length
		}
	}

	return node, length
}

func (it *LinkedCollectionNode) CreateLinkedList() *LinkedCollections {
	return Empty.LinkedCollections().
		AppendChainOfNodes(it)
}

func (it *LinkedCollectionNode) Clone() *LinkedCollectionNode {
	return &LinkedCollectionNode{
		Element: it.Element,
		next:    nil,
	}
}

func (it *LinkedCollectionNode) String() string {
	return it.Element.String()
}

// List returns all elements from this node onwards as a string slice.
func (it *LinkedCollectionNode) List() []string {
	list := make([]string, 0, constants.ArbitraryCapacity100)

	node := it
	list = append(list, node.Element.List()...)

	for node.HasNext() {
		node = node.Next()

		list = append(list, node.Element.List()...)
	}

	return list
}

// Deprecated: Use List instead.
func (it *LinkedCollectionNode) ListPtr() *[]string {
	list := it.List()
	return &list
}

func (it *LinkedCollectionNode) Join(separator string) string {
	list := it.List()
	return strings.Join(list, separator)
}

func (it *LinkedCollectionNode) StringList(header string) string {
	return header + it.Join(commonJoiner)
}

func (it *LinkedCollectionNode) Print(header string) {
	finalString := it.StringList(header)
	fmt.Println(finalString)
}
