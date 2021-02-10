package corestr

import (
	"fmt"
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

type LinkedCollectionNode struct {
	Element *Collection
	next    *LinkedCollectionNode
}

func (linkedCollectionNode *LinkedCollectionNode) IsEmpty() bool {
	return linkedCollectionNode.Element == nil
}

func (linkedCollectionNode *LinkedCollectionNode) HasElement() bool {
	return linkedCollectionNode.Element != nil
}

func (linkedCollectionNode *LinkedCollectionNode) HasNext() bool {
	return linkedCollectionNode.next != nil
}

func (linkedCollectionNode *LinkedCollectionNode) Next() *LinkedCollectionNode {
	return linkedCollectionNode.next
}

func (linkedCollectionNode *LinkedCollectionNode) AddNext(
	linkedCollection *LinkedCollections,
	collection *Collection,
) *LinkedCollectionNode {
	newNode := &LinkedCollectionNode{
		Element: collection,
		next:    linkedCollectionNode.Next(),
	}

	linkedCollectionNode.next = newNode

	linkedCollection.incrementLength()

	return newNode
}

func (linkedCollectionNode *LinkedCollectionNode) AddStringsPtrToNode(
	linkedCollection *LinkedCollections,
	isSkipOnNull bool,
	items *[]string,
	isMakeClone bool,
) *LinkedCollections {
	collection := NewCollectionUsingStrings(items, isMakeClone)

	return linkedCollection.AddCollectionToNode(
		isSkipOnNull,
		linkedCollectionNode,
		collection)
}

func (linkedCollectionNode *LinkedCollectionNode) AddCollectionToNode(
	linkedCollection *LinkedCollections,
	isSkipOnNull bool,
	collection *Collection,
) *LinkedCollections {
	return linkedCollection.AddCollectionToNode(
		isSkipOnNull,
		linkedCollectionNode,
		collection)
}

func (linkedCollectionNode *LinkedCollectionNode) AddNextNode(
	linkedCollection *LinkedCollections,
	nextNode *LinkedCollectionNode,
) *LinkedCollectionNode {
	nextNode.next = linkedCollectionNode.Next()
	linkedCollectionNode.next = nextNode

	linkedCollection.incrementLength()

	return nextNode
}

func (linkedCollectionNode *LinkedCollectionNode) IsChainEqual(another *LinkedCollectionNode) bool {
	if linkedCollectionNode == another {
		return true
	}

	if another == nil && linkedCollectionNode == nil {
		return true
	}

	if another == nil || linkedCollectionNode == nil {
		return false
	}

	return linkedCollectionNode.IsEqual(another) &&
		linkedCollectionNode.isNextChainEqual(another)
}

func (linkedCollectionNode *LinkedCollectionNode) IsEqual(another *LinkedCollectionNode) bool {
	if linkedCollectionNode == nil && nil == another {
		return true
	}

	if linkedCollectionNode == nil || nil == another {
		return false
	}

	if linkedCollectionNode == another {
		return true
	}

	//goland:noinspection GoNilness

	elem1 := linkedCollectionNode.Element
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

	isElementSame := elem1.IsEqualsPtr(elem2)

	return isElementSame &&
		linkedCollectionNode.isNextEqual(another)
}

func (linkedCollectionNode *LinkedCollectionNode) isNextEqual(
	another *LinkedCollectionNode,
) bool {
	next1 := linkedCollectionNode.Next()
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
		IsEqualsPtr(
			next2.Element)
}

func (linkedCollectionNode *LinkedCollectionNode) isNextChainEqual(
	another *LinkedCollectionNode,
) bool {
	next1 := linkedCollectionNode.Next()
	next2 := another.Next()

	if next1 == nil && nil == next2 {
		return true
	}

	if next1 == nil || nil == next2 {
		return false
	}

	return next1.IsChainEqual(next2)
}

func (linkedCollectionNode *LinkedCollectionNode) IsEqualValue(collection *Collection) bool {
	elem1 := linkedCollectionNode.Element

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

	return elem1.IsEqualsPtr(collection)
}

func (linkedCollectionNode *LinkedCollectionNode) EndOfChain() (
	endOfChain *LinkedCollectionNode,
	length int,
) {
	node := linkedCollectionNode
	length++

	for node.HasNext() {
		node = node.Next()
		length++
	}

	return node, length
}

func (linkedCollectionNode *LinkedCollectionNode) LoopEndOfChain(
	processor LinkedCollectionSimpleProcessor,
) (endOfLoop *LinkedCollectionNode, length int) {
	node := linkedCollectionNode
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

func (linkedCollectionNode *LinkedCollectionNode) CreateLinkedList() *LinkedCollections {
	return NewLinkedCollections().
		AppendChainOfNodes(linkedCollectionNode)
}

func (linkedCollectionNode *LinkedCollectionNode) Clone() *LinkedCollectionNode {
	return &LinkedCollectionNode{
		Element: linkedCollectionNode.Element,
		next:    nil,
	}
}

func (linkedCollectionNode *LinkedCollectionNode) String() string {
	return linkedCollectionNode.Element.String()
}

func (linkedCollectionNode *LinkedCollectionNode) ListPtr() *[]string {
	list := make([]string, 0, constants.ArbitraryCapacity100)

	node := linkedCollectionNode
	list = append(list, node.Element.List()...)

	for node.HasNext() {
		node = node.Next()

		list = append(list, node.Element.List()...)
	}

	return &list
}

func (linkedCollectionNode *LinkedCollectionNode) Join(separator string) *string {
	list := linkedCollectionNode.ListPtr()
	toString := strings.Join(*list, separator)

	return &toString
}

func (linkedCollectionNode *LinkedCollectionNode) StringListPtr(header string) *string {
	finalString := header +
		*linkedCollectionNode.Join(commonJoiner)

	return &finalString
}

func (linkedCollectionNode *LinkedCollectionNode) Print(header string) {
	finalString := linkedCollectionNode.StringListPtr(header)
	fmt.Println(finalString)
}
