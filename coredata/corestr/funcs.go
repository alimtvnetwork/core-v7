package corestr

type OnCompleteCharCollectionMap func(charCollection *CharCollectionMap)
type OnCompleteCharHashsetMap func(charHashset *CharHashsetMap)
type IsStringFilter func(str string) (result string, isKeep bool)
type IsKeyAnyValueFilter func(pair KeyAnyValuePair) (result string, isKeep bool)
type IsKeyValueFilter func(pair KeyValuePair) (result string, isKeep bool)
type IsStringPointerFilter func(stringPointer *string) (result *string, isKeep bool)
type LinkedListFilter func(list *LinkedList, index int, node *LinkedListNode) (result *LinkedListNode, isKeep bool)
type LinkedListSimpleProcessor func(
	index int, currentNode, prevNode *LinkedListNode, isFirstIndex, isEndingIndex bool,
) (isBreak bool)
type LinkedCollectionFilter func(
	list *LinkedCollections, index int, node *LinkedCollectionNode,
) (result *LinkedCollectionNode, isKeep bool)
type LinkedCollectionSimpleProcessor func(
	index int, currentNode, prevNode *LinkedCollectionNode, isFirstIndex, isEndingIndex bool,
) (isBreak bool)

type LinkedCollectionReverseProcessor func(
	index int,
	currentNode, nextNode *LinkedCollectionNode,
	isFirstIndex,
	isEndingIndex bool,
) (isBreak bool)
