package corestr

type CollectionDataModel struct {
	Items *[]string `json:"StringsCollection"`
}

func NewCollectionDataModelUsingDataModel(dataModel *CollectionDataModel) *Collection {
	return &Collection{
		items: dataModel.Items,
	}
}

func NewCollectionDataModelUsing(collection *Collection) *CollectionDataModel {
	return &CollectionDataModel{
		Items: collection.items,
	}
}
