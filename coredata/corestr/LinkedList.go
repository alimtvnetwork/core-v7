package corestr

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coreindexes"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/internal/strutilinternal"
)

type LinkedList struct {
	head, tail *LinkedListNode
	length     int
	sync.Mutex
}

func (it *LinkedList) Tail() *LinkedListNode {
	return it.tail
}

func (it *LinkedList) Head() *LinkedListNode {
	return it.head
}

func (it *LinkedList) Length() int {
	return it.length
}

func (it *LinkedList) incrementLength() int {
	it.length++

	return it.length
}

func (it *LinkedList) incrementLengthUsingNumber(number int) int {
	it.length += number

	return it.length
}

func (it *LinkedList) setLengthToZero() int {
	it.length = 0

	return it.length
}

func (it *LinkedList) setLength(number int) int {
	it.length = number

	return it.length
}

func (it *LinkedList) decrementLength() int {
	it.length--

	return it.length
}

func (it *LinkedList) LengthLock() int {
	it.Lock()
	defer it.Unlock()

	return it.length
}

//goland:noinspection GoVetCopyLock
func (it *LinkedList) IsEquals(
	anotherLinkedList LinkedList,
) bool {
	return it.IsEqualsWithSensitivePtr(
		&anotherLinkedList,
		true)
}

func (it *LinkedList) IsEqualsPtr(
	anotherLinkedList *LinkedList,
) bool {
	return it.IsEqualsWithSensitivePtr(
		anotherLinkedList,
		true)
}

func (it *LinkedList) IsEqualsWithSensitivePtr(
	anotherLinkedList *LinkedList,
	isCaseSensitive bool,
) bool {
	if anotherLinkedList == nil && it == nil {
		return true
	}

	if anotherLinkedList == nil || it == nil {
		return false
	}

	if it == anotherLinkedList {
		return true
	}

	if it.IsEmpty() && anotherLinkedList.IsEmpty() {
		return true
	}

	if it.IsEmpty() || anotherLinkedList.IsEmpty() {
		return false
	}

	if it.Length() != anotherLinkedList.Length() {
		return false
	}

	leftNode := it.head
	rightNode := anotherLinkedList.head

	if leftNode == nil && rightNode == nil {
		return true
	}

	if leftNode == nil || rightNode == nil {
		return false
	}

	return leftNode.IsChainEqual(rightNode, isCaseSensitive)
}

func (it *LinkedList) IsEmptyLock() bool {
	it.Lock()
	defer it.Unlock()

	return it.head == nil || it.length == 0
}

func (it *LinkedList) IsEmpty() bool {
	return it.head == nil ||
		it.length == 0
}

func (it *LinkedList) HasItems() bool {
	return it.head != nil &&
		it.length > 0
}

func (it *LinkedList) Add(item string) *LinkedList {
	if it.IsEmpty() {
		it.head = &LinkedListNode{
			Element: item,
			next:    nil,
		}

		it.tail = it.head
		it.incrementLength()

		return it
	}

	it.tail.next = &LinkedListNode{
		Element: item,
		next:    nil,
	}

	it.tail = it.tail.next
	it.incrementLength()

	return it
}

func (it *LinkedList) AddItemsMap(itemsMap map[string]bool) *LinkedList {
	if len(itemsMap) == 0 {
		return it
	}

	for key, isAdd := range itemsMap {
		if !isAdd {
			continue
		}

		it.Add(key)
	}

	return it
}

func (it *LinkedList) AddLock(item string) *LinkedList {
	it.Lock()
	defer it.Unlock()

	return it.Add(item)
}

// InsertAt BigO(n) expensive operation.
func (it *LinkedList) InsertAt(index int, item string) *LinkedList {
	if index < 1 {
		return it.AddFront(item)
	}

	node := it.IndexAt(index - 1)
	it.AddAfterNode(node, item)

	return it
}

func (it *LinkedList) AddBackNode(node *LinkedListNode) *LinkedList {
	return it.AppendNode(node)
}

func (it *LinkedList) AppendNode(node *LinkedListNode) *LinkedList {
	if it.IsEmpty() {
		it.head = node
		it.tail = it.head
		it.incrementLength()

		return it
	}

	it.tail.next = node
	it.tail = it.tail.next
	it.incrementLength()

	return it
}

func (it *LinkedList) AppendChainOfNodes(nodeHead *LinkedListNode) *LinkedList {
	endOfChain, length := nodeHead.EndOfChain()

	if it.IsEmpty() {
		it.head = nodeHead
	} else {
		it.tail.next = nodeHead
	}

	it.tail = endOfChain
	it.incrementLengthUsingNumber(length)

	return it
}

func (it *LinkedList) PushBack(item string) *LinkedList {
	return it.Add(item)
}

func (it *LinkedList) AddNonEmpty(item string) *LinkedList {
	if item == "" {
		return it
	}

	return it.Add(item)
}

func (it *LinkedList) AddNonEmptyWhitespace(item string) *LinkedList {
	if strutilinternal.IsEmptyOrWhitespace(item) {
		return it
	}

	return it.Add(item)
}

func (it *LinkedList) AddIf(isAdd bool, item string) *LinkedList {
	if !isAdd {
		return it
	}

	return it.Add(item)
}

func (it *LinkedList) AddIfMany(
	isAdd bool,
	addingStrings ...string,
) *LinkedList {
	if !isAdd {
		return it
	}

	return it.AddStringsPtr(&addingStrings)
}

func (it *LinkedList) AddFunc(f func() string) *LinkedList {
	return it.Add(f())
}

func (it *LinkedList) AddFuncErr(
	funcReturnsStringError func() (result string, err error),
	errHandler func(errInput error),
) *LinkedList {
	r, err := funcReturnsStringError()

	if err != nil {
		errHandler(err)

		return it
	}

	return it.Add(r)
}

func (it *LinkedList) Push(item string) *LinkedList {
	return it.Add(item)
}

func (it *LinkedList) PushFront(item string) *LinkedList {
	return it.AddFront(item)
}

func (it *LinkedList) AddFront(item string) *LinkedList {
	if it.IsEmpty() {
		return it.Add(item)
	}

	node := &LinkedListNode{
		Element: item,
		next:    it.head,
	}

	it.head = node
	it.incrementLength()

	return it
}

func (it *LinkedList) AttachWithNode(currentNode, addingNode *LinkedListNode) error {
	if currentNode == nil {
		return errcore.
			CannotBeNilType.
			Error("CurrentNode cannot be nil.", nil)
	}

	if currentNode.next != nil {
		return errcore.
			ShouldBeNilType.
			Error("CurrentNode.next", nil)
	}

	addingNode.next = currentNode.next
	currentNode.next = addingNode

	it.incrementLength()

	return nil
}

// AddCollectionToNode iSkipOnNil
func (it *LinkedList) AddCollectionToNode(
	isSkipOnNull bool,
	node *LinkedListNode,
	collection *Collection,
) *LinkedList {
	return it.AddStringsPtrToNode(
		isSkipOnNull,
		node,
		collection.ListPtr())
}

func (it *LinkedList) Loop(
	simpleProcessor LinkedListSimpleProcessor,
) *LinkedList {
	length := it.Length()
	if length == 0 {
		return it
	}

	node := it.head
	arg := &LinkedListProcessorParameter{
		Index:         0,
		CurrentNode:   node,
		PrevNode:      nil,
		IsFirstIndex:  true,
		IsEndingIndex: false,
	}

	isBreak := simpleProcessor(arg)

	if isBreak {
		return it
	}

	lenMinusOne := length - 1
	index := 1
	isEndingIndex := false

	for node.HasNext() {
		prev := node
		node = node.Next()
		isEndingIndex = lenMinusOne == index

		arg2 := &LinkedListProcessorParameter{
			Index:         index,
			CurrentNode:   node,
			PrevNode:      prev,
			IsFirstIndex:  false,
			IsEndingIndex: isEndingIndex,
		}

		isBreak2 := simpleProcessor(arg2)

		if isBreak2 {
			return it
		}

		index++
	}

	return it
}

func (it *LinkedList) Filter(
	filter LinkedListFilter,
) *[]*LinkedListNode {
	length := it.Length()
	list := make([]*LinkedListNode, 0, length)

	if length == 0 {
		return &list
	}

	node := it.head
	arg := &LinkedListFilterParameter{
		Node:  node,
		Index: 0,
	}
	result := filter(arg)

	if result.IsKeep {
		list = append(list, result.Value)
	}

	if result.IsBreak {
		return &list
	}

	index := 1

	for node.HasNext() {
		node = node.Next()

		arg2 := &LinkedListFilterParameter{
			Node:  node,
			Index: index,
		}
		result2 := filter(arg2)

		if result2.IsKeep {
			list = append(list, result2.Value)
		}

		if result2.IsBreak {
			return &list
		}

		index++
	}

	return &list
}

func (it *LinkedList) RemoveNodeByElementValue(
	element string,
	isCaseSensitive bool,
	isIgnorePanic bool,
) *LinkedList {
	if !isIgnorePanic && it.IsEmpty() {
		errcore.
			CannotRemoveIndexesFromEmptyCollectionType.
			HandleUsingPanic("element cannot be removed from Empty linkedlist.", element)
	}

	var processor LinkedListSimpleProcessor = func(
		arg *LinkedListProcessorParameter,
	) (isBreak bool) {
		isSameNode :=
			(isCaseSensitive && arg.CurrentNode.Element == element) ||
				(!isCaseSensitive && strings.EqualFold(element, arg.CurrentNode.Element))

		if isSameNode && arg.IsFirstIndex {
			it.head = arg.CurrentNode.next
			it.decrementLength()

			return false
		}

		if isSameNode {
			arg.PrevNode.next = arg.CurrentNode.next
			it.decrementLength()
		}

		return false
	}

	return it.Loop(processor)
}

func (it *LinkedList) RemoveNodeByIndex(
	removingIndex int,
) *LinkedList {
	if removingIndex < 0 {
		errcore.
			CannotBeNegativeIndexType.
			HandleUsingPanic(
				"removeIndex was less than 0.",
				removingIndex)
	}

	var singleProcessor LinkedListSimpleProcessor = func(
		arg *LinkedListProcessorParameter,
	) (isBreak bool) {
		hasIndex := removingIndex == arg.Index

		if !hasIndex {
			return false
		}

		isBreak = hasIndex
		it.decrementLength()

		if arg.IsFirstIndex {
			it.head =
				arg.CurrentNode.next
			arg.CurrentNode = nil
			return isBreak
		}

		if arg.IsEndingIndex {
			arg.PrevNode.next = nil
			arg.CurrentNode = nil

			return isBreak
		}

		arg.PrevNode.next = arg.CurrentNode.next
		arg.CurrentNode = nil

		return isBreak
	}

	return it.Loop(singleProcessor)
}

func (it *LinkedList) RemoveNodeByIndexes(
	isIgnorePanic bool,
	removingIndexes ...int,
) *LinkedList {
	length := len(removingIndexes)

	if length == 0 {
		return it
	}

	if !isIgnorePanic && it.IsEmpty() && length > 0 {
		errcore.
			CannotRemoveIndexesFromEmptyCollectionType.
			HandleUsingPanic("removingIndexes cannot be removed from Empty linkedlist.", removingIndexes)
	}

	removingIndexesCopy := removingIndexes

	nonChainedNodes := it.Filter(
		func(
			arg *LinkedListFilterParameter,
		) *LinkedListFilterResult {
			hasIndex := coreindexes.HasIndexPlusRemoveIndex(removingIndexesCopy, arg.Index)
			if hasIndex {
				// remove
				return &LinkedListFilterResult{
					Value:   arg.Node,
					IsKeep:  false,
					IsBreak: false,
				}
			}

			// not remove
			return &LinkedListFilterResult{
				Value:   arg.Node,
				IsKeep:  true,
				IsBreak: false,
			}
		})

	nonChainedCollection := &NonChainedLinkedListNodes{
		items:             nonChainedNodes,
		isChainingApplied: false,
	}

	if nonChainedCollection.IsEmpty() {
		return it
	}

	it.setLength(nonChainedCollection.Length())
	it.head = nonChainedCollection.ApplyChaining().First()

	return it
}

func (it *LinkedList) GetCompareSummary(
	right *LinkedList,
	leftName, rightName string,
) string {
	lLen := it.Length()
	rLen := right.Length()

	leftStr := fmt.Sprintf(
		linkedListCollectionCompareHeaderLeft,
		leftName,
		lLen,
		it)

	rightStr := fmt.Sprintf(
		linkedListCollectionCompareHeaderRight,
		rightName,
		rLen,
		right,
		it.IsEqualsPtr(right),
		lLen,
		rLen)

	return leftStr + rightStr
}

// RemoveNode skip if removingNode is nil
func (it *LinkedList) RemoveNode(
	removingNode *LinkedListNode,
) *LinkedList {
	if removingNode == nil {
		return it
	}

	if it.IsEmpty() {
		errcore.
			CannotRemoveIndexesFromEmptyCollectionType.
			HandleUsingPanic("removingNode cannot be removed from Empty linkedlist.", removingNode.String())
	}

	var processor LinkedListSimpleProcessor = func(
		arg *LinkedListProcessorParameter,
	) (isBreak bool) {
		isSameNode := arg.CurrentNode == removingNode
		if isSameNode && arg.IsFirstIndex {
			it.head = arg.CurrentNode.next
			it.decrementLength()

			return true
		}

		if isSameNode {
			arg.PrevNode.next = arg.CurrentNode.next
			it.decrementLength()

			return true
		}

		return false
	}

	return it.Loop(processor)
}

// AddStringsPtrToNode iSkipOnNil
func (it *LinkedList) AddStringsPtrToNode(
	isSkipOnNull bool,
	node *LinkedListNode,
	items *[]string,
) *LinkedList {
	if items == nil || node == nil && isSkipOnNull {
		return it
	}

	if node == nil {
		errcore.
			CannotBeNilType.
			HandleUsingPanic(
				"node cannot be nil.",
				nil)
	}

	length := len(*items)

	if length == 0 {
		return it
	}

	if length == 1 {
		it.AddAfterNode(node, (*items)[0])

		return it
	}

	finalHead := &LinkedListNode{
		Element: (*items)[0],
		next:    nil,
	}

	nextNode := finalHead

	for _, item := range (*items)[1:] {
		nextNode = nextNode.AddNext(it, item)
	}

	//goland:noinspection GoNilness
	nextNode.next = node.next
	//goland:noinspection GoNilness
	node.next = finalHead
	it.incrementLength()

	return it
}

func (it *LinkedList) AddAfterNode(
	node *LinkedListNode,
	item string,
) *LinkedListNode {
	newNode := &LinkedListNode{
		Element: item,
		next:    node.next,
	}

	node.next = newNode
	it.incrementLength()

	return newNode
}

// AddStringsPtr add to back
func (it *LinkedList) AddStringsPtr(items *[]string) *LinkedList {
	if items == nil {
		return it
	}

	for _, item := range *items {
		it.Add(item)
	}

	return it
}

// AddStrings items add to back
func (it *LinkedList) AddStrings(items ...string) *LinkedList {
	if len(items) == 0 {
		return it
	}

	for _, item := range items {
		it.Add(item)
	}

	return it
}

// AddStringsPtrLock add to back
func (it *LinkedList) AddStringsPtrLock(items *[]string) *LinkedList {
	it.Lock()
	defer it.Unlock()

	return it.AddStringsPtr(items)
}

// IndexAt Expensive operation BigO(n)
func (it *LinkedList) IndexAt(index int) *LinkedListNode {
	length := it.Length()
	if index < 0 {
		return nil
	}

	if length == 0 || length-1 < index {
		errcore.OutOfRangeType.HandleUsingPanic(
			"Given index is out of range. Whereas length:",
			length)
	}

	if index == 0 {
		return it.head
	}

	node := it.head
	i := 1
	for node.HasNext() {
		node = node.Next()

		if i == index {
			return node
		}

		i++
	}

	return nil
}

// SafePointerIndexAt Expensive operation BigO(n)
func (it *LinkedList) SafePointerIndexAt(index int) *string {
	node := it.SafeIndexAt(index)

	if node == nil {
		return nil
	}

	return &node.Element
}

// SafePointerIndexAtUsingDefault Expensive operation BigO(n)
func (it *LinkedList) SafePointerIndexAtUsingDefault(
	index int,
	defaultString string,
) string {
	node := it.SafeIndexAt(index)

	if node == nil {
		return defaultString
	}

	return node.Element
}

// SafeIndexAt Expensive operation BigO(n)
func (it *LinkedList) SafeIndexAt(index int) *LinkedListNode {
	length := it.Length()
	isExitCondition := index < 0 || length == 0 || length-1 < index
	if isExitCondition {
		return nil
	}

	if index == 0 {
		return it.head
	}

	node := it.head
	i := 1
	for node.HasNext() {
		node = node.Next()

		if i == index {
			return node
		}

		i++
	}

	return nil
}

// SafeIndexAtLock Expensive operation BigO(n)
func (it *LinkedList) SafeIndexAtLock(index int) *LinkedListNode {
	it.Lock()
	defer it.Unlock()

	return it.SafeIndexAt(index)
}

// SafePointerIndexAtUsingDefaultLock Expensive operation BigO(n)
func (it *LinkedList) SafePointerIndexAtUsingDefaultLock(
	index int,
	defaultString string,
) string {
	it.Lock()
	defer it.Unlock()

	return it.SafePointerIndexAtUsingDefault(index, defaultString)
}

func (it *LinkedList) GetNextNodes(count int) *[]*LinkedListNode {
	counter := 0

	return it.Filter(
		func(
			arg *LinkedListFilterParameter,
		) *LinkedListFilterResult {
			isBreak := counter >= count-1
			return &LinkedListFilterResult{
				Value:   arg.Node,
				IsKeep:  true,
				IsBreak: isBreak,
			}
		})
}

func (it *LinkedList) GetAllLinkedNodes() *[]*LinkedListNode {
	return it.Filter(
		func(
			arg *LinkedListFilterParameter,
		) *LinkedListFilterResult {
			return &LinkedListFilterResult{
				Value:   arg.Node,
				IsKeep:  true,
				IsBreak: false,
			}
		})
}

// AddPointerStringsPtr skip on nil, add to back
func (it *LinkedList) AddPointerStringsPtr(items *[]*string) *LinkedList {
	if items == nil {
		return it
	}

	for _, item := range *items {
		if item == nil {
			continue
		}

		it.Add(*item)
	}

	return it
}

// AddCollection skip on nil
func (it *LinkedList) AddCollection(collection *Collection) *LinkedList {
	if collection == nil {
		return it
	}

	for _, item := range collection.items {
		it.Add(item)
	}

	return it
}

func (it *LinkedList) ToCollection(addCapacity int) *Collection {
	newLength := it.Length() + addCapacity
	collection := New.Collection.Cap(newLength)

	if it.IsEmpty() {
		return collection
	}

	node := it.head
	collection.Add(node.Element)

	for node.HasNext() {
		node = node.Next()
		collection.Add(node.Element)
	}

	return collection
}

// List must return slice.
func (it *LinkedList) List() []string {
	list := make(
		[]string,
		0,
		it.Length())

	if it.IsEmpty() {
		return list
	}

	node := it.head
	list = append(list, node.Element)

	for node.HasNext() {
		node = node.Next()
		list = append(list, node.Element)
	}

	return list
}

// ListPtr must return slice.
func (it *LinkedList) ListPtr() *[]string {
	list := it.List()

	return &list
}

// ListPtrLock must return slice.
func (it *LinkedList) ListPtrLock() *[]string {
	it.Lock()
	defer it.Unlock()

	return it.ListPtr()
}

func (it *LinkedList) String() string {
	if it.IsEmpty() {
		return commonJoiner + NoElements
	}

	return commonJoiner +
		strings.Join(
			it.List(),
			commonJoiner)
}

func (it *LinkedList) StringLock() string {
	if it.IsEmptyLock() {
		return commonJoiner + NoElements
	}

	it.Lock()
	defer it.Unlock()

	return commonJoiner +
		strings.Join(
			*it.ListPtr(),
			commonJoiner)
}

func (it *LinkedList) Join(
	separator string,
) string {
	return strings.Join(*it.ListPtr(), separator)
}

func (it *LinkedList) JoinLock(
	separator string,
) string {
	it.Lock()
	defer it.Unlock()

	return strings.Join(*it.ListPtr(), separator)
}

func (it *LinkedList) Joins(
	separator string,
	items ...string,
) string {
	if items == nil || it.Length() == 0 {
		return strings.Join(items, separator)
	}

	collection := it.ToCollection(len(items) +
		constants.ArbitraryCapacity2)
	collection.AddStringsPtr(&items)

	return collection.Join(separator)
}

func (it *LinkedList) JsonModel() []string {
	return it.ToCollection(0).JsonModel()
}

func (it *LinkedList) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it *LinkedList) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.JsonModel())
}

func (it *LinkedList) UnmarshalJSON(data []byte) error {
	var dataModelStrings []string
	err := json.Unmarshal(data, &dataModelStrings)

	if err == nil {
		it.Clear()
		it.AddStrings(dataModelStrings...)
	}

	return err
}

func (it *LinkedList) RemoveAll() *LinkedList {
	return it.Clear()
}

func (it *LinkedList) Clear() *LinkedList {
	if it.IsEmpty() {
		return it
	}

	it.head = nil
	it.tail = nil
	it.setLengthToZero()

	return it
}

func (it LinkedList) Json() corejson.Result {
	return corejson.New(it)
}

func (it LinkedList) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it *LinkedList) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*LinkedList, error) {
	err := jsonResult.Unmarshal(it)

	if err != nil {
		return New.LinkedList.Create(), err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
func (it *LinkedList) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *LinkedList {
	newUsingJson, err :=
		it.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return newUsingJson
}

// JsonParseSelfInject Panic if error
func (it *LinkedList) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it *LinkedList) AsJsonMarshaller() corejson.JsonMarshaller {
	return it
}
