package corestr

type newCreator struct {
	Collection              newCollectionCreator
	CharHashsetMap          newCharHashsetMapCreator
	CharCollectionMap       newCharCollectionMapCreator
	SimpleStringOnce        newSimpleStringOnceCreator
	SimpleSlice             newSimpleSliceCreator
	Hashset                 newHashsetCreator
	KeyValues               newKeyValuesCreator
	HashsetsCollection      newHashsetsCollectionCreator
	Hashmap                 newHashmapCreator
	LinkedList              newLinkedListCreator
	LinkedCollection        newLinkedListCollectionsCreator
	CollectionsOfCollection newCollectionsOfCollectionCreator
}
