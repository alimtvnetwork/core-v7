package corestr

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coreindexes"
	"gitlab.com/evatix-go/core/defaulterr"
	"gitlab.com/evatix-go/core/internal/strutilinternal"
	"gitlab.com/evatix-go/core/msgtype"
)

type LinkedList struct {
	head, tail *LinkedListNode
	length     int
	sync.Mutex
}

func (linkedList *LinkedList) Tail() *LinkedListNode {
	return linkedList.tail
}

func (linkedList *LinkedList) Head() *LinkedListNode {
	return linkedList.head
}

func (linkedList *LinkedList) Length() int {
	return linkedList.length
}

func (linkedList *LinkedList) incrementLength() int {
	linkedList.length++

	return linkedList.length
}

func (linkedList *LinkedList) incrementLengthUsingNumber(number int) int {
	linkedList.length += number

	return linkedList.length
}

func (linkedList *LinkedList) setLengthToZero() int {
	linkedList.length = 0

	return linkedList.length
}

func (linkedList *LinkedList) setLength(number int) int {
	linkedList.length = number

	return linkedList.length
}

func (linkedList *LinkedList) decrementLength() int {
	linkedList.length--

	return linkedList.length
}

func (linkedList *LinkedList) LengthLock() int {
	linkedList.Lock()
	defer linkedList.Unlock()

	return linkedList.length
}

//goland:noinspection GoVetCopyLock
func (linkedList *LinkedList) IsEquals(
	anotherLinkedList LinkedList,
) bool {
	return linkedList.IsEqualsWithSensitivePtr(
		&anotherLinkedList,
		true)
}

func (linkedList *LinkedList) IsEqualsPtr(
	anotherLinkedList *LinkedList,
) bool {
	return linkedList.IsEqualsWithSensitivePtr(
		anotherLinkedList,
		true)
}

func (linkedList *LinkedList) IsEqualsWithSensitivePtr(
	anotherLinkedList *LinkedList,
	isCaseSensitive bool,
) bool {
	if anotherLinkedList == nil && linkedList == nil {
		return true
	}

	if anotherLinkedList == nil || linkedList == nil {
		return false
	}

	if linkedList == anotherLinkedList {
		return true
	}

	if linkedList.IsEmpty() && anotherLinkedList.IsEmpty() {
		return true
	}

	if linkedList.IsEmpty() || anotherLinkedList.IsEmpty() {
		return false
	}

	if linkedList.Length() != anotherLinkedList.Length() {
		return false
	}

	leftNode := linkedList.head
	rightNode := anotherLinkedList.head

	if leftNode == nil && rightNode == nil {
		return true
	}

	if leftNode == nil || rightNode == nil {
		return false
	}

	return leftNode.IsChainEqual(rightNode, isCaseSensitive)
}

func (linkedList *LinkedList) IsEmptyLock() bool {
	linkedList.Lock()
	defer linkedList.Unlock()

	return linkedList.head == nil || linkedList.length == 0
}

func (linkedList *LinkedList) IsEmpty() bool {
	return linkedList.head == nil ||
		linkedList.length == 0
}

func (linkedList *LinkedList) HasItems() bool {
	return linkedList.head != nil &&
		linkedList.length > 0
}

func (linkedList *LinkedList) Add(item string) *LinkedList {
	if linkedList.IsEmpty() {
		linkedList.head = &LinkedListNode{
			Element: item,
			next:    nil,
		}

		linkedList.tail = linkedList.head
		linkedList.incrementLength()

		return linkedList
	}

	linkedList.tail.next = &LinkedListNode{
		Element: item,
		next:    nil,
	}

	linkedList.tail = linkedList.tail.next
	linkedList.incrementLength()

	return linkedList
}

func (linkedList *LinkedList) AddLock(item string) *LinkedList {
	linkedList.Lock()
	defer linkedList.Unlock()

	return linkedList.Add(item)
}

// InsertAt BigO(n) expensive operation.
func (linkedList *LinkedList) InsertAt(index int, item string) *LinkedList {
	if index < 1 {
		return linkedList.AddFront(item)
	}

	node := linkedList.IndexAt(index - 1)
	linkedList.AddAfterNode(node, item)

	return linkedList
}

func (linkedList *LinkedList) AddBackNode(node *LinkedListNode) *LinkedList {
	return linkedList.AppendNode(node)
}

func (linkedList *LinkedList) AppendNode(node *LinkedListNode) *LinkedList {
	if linkedList.IsEmpty() {
		linkedList.head = node
		linkedList.tail = linkedList.head
		linkedList.incrementLength()

		return linkedList
	}

	linkedList.tail.next = node
	linkedList.tail = linkedList.tail.next
	linkedList.incrementLength()

	return linkedList
}

func (linkedList *LinkedList) AppendChainOfNodes(nodeHead *LinkedListNode) *LinkedList {
	endOfChain, length := nodeHead.EndOfChain()

	if linkedList.IsEmpty() {
		linkedList.head = nodeHead
	} else {
		linkedList.tail.next = nodeHead
	}

	linkedList.tail = endOfChain
	linkedList.incrementLengthUsingNumber(length)

	return linkedList
}

func (linkedList *LinkedList) PushBack(item string) *LinkedList {
	return linkedList.Add(item)
}

func (linkedList *LinkedList) AddNonEmpty(item string) *LinkedList {
	if item == "" {
		return linkedList
	}

	return linkedList.Add(item)
}

func (linkedList *LinkedList) AddNonEmptyWhitespace(item string) *LinkedList {
	if strutilinternal.IsEmptyOrWhitespace(item) {
		return linkedList
	}

	return linkedList.Add(item)
}

func (linkedList *LinkedList) AddIf(isAdd bool, item string) *LinkedList {
	if !isAdd {
		return linkedList
	}

	return linkedList.Add(item)
}

func (linkedList *LinkedList) AddIfMany(
	isAdd bool,
	addingStrings ...string,
) *LinkedList {
	if !isAdd {
		return linkedList
	}

	return linkedList.AddStringsPtr(&addingStrings)
}

func (linkedList *LinkedList) AddFunc(f func() string) *LinkedList {
	return linkedList.Add(f())
}

func (linkedList *LinkedList) AddFuncErr(
	funcReturnsStringError func() (result string, err error),
	errHandler func(errInput error),
) *LinkedList {
	r, err := funcReturnsStringError()

	if err != nil {
		errHandler(err)

		return linkedList
	}

	return linkedList.Add(r)
}

func (linkedList *LinkedList) Push(item string) *LinkedList {
	return linkedList.Add(item)
}

func (linkedList *LinkedList) PushFront(item string) *LinkedList {
	return linkedList.AddFront(item)
}

func (linkedList *LinkedList) AddFront(item string) *LinkedList {
	if linkedList.IsEmpty() {
		return linkedList.Add(item)
	}

	node := &LinkedListNode{
		Element: item,
		next:    linkedList.head,
	}

	linkedList.head = node
	linkedList.incrementLength()

	return linkedList
}

func (linkedList *LinkedList) AttachWithNode(currentNode, addingNode *LinkedListNode) error {
	if currentNode == nil {
		return msgtype.
			CannotBeNilMessage.
			Error("CurrentNode cannot be nil.", nil)
	}

	if currentNode.next != nil {
		return msgtype.
			ShouldBeNilMessage.
			Error("CurrentNode.next", nil)
	}

	addingNode.next = currentNode.next
	currentNode.next = addingNode

	linkedList.incrementLength()

	return nil
}

// AddCollectionToNode iSkipOnNil
func (linkedList *LinkedList) AddCollectionToNode(
	isSkipOnNull bool,
	node *LinkedListNode,
	collection *Collection,
) *LinkedList {
	return linkedList.AddStringsPtrToNode(
		isSkipOnNull,
		node,
		collection.items)
}

func (linkedList *LinkedList) Loop(
	simpleProcessor LinkedListSimpleProcessor,
) *LinkedList {
	length := linkedList.Length()
	if length == 0 {
		return linkedList
	}

	node := linkedList.head
	arg := &LinkedListProcessorParameter{
		Index:         0,
		CurrentNode:   node,
		PrevNode:      nil,
		IsFirstIndex:  true,
		IsEndingIndex: false,
	}

	isBreak := simpleProcessor(arg)

	if isBreak {
		return linkedList
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
			return linkedList
		}

		index++
	}

	return linkedList
}

func (linkedList *LinkedList) Filter(
	filter LinkedListFilter,
) *[]*LinkedListNode {
	length := linkedList.Length()
	list := make([]*LinkedListNode, 0, length)

	if length == 0 {
		return &list
	}

	node := linkedList.head
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

func (linkedList *LinkedList) RemoveNodeByElementValue(
	element string,
	isCaseSensitive bool,
	isIgnorePanic bool,
) *LinkedList {
	if !isIgnorePanic && linkedList.IsEmpty() {
		msgtype.
			CannotRemoveIndexesFromEmptyCollection.
			HandleUsingPanic("element cannot be removed from empty linkedlist.", element)
	}

	var processor LinkedListSimpleProcessor = func(
		arg *LinkedListProcessorParameter,
	) (isBreak bool) {
		isSameNode :=
			(isCaseSensitive && arg.CurrentNode.Element == element) ||
				(!isCaseSensitive && strings.EqualFold(element, arg.CurrentNode.Element))

		if isSameNode && arg.IsFirstIndex {
			linkedList.head = arg.CurrentNode.next
			linkedList.decrementLength()

			return false
		}

		if isSameNode {
			arg.PrevNode.next = arg.CurrentNode.next
			linkedList.decrementLength()
		}

		return false
	}

	return linkedList.Loop(processor)
}

func (linkedList *LinkedList) RemoveNodeByIndex(
	removingIndex int,
) *LinkedList {
	if removingIndex < 0 {
		msgtype.
			CannotBeNegativeIndex.
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
		linkedList.decrementLength()

		if arg.IsFirstIndex {
			linkedList.head =
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

	return linkedList.Loop(singleProcessor)
}

func (linkedList *LinkedList) RemoveNodeByIndexes(
	isIgnorePanic bool,
	removingIndexes ...int,
) *LinkedList {
	length := len(removingIndexes)

	if length == 0 {
		return linkedList
	}

	if !isIgnorePanic && linkedList.IsEmpty() && length > 0 {
		msgtype.
			CannotRemoveIndexesFromEmptyCollection.
			HandleUsingPanic("removingIndexes cannot be removed from empty linkedlist.", removingIndexes)
	}

	removingIndexesCopy := removingIndexes
	removingIndexesCopyPtr := &removingIndexesCopy

	nonChainedNodes := linkedList.Filter(
		func(
			arg *LinkedListFilterParameter,
		) *LinkedListFilterResult {
			hasIndex := coreindexes.HasIndexPlusRemoveIndex(removingIndexesCopyPtr, arg.Index)
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
		return linkedList
	}

	linkedList.setLength(nonChainedCollection.Length())
	linkedList.head = nonChainedCollection.ApplyChaining().First()

	return linkedList
}

func (linkedList *LinkedList) GetCompareSummary(
	right *LinkedList,
	leftName, rightName string,
) string {
	lLen := linkedList.Length()
	rLen := right.Length()

	leftStr := fmt.Sprintf(
		linkedListCollectionCompareHeaderLeft,
		leftName,
		lLen,
		linkedList)

	rightStr := fmt.Sprintf(
		linkedListCollectionCompareHeaderRight,
		rightName,
		rLen,
		right,
		linkedList.IsEqualsPtr(right),
		lLen,
		rLen)

	return leftStr + rightStr
}

// RemoveNode skip if removingNode is nil
func (linkedList *LinkedList) RemoveNode(
	removingNode *LinkedListNode,
) *LinkedList {
	if removingNode == nil {
		return linkedList
	}

	if linkedList.IsEmpty() {
		msgtype.
			CannotRemoveIndexesFromEmptyCollection.
			HandleUsingPanic("removingNode cannot be removed from empty linkedlist.", removingNode.String())
	}

	var processor LinkedListSimpleProcessor = func(
		arg *LinkedListProcessorParameter,
	) (isBreak bool) {
		isSameNode := arg.CurrentNode == removingNode
		if isSameNode && arg.IsFirstIndex {
			linkedList.head = arg.CurrentNode.next
			linkedList.decrementLength()

			return true
		}

		if isSameNode {
			arg.PrevNode.next = arg.CurrentNode.next
			linkedList.decrementLength()

			return true
		}

		return false
	}

	return linkedList.Loop(processor)
}

// AddStringsPtrToNode iSkipOnNil
func (linkedList *LinkedList) AddStringsPtrToNode(
	isSkipOnNull bool,
	node *LinkedListNode,
	items *[]string,
) *LinkedList {
	if items == nil || node == nil && isSkipOnNull {
		return linkedList
	}

	if node == nil {
		msgtype.
			CannotBeNilMessage.
			HandleUsingPanic(
				"node cannot be nil.",
				nil)
	}

	length := len(*items)

	if length == 0 {
		return linkedList
	}

	if length == 1 {
		linkedList.AddAfterNode(node, (*items)[0])

		return linkedList
	}

	finalHead := &LinkedListNode{
		Element: (*items)[0],
		next:    nil,
	}

	nextNode := finalHead

	for _, item := range (*items)[1:] {
		nextNode = nextNode.AddNext(linkedList, item)
	}

	//goland:noinspection GoNilness
	nextNode.next = node.next
	//goland:noinspection GoNilness
	node.next = finalHead
	linkedList.incrementLength()

	return linkedList
}

func (linkedList *LinkedList) AddAfterNode(
	node *LinkedListNode,
	item string,
) *LinkedListNode {
	newNode := &LinkedListNode{
		Element: item,
		next:    node.next,
	}

	node.next = newNode
	linkedList.incrementLength()

	return newNode
}

// AddStringsPtr add to back
func (linkedList *LinkedList) AddStringsPtr(items *[]string) *LinkedList {
	if items == nil {
		return linkedList
	}

	for _, item := range *items {
		linkedList.Add(item)
	}

	return linkedList
}

// AddStringsPtrLock add to back
func (linkedList *LinkedList) AddStringsPtrLock(items *[]string) *LinkedList {
	linkedList.Lock()
	defer linkedList.Unlock()

	return linkedList.AddStringsPtr(items)
}

// IndexAt Expensive operation BigO(n)
func (linkedList *LinkedList) IndexAt(index int) *LinkedListNode {
	length := linkedList.Length()
	if index < 0 {
		return nil
	}

	if length == 0 || length-1 < index {
		msgtype.OutOfRange.HandleUsingPanic(
			"Given index is out of range. Whereas length:",
			length)
	}

	if index == 0 {
		return linkedList.head
	}

	node := linkedList.head
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
func (linkedList *LinkedList) SafePointerIndexAt(index int) *string {
	node := linkedList.SafeIndexAt(index)

	if node == nil {
		return nil
	}

	return &node.Element
}

// SafePointerIndexAtUsingDefault Expensive operation BigO(n)
func (linkedList *LinkedList) SafePointerIndexAtUsingDefault(
	index int,
	defaultString string,
) string {
	node := linkedList.SafeIndexAt(index)

	if node == nil {
		return defaultString
	}

	return node.Element
}

// SafeIndexAt Expensive operation BigO(n)
func (linkedList *LinkedList) SafeIndexAt(index int) *LinkedListNode {
	length := linkedList.Length()
	isExitCondition := index < 0 || length == 0 || length-1 < index
	if isExitCondition {
		return nil
	}

	if index == 0 {
		return linkedList.head
	}

	node := linkedList.head
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
func (linkedList *LinkedList) SafeIndexAtLock(index int) *LinkedListNode {
	linkedList.Lock()
	defer linkedList.Unlock()

	return linkedList.SafeIndexAt(index)
}

// SafePointerIndexAtUsingDefaultLock Expensive operation BigO(n)
func (linkedList *LinkedList) SafePointerIndexAtUsingDefaultLock(
	index int,
	defaultString string,
) string {
	linkedList.Lock()
	defer linkedList.Unlock()

	return linkedList.SafePointerIndexAtUsingDefault(index, defaultString)
}

func (linkedList *LinkedList) GetNextNodes(count int) *[]*LinkedListNode {
	counter := 0

	return linkedList.Filter(
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

func (linkedList *LinkedList) GetAllLinkedNodes() *[]*LinkedListNode {
	return linkedList.Filter(
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
func (linkedList *LinkedList) AddPointerStringsPtr(items *[]*string) *LinkedList {
	if items == nil {
		return linkedList
	}

	for _, item := range *items {
		if item == nil {
			continue
		}

		linkedList.Add(*item)
	}

	return linkedList
}

// AddCollection skip on nil
func (linkedList *LinkedList) AddCollection(collection *Collection) *LinkedList {
	if collection == nil {
		return linkedList
	}

	for _, item := range *collection.items {
		linkedList.Add(item)
	}

	return linkedList
}

func (linkedList *LinkedList) ToCollection(addCapacity int) *Collection {
	newLength := linkedList.Length() + addCapacity
	collection := NewCollection(newLength)

	if linkedList.IsEmpty() {
		return collection
	}

	node := linkedList.head
	collection.Add(node.Element)

	for node.HasNext() {
		node = node.Next()
		collection.Add(node.Element)
	}

	return collection
}

// ListPtr must return slice.
func (linkedList *LinkedList) ListPtr() *[]string {
	list := make([]string, 0, linkedList.Length())

	if linkedList.IsEmpty() {
		return &list
	}

	node := linkedList.head
	list = append(list, node.Element)

	for node.HasNext() {
		node = node.Next()
		list = append(list, node.Element)
	}

	return &list
}

// ListPtrLock must return slice.
func (linkedList *LinkedList) ListPtrLock() *[]string {
	linkedList.Lock()
	defer linkedList.Unlock()

	return linkedList.ListPtr()
}

func (linkedList *LinkedList) String() string {
	if linkedList.IsEmpty() {
		return commonJoiner + NoElements
	}

	return commonJoiner +
		strings.Join(
			*linkedList.ListPtr(),
			commonJoiner)
}

func (linkedList *LinkedList) StringLock() string {
	if linkedList.IsEmptyLock() {
		return commonJoiner + NoElements
	}

	linkedList.Lock()
	defer linkedList.Unlock()

	return commonJoiner +
		strings.Join(
			*linkedList.ListPtr(),
			commonJoiner)
}

func (linkedList *LinkedList) Join(
	separator string,
) string {
	return strings.Join(*linkedList.ListPtr(), separator)
}

func (linkedList *LinkedList) JoinLock(
	separator string,
) string {
	linkedList.Lock()
	defer linkedList.Unlock()

	return strings.Join(*linkedList.ListPtr(), separator)
}

func (linkedList *LinkedList) Joins(
	separator string,
	items ...string,
) string {
	if items == nil || linkedList.Length() == 0 {
		return strings.Join(items, separator)
	}

	collection := linkedList.ToCollection(len(items) +
		constants.ArbitraryCapacity2)
	collection.AddStringsPtr(&items)

	return collection.Join(separator)
}

func (linkedList *LinkedList) JsonModel() *CollectionDataModel {
	return linkedList.ToCollection(0).JsonModel()
}

func (linkedList *LinkedList) JsonModelAny() interface{} {
	return linkedList.JsonModel()
}

func (linkedList *LinkedList) MarshalJSON() ([]byte, error) {
	return json.Marshal(*linkedList.JsonModel())
}

func (linkedList *LinkedList) UnmarshalJSON(data []byte) error {
	var dataModel CollectionDataModel
	err := json.Unmarshal(data, &dataModel)

	if err == nil {
		linkedList.Clear()
		linkedList.AddStringsPtr(dataModel.Items)
	}

	return err
}

func (linkedList *LinkedList) RemoveAll() *LinkedList {
	return linkedList.Clear()
}

func (linkedList *LinkedList) Clear() *LinkedList {
	if linkedList.IsEmpty() {
		return linkedList
	}

	linkedList.head = nil
	linkedList.tail = nil
	linkedList.setLengthToZero()

	return linkedList
}

func (linkedList *LinkedList) Json() *corejson.Result {
	if linkedList.IsEmpty() {
		return corejson.EmptyWithoutErrorPtr()
	}

	jsonBytes, err := json.Marshal(linkedList)

	return corejson.NewPtr(jsonBytes, err)
}

func (linkedList *LinkedList) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*LinkedList, error) {
	if jsonResult == nil || jsonResult.IsEmptyJsonBytes() {
		return EmptyLinkedList(), defaulterr.UnMarshallingFailedDueToNilOrEmpty
	}

	err := json.Unmarshal(*jsonResult.Bytes, &linkedList)

	if err != nil {
		return EmptyLinkedList(), err
	}

	return linkedList, nil
}

// ParseInjectUsingJsonMust Panic if error
func (linkedList *LinkedList) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *LinkedList {
	newUsingJson, err :=
		linkedList.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return newUsingJson
}

// JsonParseSelfInject Panic if error
func (linkedList *LinkedList) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := linkedList.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (linkedList *LinkedList) AsJsonMarshaller() corejson.JsonMarshaller {
	return linkedList
}
