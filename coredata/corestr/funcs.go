package corestr

type ReturningBool struct {
	IsBreak, IsKeep bool
}

type LinkedCollectionFilterResult struct {
	Value           *LinkedCollectionNode
	IsKeep, IsBreak bool
}

type LinkedListFilterResult struct {
	Value           *LinkedListNode
	IsKeep, IsBreak bool
}

type LinkedCollectionFilterParameter struct {
	Node  *LinkedCollectionNode
	Index int
}

type LinkedListFilterParameter struct {
	Node  *LinkedListNode
	Index int
}

type LinkedListProcessorParameter struct {
	Index                       int
	CurrentNode, PrevNode       *LinkedListNode
	IsFirstIndex, IsEndingIndex bool
}

type LinkedCollectionProcessorParameter struct {
	Index                       int
	CurrentNode, PrevNode       *LinkedCollectionNode
	IsFirstIndex, IsEndingIndex bool
}

type OnCompleteCharCollectionMap func(charCollection *CharCollectionMap)
type OnCompleteLinkedCollections func(linkedCollections *LinkedCollections)
type AnyToCollectionProcessor func(any interface{}) *Collection
type OnCompleteCharHashsetMap func(charHashset *CharHashsetMap)
type IsStringFilter func(str string) (result string, isKeep bool, isBreak bool)
type IsKeyAnyValueFilter func(pair KeyAnyValuePair) (result string, isKeep bool, isBreak bool)
type IsKeyValueFilter func(pair KeyValuePair) (result string, isKeep bool, isBreak bool)
type IsStringPointerFilter func(stringPointer *string) (result *string, isKeep bool, isBreak bool)
type LinkedListFilter func(arg *LinkedListFilterParameter) *LinkedListFilterResult
type LinkedListSimpleProcessor func(
	arg *LinkedListProcessorParameter,
) (isBreak bool)
type LinkedCollectionFilter func(
	arg *LinkedCollectionFilterParameter,
) *LinkedCollectionFilterResult
type LinkedCollectionSimpleProcessor func(
	arg *LinkedCollectionProcessorParameter,
) (isBreak bool)
