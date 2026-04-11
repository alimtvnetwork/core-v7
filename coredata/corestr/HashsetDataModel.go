package corestr

import "sync"

type HashsetDataModel struct {
	Items map[string]bool `json:"Hashset"`
}

func NewHashsetUsingDataModel(dataModel *HashsetDataModel) *Hashset {
	return &Hashset{
		items:         dataModel.Items,
		hasMapUpdated: false,
		cachedList:    nil,
		Mutex:         sync.RWMutex{},
	}
}

func NewHashsetsDataModelUsing(collection *Hashset) *HashsetDataModel {
	return &HashsetDataModel{
		Items: collection.items,
	}
}
