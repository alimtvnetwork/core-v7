package corestr

import "sync"

type HashmapDataModel struct {
	Items map[string]string `json:"Hashmap"`
}

func NewHashmapUsingDataModel(dataModel *HashmapDataModel) *Hashmap {
	return &Hashmap{
		items:         dataModel.Items,
		hasMapUpdated: false,
		cachedList:    nil,
		Mutex:         sync.RWMutex{},
	}
}

func NewHashmapsDataModelUsing(collection *Hashmap) *HashmapDataModel {
	return &HashmapDataModel{
		Items: collection.items,
	}
}
