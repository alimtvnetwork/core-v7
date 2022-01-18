package corestr

type newCreator struct {
	Collection                 newCollectionCreator
	CollectionPtr              newCollectionPtrCreator
	CharHashsetMap             newCharHashsetMapCreator
	CharCollectionMap          newCharCollectionMapCreator
	SimpleStringOnce           newSimpleStringOnceCreator
	SimpleSlice                newSimpleSliceCreator
	Hashset                    newHashsetCreator
	KeyValues                  newKeyValuesCreator
	HashsetsCollection         newHashsetsCollectionCreator
	Hashmap                    newHashmapCreator
	LinkedList                 newLinkedListCreator
	LinkedCollection           newLinkedListCollectionsCreator
	CollectionsOfCollection    newCollectionsOfCollectionCreator
	CollectionsOfCollectionPtr newCollectionsOfCollectionPtrCreator
}
