package corestr

type CharHashsetDataModel struct {
	Items               *map[byte]*Hashset `json:"CharacterVsHashsetMap"`
	EachHashsetCapacity int
}

func NewCharHashsetMapUsingDataModel(dataModel *CharHashsetDataModel) *CharHashsetMap {
	return &CharHashsetMap{
		items:               dataModel.Items,
		eachHashsetCapacity: dataModel.EachHashsetCapacity,
	}
}

func NewCharHashsetMapDataModelUsing(dataIn *CharHashsetMap) *CharHashsetDataModel {
	return &CharHashsetDataModel{
		Items:               dataIn.items,
		EachHashsetCapacity: dataIn.eachHashsetCapacity,
	}
}
