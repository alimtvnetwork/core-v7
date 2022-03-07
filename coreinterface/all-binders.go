package coreinterface

type BasicSlicerContractsBinder interface {
	BasicSlicer
	AsBasicSliceContractsBinder() BasicSlicerContractsBinder
}

type StandardSlicerContractsBinder interface {
	StandardSlicer
	AsStandardSlicerContractsBinder() StandardSlicerContractsBinder
}

type DynamicLinqWithPagingContractsBinder interface {
	DynamicLinqWithPaging
	AsDynamicLinqWithPagingContractsBinder() DynamicLinqWithPagingContractsBinder
}

type DynamicLinqContractsBinder interface {
	DynamicLinq
	AsDynamicLinqContractsBinder() DynamicLinqContractsBinder
}

type SimpleValidInvalidCheckerContractsBinder interface {
	SimpleValidInvalidChecker
	AsSimpleValidInvalidChecker() SimpleValidInvalidChecker
}

type JsonBytesStringerContractsBinder interface {
	JsonByter
	JsonCombineStringer
	AsJsonBytesStringerContractsBinder() JsonBytesStringerContractsBinder
}

type CountStateTrackerBinder interface {
	CountStateTracker
	AsCountStateTrackerBinder() CountStateTrackerBinder
}

type KeyValueStringDefinerBinder interface {
	KeyValueStringDefiner
	AsKeyValueStringDefinerBinder() KeyValueStringDefinerBinder
}

type KeyAnyValueDefinerBinder interface {
	KeyAnyValueDefiner
	AsKeyAnyValueDefinerBinder() KeyAnyValueDefinerBinder
}

type KeyStringValuesCollectionDefinerBinder interface {
	KeyStringValuesCollectionDefiner
	AsKeyStringValuesCollectionDefinerBinder() KeyStringValuesCollectionDefinerBinder
}

type KeyAnyValuesCollectionDefinerBinder interface {
	KeyAnyValuesCollectionDefiner
	AsKeyAnyValuesCollectionDefinerBinder() KeyAnyValuesCollectionDefinerBinder
}
