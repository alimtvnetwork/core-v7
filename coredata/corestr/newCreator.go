package corestr

type newCreator struct {
	Collection                 *newCollectionCreator
	CollectionPtr              *newCollectionPtrCreator
	CharHashsetMap             *newCharHashsetMapCreator
	CharCollectionMap          *newCharCollectionMapCreator
	Hashset                    *newHashsetCreator
	HashsetsCollection         *newHashsetsCollectionCreator
	Hashmap                    *newHashmapCreator
	LinkedList                 *newLinkedListCreator
	LinkedCollection           *newLinkedListCollectionsCreator
	CollectionsOfCollection    *newCollectionsOfCollectionCreator
	CollectionsOfCollectionPtr *newCollectionsOfCollectionPtrCreator
}
