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

// including all nested ones
func (linkedCollections *LinkedCollections) AllIndividualItemsLength() int {
	allLengthSum := 0

	var processor LinkedCollectionSimpleProcessor = func(
		index int,
		currentNode,
		prevNode *LinkedCollectionNode,
		isFirstIndex,
		isEndingIndex bool,
	) (isBreak bool) {
		allLengthSum += currentNode.Element.Length()

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

// BigO(n) expensive operation.
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

func (linkedCollections *LinkedCollections) PushBack(collection *Collection) *LinkedCollections {
	return linkedCollections.Add(collection)
}

func (linkedCollections *LinkedCollections) Push(collection *Collection) *LinkedCollections {
	return linkedCollections.Add(collection)
}

func (linkedCollections *LinkedCollections) PushFront(collection *Collection) *LinkedCollections {
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

// iSkipOnNil
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

func (linkedCollections *LinkedCollections) Loop(
	simpleProcessor LinkedCollectionSimpleProcessor,
) *LinkedCollections {
	length := linkedCollections.Length()
	if length == 0 {
		return linkedCollections
	}

	node := linkedCollections.head
	isBreak := simpleProcessor(
		0,
		node,
		nil,
		true,
		false)

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
		isBreak = simpleProcessor(
			index,
			node,
			prev,
			false,
			isEndingIndex)

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
	result, isKeep := filter(
		linkedCollections,
		0,
		node)

	if isKeep {
		list = append(list, result)
	}

	index := 1
	for node.HasNext() {
		node = node.Next()
		result2, isKeep2 := filter(
			linkedCollections,
			index,
			node)

		if isKeep2 {
			list = append(list, result2)
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
	removingIndexes ...int,
) *LinkedCollections {
	if removingIndexes == nil {
		return linkedCollections
	}

	removingIndexesCopy := removingIndexes
	removingIndexesCopyPtr := &removingIndexesCopy

	var processor LinkedCollectionSimpleProcessor = func(
		index int,
		currentNode, prevNode *LinkedCollectionNode,
		isFirstIndex,
		isEndingIndex bool,
	) (isBreak bool) {
		hasIndex := coreindexes.HasIndexPlusRemoveIndex(removingIndexesCopyPtr, index)

		if !hasIndex {
			return isBreak
		}

		isBreak = len(*removingIndexesCopyPtr) == 0
		linkedCollections.decrementLength()

		if isFirstIndex {
			linkedCollections.head = currentNode.next
			currentNode = nil
			return isBreak
		}

		if isEndingIndex {
			prevNode.next = nil
			currentNode = nil

			return isBreak
		}

		prevNode.next = currentNode.next
		currentNode = nil

		return isBreak
	}

	return linkedCollections.Loop(processor)
}

func (linkedCollections *LinkedCollections) RemoveNode(
	removingNode *LinkedCollectionNode,
) *LinkedCollections {
	var processor LinkedCollectionSimpleProcessor = func(
		index int,
		currentNode,
		prevNode *LinkedCollectionNode,
		isFirstIndex,
		isEndingIndex bool,
	) (isBreak bool) {
		isSameNode := currentNode == removingNode
		if isSameNode && isFirstIndex {
			linkedCollections.head = currentNode.next
			linkedCollections.decrementLength()

			return true
		}

		if isSameNode {
			prevNode.next = currentNode.next
			linkedCollections.decrementLength()

			return true
		}

		return false
	}

	return linkedCollections.Loop(processor)
}

// iSkipOnNil
func (linkedCollections *LinkedCollections) AppendCollections(
	isSkipOnNull bool,
	collections ...*Collection,
) *LinkedCollections {
	if isSkipOnNull && collections == nil {
		return linkedCollections
	}

	for i := range collections {
		linkedCollections.Add(collections[i])
	}

	return linkedCollections
}

// iSkipOnNil
func (linkedCollections *LinkedCollections) AppendCollectionsPointers(
	isSkipOnNull bool,
	collections *[]*Collection,
) *LinkedCollections {
	if isSkipOnNull && collections == nil {
		return linkedCollections
	}

	for i := range *collections {
		linkedCollections.Add((*collections)[i])
	}

	return linkedCollections
}

// iSkipOnNil
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

// iSkipOnNil
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

// iSkipOnNil
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

// add to back
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

// add to back
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

// Expensive operation BigO(n)
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

// Expensive operation BigO(n)
func (linkedCollections *LinkedCollections) SafePointerIndexAt(
	index int,
) *Collection {
	node := linkedCollections.SafeIndexAt(index)

	if node == nil {
		return nil
	}

	return node.Element
}

// Expensive operation BigO(n)
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

// skip on nil, add to back
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

// skip on nil, add to back
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

// skip on nil
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

	collection := NewCollection(
		linkedCollections.AllIndividualItemsLength() +
			addCapacity)

	var processor LinkedCollectionSimpleProcessor = func(
		index int,
		currentNode,
		prevNode *LinkedCollectionNode,
		isFirstIndex,
		isEndingIndex bool,
	) (isBreak bool) {
		if currentNode == nil {
			return false
		}

		collection.AddCollection(currentNode.Element)

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

	collection := NewCollectionsOfCollection(
		linkedCollections.AllIndividualItemsLength() +
			addCapacity)

	var processor LinkedCollectionSimpleProcessor = func(
		index int,
		currentNode,
		prevNode *LinkedCollectionNode,
		isFirstIndex,
		isEndingIndex bool,
	) (isBreak bool) {
		if currentNode == nil {
			return false
		}

		collection.Adds(currentNode.Element)

		return false
	}

	linkedCollections.Loop(processor)

	return collection
}

// must return slice.
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
		return EmptyLinkedCollections(), nil
	}

	err := json.Unmarshal(*jsonResult.Bytes, &linkedCollections)

	if err != nil {
		return EmptyLinkedCollections(), err
	}

	return linkedCollections, nil
}

// Panic if error
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

// Panic if error
func (linkedCollections *LinkedCollections) JsonParseSelfInject(
	jsonResult *corejson.Result,
) {
	linkedCollections.ParseInjectUsingJsonMust(jsonResult)
}

func (linkedCollections *LinkedCollections) AsJsoner() *corejson.Jsoner {
	var jsoner corejson.Jsoner = linkedCollections

	return &jsoner
}

func (linkedCollections *LinkedCollections) AsJsonParseSelfInjector() *corejson.ParseSelfInjector {
	var jsonInjector corejson.ParseSelfInjector = linkedCollections

	return &jsonInjector
}

func (linkedCollections *LinkedCollections) AsJsonMarshaller() *corejson.JsonMarshaller {
	var jsonMarshaller corejson.JsonMarshaller = linkedCollections

	return &jsonMarshaller
}
