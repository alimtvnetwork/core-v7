package corestr

import "sync"

type HashmapDataModel struct {
	Items *map[string]string `json:"Hashmap"`
}

func NewHashmapUsingDataModel(dataModel *HashmapDataModel) *Hashmap {
	length := 0

	if dataModel.Items != nil {
		length = len(*dataModel.Items)
	}

	return &Hashmap{
		items:         dataModel.Items,
		hasMapUpdated: false,
		cachedList:    nil,
		length:        length,
		isEmptySet:    length == 0,
		Mutex:         sync.Mutex{},
	}
}

func NewHashmapsDataModelUsing(collection *Hashmap) *HashmapDataModel {
	return &HashmapDataModel{
		Items: collection.items,
	}
}
