package corestr

type HashsetsCollectionDataModel struct {
	Items *[]*Hashset `json:"HashsetsCollections"`
}

func NewHashsetsCollectionUsingDataModel(dataModel *HashsetsCollectionDataModel) *HashsetsCollection {
	return &HashsetsCollection{
		items: dataModel.Items,
	}
}

func NewHashsetsCollectionDataModelUsing(collection *HashsetsCollection) *HashsetsCollectionDataModel {
	return &HashsetsCollectionDataModel{
		Items: collection.items,
	}
}
