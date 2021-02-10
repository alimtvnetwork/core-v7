package corestr

type HashsetsCollectionDataModel struct {
	Items *[]*Hashset `json:"HashsetsCollections"`
}

//goland:noinspection GoUnusedExportedFunction
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
