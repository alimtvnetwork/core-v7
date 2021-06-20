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
	"gitlab.com/evatix-go/core/msgtype"
)

type LinkedCollections struct {
	head, tail *LinkedCollectionNode
	length     int
	sync.Mutex
}

func (linkedCollections *LinkedCollections) Tail() *LinkedCollectionNode {
	return linkedCollections.tail
}

func (linkedCollections *LinkedCollections) Head() *LinkedCollectionNode {
	return linkedCollections.head
}

func (linkedCollections *LinkedCollections) First() *Collection {
	return linkedCollections.head.Element
}

func (linkedCollections *LinkedCollections) Single() *Collection {
	return linkedCollections.head.Element
}

func (linkedCollections *LinkedCollections) Last() *Collection {
	return linkedCollections.tail.Element
}

func (linkedCollections *LinkedCollections) LastOrDefault() *Collection {
	if linkedCollections.IsEmpty() {
		return EmptyCollection()
	}

	return linkedCollections.tail.Element
}

func (linkedCollections *LinkedCollections) FirstOrDefault() *Collection {
	if linkedCollections.IsEmpty() {
		return EmptyCollection()
	}

	return linkedCollections.head.Element
}

func (linkedCollections *LinkedCollections) Length() int {
	return linkedCollections.length
}

// AllIndividualItemsLength including all nested ones
func (linkedCollections *LinkedCollections) AllIndividualItemsLength() int {
	allLengthSum := 0

	var processor LinkedCollectionSimpleProcessor = func(
		arg *LinkedCollectionProcessorParameter,
	) (isBreak bool) {
		allLengthSum += arg.CurrentNode.Element.Length()

		return false
	}

	linkedCollections.Loop(processor)

	return allLengthSum
}

func (linkedCollections *LinkedCollections) incrementLength() int {
	linkedCollections.length++

	return linkedCollections.length
}

func (linkedCollections *LinkedCollections) setLengthToZero() int {
	linkedCollections.length = 0

	return linkedCollections.length
}

func (linkedCollections *LinkedCollections) setLength(number int) int {
	linkedCollections.length = number

	return linkedCollections.length
}

func (linkedCollections *LinkedCollections) decrementLength() int {
	linkedCollections.length--

	return linkedCollections.length
}

func (linkedCollections *LinkedCollections) incrementLengthLock() {
	linkedCollections.Lock()
	linkedCollections.length++
	linkedCollections.Unlock()
}

func (linkedCollections *LinkedCollections) incrementLengthUsingNumber(number int) int {
	linkedCollections.length += number

	return linkedCollections.length
}

func (linkedCollections *LinkedCollections) LengthLock() int {
	linkedCollections.Lock()
	defer linkedCollections.Unlock()

	return linkedCollections.length
}

func (linkedCollections *LinkedCollections) IsEqualsPtr(
	anotherLinkedCollections *LinkedCollections,
) bool {
	if anotherLinkedCollections == nil {
		return false
	}

	if linkedCollections == anotherLinkedCollections {
		return true
	}

	if linkedCollections.IsEmpty() && anotherLinkedCollections.IsEmpty() {
		return true
	}

	if linkedCollections.IsEmpty() || anotherLinkedCollections.IsEmpty() {
		return false
	}

	if linkedCollections.Length() != anotherLinkedCollections.Length() {
		return false
	}

	leftNode := linkedCollections.head
	rightNode := anotherLinkedCollections.head

	if leftNode == nil && rightNode == nil {
		return true
	}

	if leftNode == nil || rightNode == nil {
		return false
	}

	return leftNode.IsChainEqual(rightNode)
}

func (linkedCollections *LinkedCollections) IsEmptyLock() bool {
	linkedCollections.Lock()
	defer linkedCollections.Unlock()

	return linkedCollections.head == nil || linkedCollections.length == 0
}

func (linkedCollections *LinkedCollections) IsEmpty() bool {
	return linkedCollections.head == nil || linkedCollections.length == 0
}

func (linkedCollections *LinkedCollections) HasItems() bool {
	return linkedCollections.head != nil &&
		linkedCollections.length > 0
}

// InsertAt BigO(n) expensive operation.
func (linkedCollections *LinkedCollections) InsertAt(
	index int,
	collection *Collection,
) *LinkedCollections {
	if index < 1 {
		return linkedCollections.AddFront(collection)
	}

	node := linkedCollections.IndexAt(index - 1)
	linkedCollections.AddAfterNode(node, collection)

	return linkedCollections
}

func (linkedCollections *LinkedCollections) AddAsync(
	collection *Collection,
	wg *sync.WaitGroup,
) *LinkedCollections {
	go func() {
		linkedCollections.Lock()
		defer linkedCollections.Unlock()
		linkedCollections.Add(collection)

		wg.Done()
	}()

	return linkedCollections
}

// AddsAsyncOnComplete Append back
func (linkedCollections *LinkedCollections) AddsAsyncOnComplete(
	onComplete OnCompleteLinkedCollections,
	isSkipOnNil bool,
	collections ...*Collection,
) *LinkedCollections {
	go func() {
		linkedCollections.Lock()
		defer linkedCollections.Unlock()

		linkedCollections.AppendCollectionsPointers(isSkipOnNil, &collections)

		onComplete(linkedCollections)
	}()

	return linkedCollections
}

// AddsUsingProcessorAsyncOnComplete Append back
func (linkedCollections *LinkedCollections) AddsUsingProcessorAsyncOnComplete(
	onComplete OnCompleteLinkedCollections,
	processor AnyToCollectionProcessor,
	isSkipOnNil bool,
	anys ...interface{},
) *LinkedCollections {
	go func() {
		linkedCollections.Lock()
		defer linkedCollections.Unlock()

		if anys == nil && isSkipOnNil {
			onComplete(linkedCollections)

			return
		}

		for i, any := range anys {
			if any == nil && isSkipOnNil {
				continue
			}

			collection := processor(any, i)
			linkedCollections.Add(collection)
		}

		onComplete(linkedCollections)
	}()

	return linkedCollections
}

// AddsUsingProcessorAsync Append back
func (linkedCollections *LinkedCollections) AddsUsingProcessorAsync(
	wg *sync.WaitGroup,
	processor AnyToCollectionProcessor,
	isSkipOnNil bool,
	anys ...interface{},
) *LinkedCollections {
	go func() {
		linkedCollections.Lock()
		defer linkedCollections.Unlock()

		if anys == nil && isSkipOnNil {
			wg.Done()

			return
		}

		for i, any := range anys {
			if any == nil && isSkipOnNil {
				continue
			}

			collection := processor(any, i)
			linkedCollections.Add(collection)
		}

		wg.Done()
	}()

	return linkedCollections
}

func (linkedCollections *LinkedCollections) AddLock(collection *Collection) *LinkedCollections {
	linkedCollections.Lock()
	defer linkedCollections.Unlock()

	return linkedCollections.Add(collection)
}

func (linkedCollections *LinkedCollections) Add(collection *Collection) *LinkedCollections {
	if linkedCollections.IsEmpty() {
		linkedCollections.head = &LinkedCollectionNode{
			Element: collection,
			next:    nil,
		}

		linkedCollections.tail = linkedCollections.head
		linkedCollections.incrementLength()

		return linkedCollections
	}

	linkedCollections.tail.next = &LinkedCollectionNode{
		Element: collection,
		next:    nil,
	}

	linkedCollections.tail = linkedCollections.tail.next
	linkedCollections.incrementLength()

	return linkedCollections
}

func (linkedCollections *LinkedCollections) AddStringsLock(stringsItems ...string) *LinkedCollections {
	if len(stringsItems) == 0 {
		return linkedCollections
	}

	linkedCollections.Lock()
	defer linkedCollections.Unlock()

	return linkedCollections.AddStringsPtr(&stringsItems, false)
}

func (linkedCollections *LinkedCollections) AddStrings(stringsItems ...string) *LinkedCollections {
	if len(stringsItems) == 0 {
		return linkedCollections
	}

	collection := NewCollectionUsingStrings(&stringsItems, false)

	return linkedCollections.Add(collection)
}

func (linkedCollections *LinkedCollections) AddBackNode(node *LinkedCollectionNode) *LinkedCollections {
	return linkedCollections.AppendNode(node)
}

func (linkedCollections *LinkedCollections) AppendNode(node *LinkedCollectionNode) *LinkedCollections {
	if linkedCollections.IsEmpty() {
		linkedCollections.head = node
		linkedCollections.tail = linkedCollections.head
		linkedCollections.incrementLength()

		return linkedCollections
	}

	linkedCollections.tail.next = node
	linkedCollections.tail = linkedCollections.tail.next
	linkedCollections.incrementLength()

	return linkedCollections
}

func (linkedCollections *LinkedCollections) AppendChainOfNodes(nodeHead *LinkedCollectionNode) *LinkedCollections {
	endOfChain, length := nodeHead.EndOfChain()

	if linkedCollections.IsEmpty() {
		linkedCollections.head = nodeHead
	} else {
		linkedCollections.tail.next = nodeHead
	}

	linkedCollections.tail = endOfChain
	linkedCollections.incrementLengthUsingNumber(length)

	return linkedCollections
}

func (linkedCollections *LinkedCollections) AppendChainOfNodesAsync(
	nodeHead *LinkedCollectionNode,
	wg *sync.WaitGroup,
) *LinkedCollections {
	go func() {
		linkedCollections.Lock()
		linkedCollections.AppendChainOfNodes(nodeHead)
		linkedCollections.Unlock()

		wg.Done()
	}()

	return linkedCollections
}

func (linkedCollections *LinkedCollections) PushBackLock(collection *Collection) *LinkedCollections {
	return linkedCollections.AddLock(collection)
}

func (linkedCollections *LinkedCollections) PushBack(collection *Collection) *LinkedCollections {
	return linkedCollections.Add(collection)
}

func (linkedCollections *LinkedCollections) Push(collection *Collection) *LinkedCollections {
	return linkedCollections.Add(collection)
}

func (linkedCollections *LinkedCollections) PushFront(collection *Collection) *LinkedCollections {
	return linkedCollections.AddFront(collection)
}

func (linkedCollections *LinkedCollections) AddFrontLock(collection *Collection) *LinkedCollections {
	linkedCollections.Lock()
	defer linkedCollections.Unlock()

	return linkedCollections.AddFront(collection)
}

func (linkedCollections *LinkedCollections) AddFront(collection *Collection) *LinkedCollections {
	if linkedCollections.IsEmpty() {
		return linkedCollections.Add(collection)
	}

	node := &LinkedCollectionNode{
		Element: collection,
		next:    linkedCollections.head,
	}

	linkedCollections.head = node
	linkedCollections.incrementLength()

	return linkedCollections
}

func (linkedCollections *LinkedCollections) AttachWithNode(
	currentNode,
	addingNode *LinkedCollectionNode,
) error {
	if currentNode == nil {
		return msgtype.
			CannotBeNilMessage.
			Error(currentNodeCannotBeNull, nil)
	}

	if currentNode.next != nil {
		return msgtype.
			ShouldBeNilMessage.
			Error("CurrentNode.next", nil)
	}

	addingNode.next = currentNode.next
	currentNode.next = addingNode
	linkedCollections.incrementLength()

	return nil
}

// AddCollectionToNode iSkipOnNil
func (linkedCollections *LinkedCollections) AddCollectionToNode(
	isSkipOnNull bool,
	node *LinkedCollectionNode,
	collection *Collection,
) *LinkedCollections {
	return linkedCollections.AddCollectionsToNode(
		isSkipOnNull,
		node,
		collection)
}

func (linkedCollections *LinkedCollections) GetNextNodes(count int) *[]*LinkedCollectionNode {
	counter := 0

	return linkedCollections.Filter(
		func(
			arg *LinkedCollectionFilterParameter,
		) *LinkedCollectionFilterResult {
			isBreak := counter >= count-1

			counter++
			return &LinkedCollectionFilterResult{
				Value:   arg.Node,
				IsKeep:  true,
				IsBreak: isBreak,
			}
		})
}

func (linkedCollections *LinkedCollections) GetAllLinkedNodes() *[]*LinkedCollectionNode {
	return linkedCollections.Filter(
		func(
			arg *LinkedCollectionFilterParameter,
		) *LinkedCollectionFilterResult {
			return &LinkedCollectionFilterResult{
				Value:   arg.Node,
				IsKeep:  true,
				IsBreak: false,
			}
		})
}

func (linkedCollections *LinkedCollections) Loop(
	simpleProcessor LinkedCollectionSimpleProcessor,
) *LinkedCollections {
	length := linkedCollections.Length()
	if length == 0 {
		return linkedCollections
	}

	node := linkedCollections.head
	arg := &LinkedCollectionProcessorParameter{
		Index:         0,
		CurrentNode:   node,
		PrevNode:      nil,
		IsFirstIndex:  true,
		IsEndingIndex: false,
	}

	isBreak := simpleProcessor(arg)

	if isBreak {
		return linkedCollections
	}

	lenMinusOne := length - 1
	index := 1
	isEndingIndex := false

	for node.HasNext() {
		prev := node
		node = node.Next()
		isEndingIndex = lenMinusOne == index

		arg2 := &LinkedCollectionProcessorParameter{
			Index:         index,
			CurrentNode:   node,
			PrevNode:      prev,
			IsFirstIndex:  false,
			IsEndingIndex: isEndingIndex,
		}

		isBreak = simpleProcessor(arg2)

		if isBreak {
			return linkedCollections
		}

		index++
	}

	return linkedCollections
}

func (linkedCollections *LinkedCollections) Filter(
	filter LinkedCollectionFilter,
) *[]*LinkedCollectionNode {
	length := linkedCollections.Length()
	list := make([]*LinkedCollectionNode, 0, length)

	if length == 0 {
		return &list
	}

	node := linkedCollections.head
	arg := &LinkedCollectionFilterParameter{
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

		arg2 := &LinkedCollectionFilterParameter{
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

func (linkedCollections *LinkedCollections) FilterAsCollection(
	filter LinkedCollectionFilter,
	additionalCapacity int,
) *Collection {
	items := linkedCollections.Filter(filter)

	if len(*items) == 0 {
		return EmptyCollection()
	}

	allLength := 0

	for _, node := range *items {
		if node != nil && node.Element != nil {
			allLength += node.Element.Length()
		}
	}

	collection := NewCollection(allLength + additionalCapacity)

	for _, node := range *items {
		if node == nil || node.Element == nil {
			continue
		}

		collection.AddCollection(node.Element)
	}

	return collection
}

func (linkedCollections *LinkedCollections) FilterAsCollections(
	filter LinkedCollectionFilter,
) *[]*Collection {
	items := linkedCollections.Filter(filter)
	collections := make([]*Collection, len(*items))

	for i := range *items {
		collections[i] = (*items)[i].Element
	}

	return &collections
}

func (linkedCollections *LinkedCollections) RemoveNodeByIndex(
	removingIndex int,
) *LinkedCollections {
	if removingIndex < 0 {
		msgtype.
			CannotBeNegativeIndex.
			HandleUsingPanic(
				"removeIndex was less than 0.",
				removingIndex)
	}

	var singleProcessor LinkedCollectionSimpleProcessor = func(
		arg *LinkedCollectionProcessorParameter,
	) (isBreak bool) {
		hasIndex := removingIndex == arg.Index

		if !hasIndex {
			return false
		}

		isBreak = hasIndex
		linkedCollections.decrementLength()

		if arg.IsFirstIndex {
			linkedCollections.head =
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

	return linkedCollections.Loop(singleProcessor)
}

func (linkedCollections *LinkedCollections) RemoveNodeByIndexes(
	isIgnorePanic bool,
	removingIndexes ...int,
) *LinkedCollections {
	length := len(removingIndexes)

	if length == 0 {
		return linkedCollections
	}

	if !isIgnorePanic && linkedCollections.IsEmpty() && length > 0 {
		msgtype.
			CannotRemoveIndexesFromEmptyCollection.
			HandleUsingPanic("removingIndexes cannot be removed from empty LinkedCollections.", removingIndexes)
	}

	removingIndexesCopy := removingIndexes
	removingIndexesCopyPtr := &removingIndexesCopy

	nonChainedNodes := linkedCollections.Filter(
		func(arg *LinkedCollectionFilterParameter) *LinkedCollectionFilterResult {
			hasIndex := coreindexes.HasIndexPlusRemoveIndex(removingIndexesCopyPtr, arg.Index)
			if hasIndex {
				// remove
				return &LinkedCollectionFilterResult{
					Value:   arg.Node,
					IsKeep:  false,
					IsBreak: false,
				}
			}

			// not remove
			return &LinkedCollectionFilterResult{
				Value:   arg.Node,
				IsKeep:  true,
				IsBreak: false,
			}
		})

	nonChainedCollection := &NonChainedLinkedCollectionNodes{
		items:             nonChainedNodes,
		isChainingApplied: false,
	}

	if nonChainedCollection.IsEmpty() {
		return linkedCollections
	}

	linkedCollections.setLength(nonChainedCollection.Length())
	linkedCollections.head = nonChainedCollection.ApplyChaining().First()

	return linkedCollections
}

func (linkedCollections *LinkedCollections) RemoveNode(
	removingNode *LinkedCollectionNode,
) *LinkedCollections {
	var processor LinkedCollectionSimpleProcessor = func(
		arg *LinkedCollectionProcessorParameter,
	) (isBreak bool) {
		isSameNode := arg.CurrentNode == removingNode
		if isSameNode && arg.IsFirstIndex {
			linkedCollections.head = arg.CurrentNode.next
			linkedCollections.decrementLength()

			return true
		}

		if isSameNode {
			arg.PrevNode.next = arg.CurrentNode.next
			linkedCollections.decrementLength()

			return true
		}

		return false
	}

	return linkedCollections.Loop(processor)
}

// AppendCollections iSkipOnNil
func (linkedCollections *LinkedCollections) AppendCollections(
	isSkipOnNull bool,
	collections ...*Collection,
) *LinkedCollections {
	if isSkipOnNull && collections == nil {
		return linkedCollections
	}

	for i := range collections {
		collection := collections[i]
		if isSkipOnNull && collection == nil {
			continue
		}

		linkedCollections.Add(collection)
	}

	return linkedCollections
}

// AppendCollectionsPointersLock iSkipOnNil
func (linkedCollections *LinkedCollections) AppendCollectionsPointersLock(
	isSkipOnNull bool,
	collections *[]*Collection,
) *LinkedCollections {
	if isSkipOnNull && collections == nil {
		return linkedCollections
	}

	for i := range *collections {
		collection := (*collections)[i]
		if isSkipOnNull && collection == nil {
			continue
		}

		linkedCollections.AddLock(collection)
	}

	return linkedCollections
}

// AppendCollectionsPointers iSkipOnNil
func (linkedCollections *LinkedCollections) AppendCollectionsPointers(
	isSkipOnNull bool,
	collections *[]*Collection,
) *LinkedCollections {
	if isSkipOnNull && collections == nil {
		return linkedCollections
	}

	for i := range *collections {
		collection := (*collections)[i]
		if isSkipOnNull && collection == nil {
			continue
		}

		linkedCollections.Add(collection)
	}

	return linkedCollections
}

// AddCollectionsToNodeAsync iSkipOnNil
func (linkedCollections *LinkedCollections) AddCollectionsToNodeAsync(
	isSkipOnNull bool,
	wg *sync.WaitGroup,
	node *LinkedCollectionNode,
	collections ...*Collection,
) *LinkedCollections {
	if isSkipOnNull && collections == nil {
		return linkedCollections
	}

	go func() {
		linkedCollections.Lock()
		linkedCollections.AddCollectionsPointerToNode(
			isSkipOnNull,
			node,
			&collections)

		linkedCollections.Unlock()

		wg.Done()
	}()

	return linkedCollections
}

// AddCollectionsToNode iSkipOnNil
func (linkedCollections *LinkedCollections) AddCollectionsToNode(
	isSkipOnNull bool,
	node *LinkedCollectionNode,
	collections ...*Collection,
) *LinkedCollections {
	if isSkipOnNull && collections == nil {
		return linkedCollections
	}

	return linkedCollections.AddCollectionsPointerToNode(
		isSkipOnNull,
		node,
		&collections)
}

// AddCollectionsPointerToNode iSkipOnNil
func (linkedCollections *LinkedCollections) AddCollectionsPointerToNode(
	isSkipOnNull bool,
	node *LinkedCollectionNode,
	items *[]*Collection,
) *LinkedCollections {
	if items == nil || node == nil && isSkipOnNull {
		return linkedCollections
	}

	if node == nil {
		msgtype.
			CannotBeNilMessage.
			HandleUsingPanic(
				nodesCannotBeNull,
				nil)
	}

	length := len(*items)

	if length == 0 {
		return linkedCollections
	}

	if length == 1 {
		linkedCollections.AddAfterNode(node, (*items)[0])

		return linkedCollections
	}

	finalHead := &LinkedCollectionNode{
		Element: (*items)[0],
		next:    nil,
	}

	nextNode := finalHead

	for _, collection := range (*items)[1:] {
		if isSkipOnNull && collection == nil {
			continue
		}

		nextNode = nextNode.AddNext(linkedCollections, collection)
	}

	//goland:noinspection GoNilness
	nextNode.next = node.next
	//goland:noinspection GoNilness
	node.next = finalHead
	linkedCollections.incrementLength()

	return linkedCollections
}

func (linkedCollections *LinkedCollections) AddAfterNode(
	node *LinkedCollectionNode,
	collection *Collection,
) *LinkedCollectionNode {
	newNode := &LinkedCollectionNode{
		Element: collection,
		next:    node.next,
	}

	node.next = newNode
	linkedCollections.incrementLength()

	return newNode
}

func (linkedCollections *LinkedCollections) AddAfterNodeAsync(
	wg *sync.WaitGroup,
	node *LinkedCollectionNode,
	collection *Collection,
) {
	go func() {
		linkedCollections.Lock()

		linkedCollections.AddAfterNode(node, collection)

		linkedCollections.Unlock()

		wg.Done()
	}()
}

// AddStringsPtrAsync add to back
func (linkedCollections *LinkedCollections) AddStringsPtrAsync(
	wg *sync.WaitGroup,
	items *[]string,
	isMakeClone bool,
) *LinkedCollections {
	if items == nil {
		return linkedCollections
	}

	go func() {
		collection := NewCollectionUsingStrings(items, isMakeClone)

		linkedCollections.Lock()

		linkedCollections.Add(collection)

		linkedCollections.Unlock()

		wg.Done()
	}()

	return linkedCollections
}

// AddAsyncFuncItems must add all the lengths to the wg
func (linkedCollections *LinkedCollections) AddAsyncFuncItems(
	wg *sync.WaitGroup,
	isMakeClone bool,
	asyncFunctions ...func() []string,
) *LinkedCollections {
	if asyncFunctions == nil {
		return linkedCollections
	}

	asyncFuncWrap := func(asyncFunc func() []string) {
		items := asyncFunc()
		collection := NewCollectionUsingStrings(&items, isMakeClone)

		linkedCollections.Lock()
		linkedCollections.Add(collection)
		linkedCollections.Unlock()

		wg.Done()
	}

	for _, function := range asyncFunctions {
		go asyncFuncWrap(function)
	}

	wg.Wait()

	return linkedCollections
}

// AddStringsPtr add to back
func (linkedCollections *LinkedCollections) AddStringsPtr(
	items *[]string,
	isMakeClone bool,
) *LinkedCollections {
	if items == nil {
		return linkedCollections
	}

	collection := NewCollectionUsingStrings(items, isMakeClone)

	return linkedCollections.Add(collection)
}

// IndexAt Expensive operation BigO(n)
func (linkedCollections *LinkedCollections) IndexAt(
	index int,
) *LinkedCollectionNode {
	length := linkedCollections.Length()
	if index < 0 {
		return nil
	}

	if length == 0 || length-1 < index {
		msgtype.OutOfRange.HandleUsingPanic(
			"Given index is out of range. Whereas length:",
			length)
	}

	if index == 0 {
		return linkedCollections.head
	}

	node := linkedCollections.head
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
func (linkedCollections *LinkedCollections) SafePointerIndexAt(
	index int,
) *Collection {
	node := linkedCollections.SafeIndexAt(index)

	if node == nil {
		return nil
	}

	return node.Element
}

// SafeIndexAt Expensive operation BigO(n)
func (linkedCollections *LinkedCollections) SafeIndexAt(
	index int,
) *LinkedCollectionNode {
	length := linkedCollections.Length()
	isExitCondition := index < 0 || length == 0 || length-1 < index
	if isExitCondition {
		return nil
	}

	if index == 0 {
		return linkedCollections.head
	}

	node := linkedCollections.head
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

// AddPointerStringsPtr skip on nil, add to back
func (linkedCollections *LinkedCollections) AddPointerStringsPtr(
	items *[]*string,
) *LinkedCollections {
	if items == nil {
		return linkedCollections
	}

	collection := NewCollectionUsingPointerStringsPlusCap(
		items,
		constants.Zero)

	return linkedCollections.Add(collection)
}

// AddPointerStringsPtrAsync skip on nil, add to back
func (linkedCollections *LinkedCollections) AddPointerStringsPtrAsync(
	wg *sync.WaitGroup,
	items *[]*string,
) *LinkedCollections {
	if items == nil {
		return linkedCollections
	}

	go func() {
		collection := NewCollectionUsingPointerStringsPlusCap(
			items,
			constants.Zero)

		linkedCollections.Lock()
		linkedCollections.Add(collection)
		linkedCollections.Unlock()

		wg.Done()
	}()

	return linkedCollections
}

// AddCollection skip on nil
func (linkedCollections *LinkedCollections) AddCollection(
	collection *Collection,
) *LinkedCollections {
	if collection == nil {
		return linkedCollections
	}

	return linkedCollections.Add(collection)
}

func (linkedCollections *LinkedCollections) ToCollection(
	addCapacity int,
) *Collection {
	if linkedCollections.IsEmpty() {
		return EmptyCollection()
	}

	newLength := linkedCollections.AllIndividualItemsLength() +
		addCapacity

	collection := NewCollection(newLength)
	var processor LinkedCollectionSimpleProcessor = func(
		arg *LinkedCollectionProcessorParameter,
	) (isBreak bool) {
		if arg.CurrentNode == nil {
			return false
		}

		collection.AddCollection(arg.CurrentNode.Element)

		return false
	}

	linkedCollections.Loop(processor)

	return collection
}

func (linkedCollections *LinkedCollections) ToCollectionsOfCollection(
	addCapacity int,
) *CollectionsOfCollection {
	if linkedCollections.IsEmpty() {
		return EmptyCollectionsOfCollection()
	}

	newLength := linkedCollections.AllIndividualItemsLength() +
		addCapacity

	collection := NewCollectionsOfCollection(newLength)

	var processor LinkedCollectionSimpleProcessor = func(
		arg *LinkedCollectionProcessorParameter,
	) (isBreak bool) {
		if arg.CurrentNode == nil {
			return false
		}

		collection.Adds(arg.CurrentNode.Element)

		return false
	}

	linkedCollections.Loop(processor)

	return collection
}

func (linkedCollections *LinkedCollections) ItemsOfItems() *[]*[]string {
	length := linkedCollections.Length()
	itemsOfItems := make([]*[]string, length)

	if length == 0 {
		return &itemsOfItems
	}

	nodes := linkedCollections.GetAllLinkedNodes()

	for i, node := range *nodes {
		itemsOfItems[i] = node.Element.items
	}

	return &itemsOfItems
}

func (linkedCollections *LinkedCollections) ItemsOfItemsCollection() *[]*Collection {
	length := linkedCollections.Length()
	itemsOfItems := make([]*Collection, length)

	if length == 0 {
		return &itemsOfItems
	}

	nodes := linkedCollections.GetAllLinkedNodes()

	for i, node := range *nodes {
		itemsOfItems[i] = node.Element
	}

	return &itemsOfItems
}

// ListPtr must return slice.
func (linkedCollections *LinkedCollections) ListPtr() *[]string {
	return linkedCollections.
		ToCollection(constants.ArbitraryCapacity5).
		items
}

func (linkedCollections *LinkedCollections) String() string {
	if linkedCollections.IsEmpty() {
		return commonJoiner + NoElements
	}

	collections := *linkedCollections.ToCollectionsOfCollection(0)

	return collections.String()
}

func (linkedCollections *LinkedCollections) StringLock() string {
	if linkedCollections.IsEmptyLock() {
		return commonJoiner + NoElements
	}

	linkedCollections.Lock()
	defer linkedCollections.Unlock()

	return commonJoiner +
		strings.Join(
			*linkedCollections.ListPtr(),
			commonJoiner)
}

func (linkedCollections *LinkedCollections) Join(
	separator string,
) string {
	return strings.Join(*linkedCollections.ListPtr(), separator)
}

func (linkedCollections *LinkedCollections) Joins(
	separator string,
	items ...string,
) string {
	if items == nil || linkedCollections.Length() == 0 {
		return strings.Join(items, separator)
	}

	collection := linkedCollections.ToCollection(len(items) +
		constants.ArbitraryCapacity2)
	collection.AddStringsPtr(&items)

	return collection.Join(separator)
}

func (linkedCollections *LinkedCollections) JsonModel() *CollectionDataModel {
	return linkedCollections.ToCollection(0).JsonModel()
}

func (linkedCollections *LinkedCollections) JsonModelAny() interface{} {
	return linkedCollections.JsonModel()
}

func (linkedCollections *LinkedCollections) MarshalJSON() ([]byte, error) {
	return json.Marshal(*linkedCollections.JsonModel())
}

func (linkedCollections *LinkedCollections) UnmarshalJSON(data []byte) error {
	var dataModel CollectionDataModel
	err := json.Unmarshal(data, &dataModel)

	if err == nil {
		linkedCollections.Clear()
		linkedCollections.AddStringsPtr(dataModel.Items, false)
	}

	return err
}

func (linkedCollections *LinkedCollections) RemoveAll() *LinkedCollections {
	return linkedCollections.Clear()
}

func (linkedCollections *LinkedCollections) Clear() *LinkedCollections {
	if linkedCollections.IsEmpty() {
		return linkedCollections
	}

	linkedCollections.head = nil
	linkedCollections.tail = nil
	linkedCollections.setLengthToZero()

	return linkedCollections
}

func (linkedCollections *LinkedCollections) Json() *corejson.Result {
	if linkedCollections.IsEmpty() {
		return corejson.EmptyWithoutErrorPtr()
	}

	jsonBytes, err := json.Marshal(linkedCollections)

	return corejson.NewPtr(jsonBytes, err)
}

func (linkedCollections *LinkedCollections) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*LinkedCollections, error) {
	if jsonResult == nil || jsonResult.IsEmptyJsonBytes() {
		return EmptyLinkedCollections(), defaulterr.UnMarshallingFailedDueToNilOrEmpty
	}

	err := json.Unmarshal(*jsonResult.Bytes, &linkedCollections)

	if err != nil {
		return EmptyLinkedCollections(), err
	}

	return linkedCollections, nil
}

// ParseInjectUsingJsonMust Panic if error
func (linkedCollections *LinkedCollections) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *LinkedCollections {
	newUsingJson, err :=
		linkedCollections.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return newUsingJson
}

func (linkedCollections *LinkedCollections) GetCompareSummary(
	right *LinkedCollections, leftName, rightName string,
) string {
	lLen := linkedCollections.Length()
	rLen := right.Length()

	leftStr := fmt.Sprintf(
		linkedListCollectionCompareHeaderLeft,
		leftName,
		lLen,
		linkedCollections)

	rightStr := fmt.Sprintf(
		linkedListCollectionCompareHeaderRight,
		rightName,
		rLen,
		right,
		linkedCollections.IsEqualsPtr(right),
		lLen,
		rLen)

	return leftStr + rightStr
}

func (linkedCollections *LinkedCollections) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := linkedCollections.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (linkedCollections *LinkedCollections) AsJsoner() corejson.Jsoner {
	return linkedCollections
}

func (linkedCollections *LinkedCollections) AsJsonParseSelfInjector() corejson.JsonParseSelfInjector {
	return linkedCollections
}

func (linkedCollections *LinkedCollections) AsJsonMarshaller() corejson.JsonMarshaller {
	return linkedCollections
}
