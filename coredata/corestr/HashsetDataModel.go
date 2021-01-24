package corestr

import "sync"

type HashsetDataModel struct {
	Items *map[string]bool `json:"Hashset"`
}

func NewHashsetUsingDataModel(dataModel *HashsetDataModel) *Hashset {
	length := 0

	if dataModel.Items != nil {
		length = len(*dataModel.Items)
	}

	return &Hashset{
		items:         dataModel.Items,
		hasMapUpdated: false,
		cachedList:    nil,
		length:        length,
		isEmptySet:    length == 0,
		Mutex:         sync.Mutex{},
	}
}

func NewHashsetsDataModelUsing(collection *Hashset) *HashsetDataModel {
	return &HashsetDataModel{
		Items: collection.items,
	}
}
