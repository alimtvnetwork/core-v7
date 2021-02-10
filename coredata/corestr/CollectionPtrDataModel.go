package corestr

type CollectionPtrDataModel struct {
	Items *[]*string `json:"PointerStringsCollection"`
}

func NewCollectionPtrDataModelUsingDataModel(dataModel *CollectionPtrDataModel) *CollectionPtr {
	return &CollectionPtr{
		items: dataModel.Items,
	}
}

func NewCollectionPtrDataModelUsing(collection *CollectionPtr) *CollectionPtrDataModel {
	return &CollectionPtrDataModel{
		Items: collection.items,
	}
}
