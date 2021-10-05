package corestr

type CharCollectionDataModel struct {
	Items                  map[byte]*Collection `json:"CharacterVsStringsCollectionMap"`
	EachCollectionCapacity int
}

func NewCharCollectionMapUsingDataModel(dataModel *CharCollectionDataModel) *CharCollectionMap {
	return &CharCollectionMap{
		items:                  dataModel.Items,
		eachCollectionCapacity: dataModel.EachCollectionCapacity,
	}
}

func NewCharCollectionMapDataModelUsing(dataIn *CharCollectionMap) *CharCollectionDataModel {
	return &CharCollectionDataModel{
		Items:                  dataIn.items,
		EachCollectionCapacity: dataIn.eachCollectionCapacity,
	}
}
