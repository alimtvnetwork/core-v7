package corestr

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coreindexes"
	"gitlab.com/evatix-go/core/msgtype"
)

type LinkedCollections struct {
	head, tail *LinkedCollectionNode
	length     int
	sync.Mutex
}

func (it *LinkedCollections) Tail() *LinkedCollectionNode {
	return it.tail
}

func (it *LinkedCollections) Head() *LinkedCollectionNode {
	return it.head
}

func (it *LinkedCollections) First() *Collection {
	return it.head.Element
}

func (it *LinkedCollections) Single() *Collection {
	return it.head.Element
}

func (it *LinkedCollections) Last() *Collection {
	return it.tail.Element
}

func (it *LinkedCollections) LastOrDefault() *Collection {
	if it.IsEmpty() {
		return EmptyCollection()
	}

	return it.tail.Element
}

func (it *LinkedCollections) FirstOrDefault() *Collection {
	if it.IsEmpty() {
		return EmptyCollection()
	}

	return it.head.Element
}

func (it *LinkedCollections) Length() int {
	return it.length
}

// AllIndividualItemsLength including all nested ones
func (it *LinkedCollections) AllIndividualItemsLength() int {
	allLengthSum := 0

	var processor LinkedCollectionSimpleProcessor = func(
		arg *LinkedCollectionProcessorParameter,
	) (isBreak bool) {
		allLengthSum += arg.CurrentNode.Element.Length()

		return false
	}

	it.Loop(processor)

	return allLengthSum
}

func (it *LinkedCollections) incrementLength() int {
	it.length++

	return it.length
}

func (it *LinkedCollections) setLengthToZero() int {
	it.length = 0

	return it.length
}

func (it *LinkedCollections) setLength(number int) int {
	it.length = number

	return it.length
}

func (it *LinkedCollections) decrementLength() int {
	it.length--

	return it.length
}

func (it *LinkedCollections) incrementLengthLock() {
	it.Lock()
	it.length++
	it.Unlock()
}

func (it *LinkedCollections) incrementLengthUsingNumber(number int) int {
	it.length += number

	return it.length
}

func (it *LinkedCollections) LengthLock() int {
	it.Lock()
	defer it.Unlock()

	return it.length
}

func (it *LinkedCollections) IsEqualsPtr(
	anotherLinkedCollections *LinkedCollections,
) bool {
	if anotherLinkedCollections == nil {
		return false
	}

	if it == anotherLinkedCollections {
		return true
	}

	if it.IsEmpty() && anotherLinkedCollections.IsEmpty() {
		return true
	}

	if it.IsEmpty() || anotherLinkedCollections.IsEmpty() {
		return false
	}

	if it.Length() != anotherLinkedCollections.Length() {
		return false
	}

	leftNode := it.head
	rightNode := anotherLinkedCollections.head

	if leftNode == nil && rightNode == nil {
		return true
	}

	if leftNode == nil || rightNode == nil {
		return false
	}

	return leftNode.IsChainEqual(rightNode)
}

func (it *LinkedCollections) IsEmptyLock() bool {
	it.Lock()
	defer it.Unlock()

	return it.head == nil || it.length == 0
}

func (it *LinkedCollections) IsEmpty() bool {
	return it.head == nil || it.length == 0
}

func (it *LinkedCollections) HasItems() bool {
	return it.head != nil &&
		it.length > 0
}

// InsertAt BigO(n) expensive operation.
func (it *LinkedCollections) InsertAt(
	index int,
	collection *Collection,
) *LinkedCollections {
	if index < 1 {
		return it.AddFront(collection)
	}

	node := it.IndexAt(index - 1)
	it.AddAfterNode(node, collection)

	return it
}

func (it *LinkedCollections) AddAsync(
	collection *Collection,
	wg *sync.WaitGroup,
) *LinkedCollections {
	go func() {
		it.Lock()
		defer it.Unlock()
		it.Add(collection)

		wg.Done()
	}()

	return it
}

// AddsAsyncOnComplete Append back
func (it *LinkedCollections) AddsAsyncOnComplete(
	onComplete OnCompleteLinkedCollections,
	isSkipOnNil bool,
	collections ...*Collection,
) *LinkedCollections {
	go func() {
		it.Lock()
		defer it.Unlock()

		it.AppendCollectionsPointers(isSkipOnNil, &collections)

		onComplete(it)
	}()

	return it
}

// AddsUsingProcessorAsyncOnComplete Append back
func (it *LinkedCollections) AddsUsingProcessorAsyncOnComplete(
	onComplete OnCompleteLinkedCollections,
	processor AnyToCollectionProcessor,
	isSkipOnNil bool,
	anys ...interface{},
) *LinkedCollections {
	go func() {
		it.Lock()
		defer it.Unlock()

		if anys == nil && isSkipOnNil {
			onComplete(it)

			return
		}

		for i, any := range anys {
			if any == nil && isSkipOnNil {
				continue
			}

			collection := processor(any, i)
			it.Add(collection)
		}

		onComplete(it)
	}()

	return it
}

// AddsUsingProcessorAsync Append back
func (it *LinkedCollections) AddsUsingProcessorAsync(
	wg *sync.WaitGroup,
	processor AnyToCollectionProcessor,
	isSkipOnNil bool,
	anys ...interface{},
) *LinkedCollections {
	go func() {
		it.Lock()
		defer it.Unlock()

		if anys == nil && isSkipOnNil {
			wg.Done()

			return
		}

		for i, any := range anys {
			if any == nil && isSkipOnNil {
				continue
			}

			collection := processor(any, i)
			it.Add(collection)
		}

		wg.Done()
	}()

	return it
}

func (it *LinkedCollections) AddLock(collection *Collection) *LinkedCollections {
	it.Lock()
	defer it.Unlock()

	return it.Add(collection)
}

func (it *LinkedCollections) Add(collection *Collection) *LinkedCollections {
	if it.IsEmpty() {
		it.head = &LinkedCollectionNode{
			Element: collection,
			next:    nil,
		}

		it.tail = it.head
		it.incrementLength()

		return it
	}

	it.tail.next = &LinkedCollectionNode{
		Element: collection,
		next:    nil,
	}

	it.tail = it.tail.next
	it.incrementLength()

	return it
}

func (it *LinkedCollections) AddStringsLock(stringsItems ...string) *LinkedCollections {
	if len(stringsItems) == 0 {
		return it
	}

	it.Lock()
	defer it.Unlock()

	return it.AddStringsPtr(&stringsItems, false)
}

func (it *LinkedCollections) AddStrings(stringsItems ...string) *LinkedCollections {
	if len(stringsItems) == 0 {
		return it
	}

	collection := NewCollectionUsingStrings(stringsItems, false)

	return it.Add(collection)
}

func (it *LinkedCollections) AddBackNode(node *LinkedCollectionNode) *LinkedCollections {
	return it.AppendNode(node)
}

func (it *LinkedCollections) AppendNode(node *LinkedCollectionNode) *LinkedCollections {
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

func (it *LinkedCollections) AppendChainOfNodes(nodeHead *LinkedCollectionNode) *LinkedCollections {
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

func (it *LinkedCollections) AppendChainOfNodesAsync(
	nodeHead *LinkedCollectionNode,
	wg *sync.WaitGroup,
) *LinkedCollections {
	go func() {
		it.Lock()
		it.AppendChainOfNodes(nodeHead)
		it.Unlock()

		wg.Done()
	}()

	return it
}

func (it *LinkedCollections) PushBackLock(collection *Collection) *LinkedCollections {
	return it.AddLock(collection)
}

func (it *LinkedCollections) PushBack(collection *Collection) *LinkedCollections {
	return it.Add(collection)
}

func (it *LinkedCollections) Push(collection *Collection) *LinkedCollections {
	return it.Add(collection)
}

func (it *LinkedCollections) PushFront(collection *Collection) *LinkedCollections {
	return it.AddFront(collection)
}

func (it *LinkedCollections) AddFrontLock(collection *Collection) *LinkedCollections {
	it.Lock()
	defer it.Unlock()

	return it.AddFront(collection)
}

func (it *LinkedCollections) AddFront(collection *Collection) *LinkedCollections {
	if it.IsEmpty() {
		return it.Add(collection)
	}

	node := &LinkedCollectionNode{
		Element: collection,
		next:    it.head,
	}

	it.head = node
	it.incrementLength()

	return it
}

func (it *LinkedCollections) AttachWithNode(
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
	it.incrementLength()

	return nil
}

func (it *LinkedCollections) AddAnother(
	another *LinkedCollections,
) *LinkedCollections {
	if another == nil || another.IsEmpty() {
		return it
	}

	node := another.Head()
	it.Add(node.Element)

	for node.HasNext() {
		node = node.Next()

		it.Add(node.Element)
	}

	return it
}

// AddCollectionToNode iSkipOnNil
func (it *LinkedCollections) AddCollectionToNode(
	isSkipOnNull bool,
	node *LinkedCollectionNode,
	collection *Collection,
) *LinkedCollections {
	return it.AddCollectionsToNode(
		isSkipOnNull,
		node,
		collection)
}

func (it *LinkedCollections) GetNextNodes(count int) *[]*LinkedCollectionNode {
	counter := 0

	return it.Filter(
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

func (it *LinkedCollections) GetAllLinkedNodes() *[]*LinkedCollectionNode {
	return it.Filter(
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

func (it *LinkedCollections) Loop(
	simpleProcessor LinkedCollectionSimpleProcessor,
) *LinkedCollections {
	length := it.Length()
	if length == 0 {
		return it
	}

	node := it.head
	arg := &LinkedCollectionProcessorParameter{
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

		arg2 := &LinkedCollectionProcessorParameter{
			Index:         index,
			CurrentNode:   node,
			PrevNode:      prev,
			IsFirstIndex:  false,
			IsEndingIndex: isEndingIndex,
		}

		isBreak = simpleProcessor(arg2)

		if isBreak {
			return it
		}

		index++
	}

	return it
}

func (it *LinkedCollections) Filter(
	filter LinkedCollectionFilter,
) *[]*LinkedCollectionNode {
	length := it.Length()
	list := make([]*LinkedCollectionNode, 0, length)

	if length == 0 {
		return &list
	}

	node := it.head
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

func (it *LinkedCollections) FilterAsCollection(
	filter LinkedCollectionFilter,
	additionalCapacity int,
) *Collection {
	items := it.Filter(filter)

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

func (it *LinkedCollections) FilterAsCollections(
	filter LinkedCollectionFilter,
) *[]*Collection {
	items := it.Filter(filter)
	collections := make([]*Collection, len(*items))

	for i := range *items {
		collections[i] = (*items)[i].Element
	}

	return &collections
}

func (it *LinkedCollections) RemoveNodeByIndex(
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

func (it *LinkedCollections) RemoveNodeByIndexes(
	isIgnorePanic bool,
	removingIndexes ...int,
) *LinkedCollections {
	length := len(removingIndexes)

	if length == 0 {
		return it
	}

	if !isIgnorePanic && it.IsEmpty() && length > 0 {
		msgtype.
			CannotRemoveIndexesFromEmptyCollection.
			HandleUsingPanic("removingIndexes cannot be removed from empty LinkedCollections.", removingIndexes)
	}

	nonChainedNodes := it.Filter(
		func(arg *LinkedCollectionFilterParameter) *LinkedCollectionFilterResult {
			hasIndex := coreindexes.HasIndexPlusRemoveIndex(removingIndexes, arg.Index)
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
		return it
	}

	it.setLength(nonChainedCollection.Length())
	it.head = nonChainedCollection.ApplyChaining().First()

	return it
}

func (it *LinkedCollections) RemoveNode(
	removingNode *LinkedCollectionNode,
) *LinkedCollections {
	var processor LinkedCollectionSimpleProcessor = func(
		arg *LinkedCollectionProcessorParameter,
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

// AppendCollections iSkipOnNil
func (it *LinkedCollections) AppendCollections(
	isSkipOnNull bool,
	collections ...*Collection,
) *LinkedCollections {
	if isSkipOnNull && collections == nil {
		return it
	}

	for i := range collections {
		collection := collections[i]
		if isSkipOnNull && collection == nil {
			continue
		}

		it.Add(collection)
	}

	return it
}

// AppendCollectionsPointersLock iSkipOnNil
func (it *LinkedCollections) AppendCollectionsPointersLock(
	isSkipOnNull bool,
	collections *[]*Collection,
) *LinkedCollections {
	if isSkipOnNull && collections == nil {
		return it
	}

	for i := range *collections {
		collection := (*collections)[i]
		if isSkipOnNull && collection == nil {
			continue
		}

		it.AddLock(collection)
	}

	return it
}

// AppendCollectionsPointers iSkipOnNil
func (it *LinkedCollections) AppendCollectionsPointers(
	isSkipOnNull bool,
	collections *[]*Collection,
) *LinkedCollections {
	if isSkipOnNull && collections == nil {
		return it
	}

	for i := range *collections {
		collection := (*collections)[i]
		if isSkipOnNull && collection == nil {
			continue
		}

		it.Add(collection)
	}

	return it
}

// AddCollectionsToNodeAsync iSkipOnNil
func (it *LinkedCollections) AddCollectionsToNodeAsync(
	isSkipOnNull bool,
	wg *sync.WaitGroup,
	node *LinkedCollectionNode,
	collections ...*Collection,
) *LinkedCollections {
	if isSkipOnNull && collections == nil {
		return it
	}

	go func() {
		it.Lock()
		it.AddCollectionsPointerToNode(
			isSkipOnNull,
			node,
			&collections)

		it.Unlock()

		wg.Done()
	}()

	return it
}

// AddCollectionsToNode iSkipOnNil
func (it *LinkedCollections) AddCollectionsToNode(
	isSkipOnNull bool,
	node *LinkedCollectionNode,
	collections ...*Collection,
) *LinkedCollections {
	if isSkipOnNull && collections == nil {
		return it
	}

	return it.AddCollectionsPointerToNode(
		isSkipOnNull,
		node,
		&collections)
}

// AddCollectionsPointerToNode iSkipOnNil
func (it *LinkedCollections) AddCollectionsPointerToNode(
	isSkipOnNull bool,
	node *LinkedCollectionNode,
	items *[]*Collection,
) *LinkedCollections {
	if items == nil || node == nil && isSkipOnNull {
		return it
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
		return it
	}

	if length == 1 {
		it.AddAfterNode(node, (*items)[0])

		return it
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

		nextNode = nextNode.AddNext(it, collection)
	}

	//goland:noinspection GoNilness
	nextNode.next = node.next
	//goland:noinspection GoNilness
	node.next = finalHead
	it.incrementLength()

	return it
}

func (it *LinkedCollections) AddAfterNode(
	node *LinkedCollectionNode,
	collection *Collection,
) *LinkedCollectionNode {
	newNode := &LinkedCollectionNode{
		Element: collection,
		next:    node.next,
	}

	node.next = newNode
	it.incrementLength()

	return newNode
}

func (it *LinkedCollections) AddAfterNodeAsync(
	wg *sync.WaitGroup,
	node *LinkedCollectionNode,
	collection *Collection,
) {
	go func() {
		it.Lock()

		it.AddAfterNode(node, collection)

		it.Unlock()

		wg.Done()
	}()
}

// AddStringsPtrAsync add to back
func (it *LinkedCollections) AddStringsPtrAsync(
	wg *sync.WaitGroup,
	items *[]string,
	isMakeClone bool,
) *LinkedCollections {
	if items == nil {
		return it
	}

	go func() {
		collection := NewCollectionUsingStringsPtr(items, isMakeClone)

		it.Lock()

		it.Add(collection)

		it.Unlock()

		wg.Done()
	}()

	return it
}

func (it *LinkedCollections) ConcatNew(
	isMakeCloneOnEmpty bool,
	linkedCollectionsOfCollection ...*LinkedCollections,
) *LinkedCollections {
	isEmpty := len(linkedCollectionsOfCollection) == 0

	if isEmpty && isMakeCloneOnEmpty {
		return NewLinkedCollections().
			AddAnother(it)
	} else if isEmpty && !isMakeCloneOnEmpty {
		return it
	}

	newLinkedCollections := NewLinkedCollections()
	newLinkedCollections.AddAnother(it)

	for _, linkedCollection := range linkedCollectionsOfCollection {
		newLinkedCollections.AddAnother(linkedCollection)
	}

	return newLinkedCollections
}

// AddAsyncFuncItems must add all the lengths to the wg
func (it *LinkedCollections) AddAsyncFuncItems(
	wg *sync.WaitGroup,
	isMakeClone bool,
	asyncFunctions ...func() []string,
) *LinkedCollections {
	if asyncFunctions == nil {
		return it
	}

	asyncFuncWrap := func(asyncFunc func() []string) {
		items := asyncFunc()

		if len(items) == 0 {
			wg.Done()

			return
		}

		collection := NewCollectionUsingStrings(items, isMakeClone)

		it.Lock()
		it.Add(collection)
		it.Unlock()

		wg.Done()
	}

	for _, function := range asyncFunctions {
		go asyncFuncWrap(function)
	}

	wg.Wait()

	return it
}

// AddAsyncFuncItemsPointer must add all the lengths to the wg
func (it *LinkedCollections) AddAsyncFuncItemsPointer(
	wg *sync.WaitGroup,
	isMakeClone bool,
	asyncFunctions ...func() *[]string,
) *LinkedCollections {
	if asyncFunctions == nil {
		return it
	}

	asyncFuncWrap := func(asyncFunc func() *[]string) {
		items := asyncFunc()

		if items == nil || len(*items) == 0 {
			wg.Done()

			return
		}

		collection := NewCollectionUsingStringsPtr(items, isMakeClone)

		it.Lock()
		it.Add(collection)
		it.Unlock()

		wg.Done()
	}

	for _, function := range asyncFunctions {
		go asyncFuncWrap(function)
	}

	wg.Wait()

	return it
}

// AddStringsPtr add to back
func (it *LinkedCollections) AddStringsPtr(
	items *[]string,
	isMakeClone bool,
) *LinkedCollections {
	if items == nil {
		return it
	}

	collection := NewCollectionUsingStringsPtr(items, isMakeClone)

	return it.Add(collection)
}

// AddStringsOfStringsPtr add to back
func (it *LinkedCollections) AddStringsOfStringsPtr(
	items *[]*[]string,
	isMakeClone bool,
) *LinkedCollections {
	if items == nil || len(*items) == 0 {
		return it
	}

	for _, stringItems := range *items {
		if stringItems == nil {
			continue
		}

		it.AddStringsPtr(
			stringItems,
			isMakeClone)
	}

	return it
}

// IndexAt Expensive operation BigO(n)
func (it *LinkedCollections) IndexAt(
	index int,
) *LinkedCollectionNode {
	length := it.Length()
	if index < 0 {
		return nil
	}

	if length == 0 || length-1 < index {
		msgtype.OutOfRange.HandleUsingPanic(
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
func (it *LinkedCollections) SafePointerIndexAt(
	index int,
) *Collection {
	node := it.SafeIndexAt(index)

	if node == nil {
		return nil
	}

	return node.Element
}

// SafeIndexAt Expensive operation BigO(n)
func (it *LinkedCollections) SafeIndexAt(
	index int,
) *LinkedCollectionNode {
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

// AddPointerStringsPtr skip on nil, add to back
func (it *LinkedCollections) AddPointerStringsPtr(
	items *[]*string,
) *LinkedCollections {
	if items == nil {
		return it
	}

	collection := NewCollectionUsingPointerStringsPlusCap(
		items,
		constants.Zero)

	return it.Add(collection)
}

// AddPointerStringsPtrAsync skip on nil, add to back
func (it *LinkedCollections) AddPointerStringsPtrAsync(
	wg *sync.WaitGroup,
	items *[]*string,
) *LinkedCollections {
	if items == nil {
		return it
	}

	go func() {
		collection := NewCollectionUsingPointerStringsPlusCap(
			items,
			constants.Zero)

		it.Lock()
		it.Add(collection)
		it.Unlock()

		wg.Done()
	}()

	return it
}

// AddCollection skip on nil
func (it *LinkedCollections) AddCollection(
	collection *Collection,
) *LinkedCollections {
	if collection == nil {
		return it
	}

	return it.Add(collection)
}

// AddCollectionsPtr skip on nil
func (it *LinkedCollections) AddCollectionsPtr(
	collectionsOfCollection *[]*Collection,
) *LinkedCollections {
	if collectionsOfCollection == nil || len(*collectionsOfCollection) == 0 {
		return it
	}

	return it.AddCollections(*collectionsOfCollection)
}

// AddCollections skip on nil
func (it *LinkedCollections) AddCollections(
	collectionsOfCollection []*Collection,
) *LinkedCollections {
	if len(collectionsOfCollection) == 0 {
		return it
	}

	for _, collection := range collectionsOfCollection {
		if collection == nil {
			continue
		}

		return it.Add(collection)
	}

	return it
}

func (it *LinkedCollections) ToStringsPtr() *[]string {
	return it.ToCollectionSimple().ListPtr()
}

func (it *LinkedCollections) ToStrings() []string {
	return it.ToCollectionSimple().List()
}

func (it *LinkedCollections) ToCollectionSimple() *Collection {
	return it.ToCollection(constants.Zero)
}

func (it *LinkedCollections) ToCollection(
	addCapacity int,
) *Collection {
	if it.IsEmpty() {
		return EmptyCollection()
	}

	newLength := it.AllIndividualItemsLength() +
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

	it.Loop(processor)

	return collection
}

func (it *LinkedCollections) ToCollectionsOfCollection(
	addCapacity int,
) *CollectionsOfCollection {
	if it.IsEmpty() {
		return EmptyCollectionsOfCollection()
	}

	newLength := it.AllIndividualItemsLength() +
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

	it.Loop(processor)

	return collection
}

func (it *LinkedCollections) ItemsOfItems() *[]*[]string {
	length := it.Length()
	itemsOfItems := make([]*[]string, length)

	if length == 0 {
		return &itemsOfItems
	}

	nodes := it.GetAllLinkedNodes()

	for i, node := range *nodes {
		itemsOfItems[i] = &node.Element.items
	}

	return &itemsOfItems
}

func (it *LinkedCollections) ItemsOfItemsCollection() *[]*Collection {
	length := it.Length()
	itemsOfItems := make([]*Collection, length)

	if length == 0 {
		return &itemsOfItems
	}

	nodes := it.GetAllLinkedNodes()

	for i, node := range *nodes {
		itemsOfItems[i] = node.Element
	}

	return &itemsOfItems
}

// ListPtr must return slice.
func (it *LinkedCollections) ListPtr() *[]string {
	return it.
		ToCollection(constants.ArbitraryCapacity5).
		ListPtr()
}

func (it *LinkedCollections) String() string {
	if it.IsEmpty() {
		return commonJoiner + NoElements
	}

	collections := *it.ToCollectionsOfCollection(0)

	return collections.String()
}

func (it *LinkedCollections) StringLock() string {
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

func (it *LinkedCollections) Join(
	separator string,
) string {
	return strings.Join(*it.ListPtr(), separator)
}

func (it *LinkedCollections) Joins(
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

func (it *LinkedCollections) JsonModel() []string {
	return it.ToCollection(0).JsonModel()
}

func (it *LinkedCollections) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it *LinkedCollections) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.JsonModel())
}

func (it *LinkedCollections) UnmarshalJSON(data []byte) error {
	var dataModelStrings []string
	err := json.Unmarshal(data, &dataModelStrings)

	if err == nil {
		it.Clear()
		it.AddStrings(dataModelStrings...)
	}

	return err
}

func (it *LinkedCollections) RemoveAll() *LinkedCollections {
	return it.Clear()
}

func (it *LinkedCollections) Clear() *LinkedCollections {
	if it.IsEmpty() {
		return it
	}

	it.head = nil
	it.tail = nil
	it.setLengthToZero()

	return it
}

func (it LinkedCollections) Json() corejson.Result {
	return corejson.NewFromAny(it)
}

func (it LinkedCollections) JsonPtr() *corejson.Result {
	return corejson.NewFromAnyPtr(it)
}

func (it *LinkedCollections) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*LinkedCollections, error) {
	err := jsonResult.Unmarshal(&it)

	if err != nil {
		return EmptyLinkedCollections(), err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
func (it *LinkedCollections) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *LinkedCollections {
	newUsingJson, err :=
		it.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return newUsingJson
}

func (it *LinkedCollections) GetCompareSummary(
	right *LinkedCollections, leftName, rightName string,
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

func (it *LinkedCollections) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it *LinkedCollections) AsJsoner() corejson.Jsoner {
	return it
}

func (it *LinkedCollections) AsJsonParseSelfInjector() corejson.JsonParseSelfInjector {
	return it
}

func (it *LinkedCollections) AsJsonMarshaller() corejson.JsonMarshaller {
	return it
}
