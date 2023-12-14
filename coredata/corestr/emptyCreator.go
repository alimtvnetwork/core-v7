package corestr

type emptyCreator struct{}

func (it *emptyCreator) Collection() *Collection {
	return &Collection{
		items: []string{},
	}
}

func (it *emptyCreator) LinkedList() *LinkedList {
	return &LinkedList{}
}

func (it *emptyCreator) SimpleSlice() *SimpleSlice {
	return New.SimpleSlice.Empty()
}

func (it *emptyCreator) KeyAnyValuePair() *KeyAnyValuePair {
	return &KeyAnyValuePair{}
}

func (it *emptyCreator) KeyValuePair() *KeyValuePair {
	return &KeyValuePair{}
}

func (it *emptyCreator) KeyValueCollection() *KeyValueCollection {
	return &KeyValueCollection{
		KeyValuePairs: []KeyValuePair{},
	}
}

func (it *emptyCreator) LinkedCollections() *LinkedCollections {
	return &LinkedCollections{}
}

func (it *emptyCreator) LeftRight() *LeftRight {
	return &LeftRight{}
}

func (it *emptyCreator) SimpleStringOnce() SimpleStringOnce {
	return SimpleStringOnce{}
}

func (it *emptyCreator) SimpleStringOncePtr() *SimpleStringOnce {
	return &SimpleStringOnce{}
}

func (it *emptyCreator) Hashset() *Hashset {
	return &Hashset{
		hasMapUpdated: false,
		isEmptySet:    true,
		length:        0,
		items:         map[string]bool{},
	}
}

func (it *emptyCreator) HashsetsCollection() *HashsetsCollection {
	return &HashsetsCollection{
		items: nil,
	}
}

func (it *emptyCreator) Hashmap() *Hashmap {
	return &Hashmap{
		isEmptySet: true,
		items:      map[string]string{},
	}
}

func (it *emptyCreator) CharCollectionMap() *CharCollectionMap {
	return &CharCollectionMap{
		items: nil,
	}
}

func (it *emptyCreator) KeyValuesCollection() *KeyValueCollection {
	return &KeyValueCollection{
		KeyValuePairs: nil,
	}
}

func (it *emptyCreator) CollectionsOfCollection() *CollectionsOfCollection {
	return &CollectionsOfCollection{
		items: nil,
	}
}

func (it *emptyCreator) CharHashsetMap() *CharHashsetMap {
	return &CharHashsetMap{
		items: nil,
	}
}
