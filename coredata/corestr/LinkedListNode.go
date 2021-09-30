package corestr

import (
	"fmt"
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

type LinkedListNode struct {
	Element string
	next    *LinkedListNode
}

func (linkedListNode *LinkedListNode) HasNext() bool {
	return linkedListNode.next != nil
}

func (linkedListNode *LinkedListNode) Next() *LinkedListNode {
	return linkedListNode.next
}

func (linkedListNode *LinkedListNode) EndOfChain() (
	endOfChain *LinkedListNode,
	length int,
) {
	node := linkedListNode
	length++

	for node.HasNext() {
		node = node.Next()
		length++
	}

	return node, length
}

func (linkedListNode *LinkedListNode) LoopEndOfChain(
	processor LinkedListSimpleProcessor,
) (endOfLoop *LinkedListNode, length int) {
	node := linkedListNode
	arg := &LinkedListProcessorParameter{
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

		arg2 := &LinkedListProcessorParameter{
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

func (linkedListNode *LinkedListNode) Clone() *LinkedListNode {
	return &LinkedListNode{
		Element: linkedListNode.Element,
		next:    nil,
	}
}

func (linkedListNode *LinkedListNode) AddNext(
	linkedListForIncrement *LinkedList,
	item string,
) *LinkedListNode {
	newNode := &LinkedListNode{
		Element: item,
		next:    linkedListNode.Next(),
	}

	linkedListNode.next = newNode
	linkedListForIncrement.incrementLength()

	return newNode
}

func (linkedListNode *LinkedListNode) AddStringsPtrToNode(
	linkedListForIncrement *LinkedList,
	isSkipOnNull bool,
	items *[]string,
) *LinkedList {
	return linkedListForIncrement.AddStringsPtrToNode(
		isSkipOnNull,
		linkedListNode,
		items)
}

func (linkedListNode *LinkedListNode) AddCollectionToNode(
	linkedListForIncrement *LinkedList,
	isSkipOnNull bool,
	collection *Collection,
) *LinkedList {
	return linkedListForIncrement.AddStringsPtrToNode(
		isSkipOnNull,
		linkedListNode,
		collection.ListPtr())
}

func (linkedListNode *LinkedListNode) AddNextNode(
	linkedListForIncrement *LinkedList,
	nextNode *LinkedListNode,
) *LinkedListNode {
	nextNode.next = linkedListNode.Next()
	linkedListNode.next = nextNode
	linkedListForIncrement.incrementLength()

	return nextNode
}

func (linkedListNode *LinkedListNode) IsEqual(another *LinkedListNode) bool {
	if linkedListNode == nil && nil == another {
		return true
	}

	if linkedListNode == nil || nil == another {
		return false
	}

	if linkedListNode == another {
		return true
	}

	//goland:noinspection GoNilness
	if linkedListNode.Element == another.Element {
		return linkedListNode.isNextEqual(
			another,
			true)
	}

	return false
}

func (linkedListNode *LinkedListNode) IsChainEqual(
	another *LinkedListNode,
	isCaseSensitive bool,
) bool {
	if linkedListNode == nil && nil == another {
		return true
	}

	if linkedListNode == nil || nil == another {
		return false
	}

	if linkedListNode == another {
		return true
	}

	elem1 := linkedListNode.Element
	elem2 := another.Element

	//goland:noinspection GoNilness
	isElementSame := (isCaseSensitive && elem1 == elem2) ||
		(!isCaseSensitive && strings.EqualFold(elem1, elem2))

	return isElementSame &&
		linkedListNode.isNextChainEqual(
			another, isCaseSensitive)
}

func (linkedListNode *LinkedListNode) IsEqualSensitive(
	another *LinkedListNode,
	isCaseSensitive bool,
) bool {
	if linkedListNode == another {
		return true
	}

	if another == nil && linkedListNode == nil {
		return true
	}

	if another == nil || linkedListNode == nil {
		return false
	}

	isSame := linkedListNode.IsEqualValueSensitive(another.Element, isCaseSensitive)

	return isSame &&
		linkedListNode.isNextEqual(another, isCaseSensitive)
}

func (linkedListNode *LinkedListNode) isNextEqual(
	another *LinkedListNode,
	isCaseSensitive bool,
) bool {
	next1 := linkedListNode.Next()
	next2 := another.Next()

	if next1 == nil && nil == next2 {
		return true
	}

	if next1 == nil || nil == next2 {
		return false
	}

	if isCaseSensitive {
		return next1.Element == next2.Element
	}

	return strings.EqualFold(next1.Element, next2.Element)
}

func (linkedListNode *LinkedListNode) isNextChainEqual(
	another *LinkedListNode,
	isCaseSensitive bool,
) bool {
	next1 := linkedListNode.Next()
	next2 := another.Next()

	if next1 == nil && nil == next2 {
		return true
	}

	if next1 == nil || nil == next2 {
		return false
	}

	return next1.
		IsChainEqual(next2, isCaseSensitive)
}

func (linkedListNode *LinkedListNode) CreateLinkedList() *LinkedList {
	return NewLinkedList().
		AppendChainOfNodes(linkedListNode)
}

func (linkedListNode *LinkedListNode) IsEqualValue(value string) bool {
	return linkedListNode.Element == value
}

func (linkedListNode *LinkedListNode) IsEqualValueSensitive(value string, isCaseSensitive bool) bool {
	if isCaseSensitive {
		return value == linkedListNode.Element
	}

	return strings.EqualFold(linkedListNode.Element, value)
}

func (linkedListNode *LinkedListNode) String() string {
	return linkedListNode.Element
}

func (linkedListNode *LinkedListNode) ListPtr() *[]string {
	list := make([]string, 0, constants.ArbitraryCapacity100)

	node := linkedListNode
	list = append(list, node.Element)

	for node.HasNext() {
		node = node.Next()

		list = append(list, node.Element)
	}

	return &list
}

func (linkedListNode *LinkedListNode) Join(separator string) *string {
	list := linkedListNode.ListPtr()
	toString := strings.Join(*list, separator)

	return &toString
}

func (linkedListNode *LinkedListNode) StringListPtr(header string) *string {
	finalString := header +
		*linkedListNode.Join(commonJoiner)

	return &finalString
}

func (linkedListNode *LinkedListNode) Print(header string) {
	finalString := linkedListNode.StringListPtr(header)
	fmt.Println(finalString)
}
