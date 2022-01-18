package corepayload

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/coredynamic"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/coredata/stringslice"
	"gitlab.com/evatix-go/core/errcore"
)

type newAttributesCreator struct{}

func (it *newAttributesCreator) Create(
	err error,
	dynamicPayloads []byte,
) *Attributes {
	return &Attributes{
		ErrorMessage:     errcore.ToString(err),
		KeyValuePairs:    corestr.Empty.Hashmap(),
		AnyKeyValuePairs: coredynamic.EmptyMapAnyItems(),
		DynamicPayloads:  dynamicPayloads,
	}
}

func (it *newAttributesCreator) UsingErrMsg(
	errMsg string,
) *Attributes {
	if errMsg == "" {
		return it.Empty()
	}

	return &Attributes{
		ErrorMessage:     errMsg,
		KeyValuePairs:    corestr.Empty.Hashmap(),
		AnyKeyValuePairs: coredynamic.EmptyMapAnyItems(),
		DynamicPayloads:  []byte{},
	}
}

func (it *newAttributesCreator) UsingErrMessages(
	errMessages ...string,
) *Attributes {
	return &Attributes{
		ErrorMessage: strings.Join(
			stringslice.NonEmptySlice(errMessages),
			constants.DefaultLine),
		KeyValuePairs:    corestr.Empty.Hashmap(),
		AnyKeyValuePairs: coredynamic.EmptyMapAnyItems(),
		DynamicPayloads:  []byte{},
	}
}

func (it *newAttributesCreator) UsingErr(
	err error,
) *Attributes {
	return &Attributes{
		ErrorMessage:     errcore.ToString(err),
		KeyValuePairs:    corestr.Empty.Hashmap(),
		AnyKeyValuePairs: coredynamic.EmptyMapAnyItems(),
		DynamicPayloads:  []byte{},
	}
}

func (it *newAttributesCreator) UsingDynamicPayloadBytes(
	dynamicPayloads []byte,
) *Attributes {
	return &Attributes{
		KeyValuePairs:    corestr.Empty.Hashmap(),
		AnyKeyValuePairs: coredynamic.EmptyMapAnyItems(),
		DynamicPayloads:  dynamicPayloads,
	}
}

func (it *newAttributesCreator) AllAny(
	keyValues *corestr.Hashmap,
	anyKeyValues *coredynamic.MapAnyItems,
	pagingInfo *PagingInfo,
	anyItem interface{},
) *Attributes {
	jsonResult := corejson.
		Serialize.
		UsingAny(anyItem)

	return &Attributes{
		ErrorMessage:     jsonResult.MeaningfulErrorMessage(),
		PagingInfo:       pagingInfo,
		KeyValuePairs:    keyValues,
		AnyKeyValuePairs: anyKeyValues,
		DynamicPayloads:  jsonResult.SafeBytes(),
	}
}

func (it *newAttributesCreator) PageInfoAny(
	pagingInfo *PagingInfo,
	anyItem interface{},
) *Attributes {
	jsonResult := corejson.
		Serialize.
		UsingAny(anyItem)

	return &Attributes{
		ErrorMessage:     jsonResult.MeaningfulErrorMessage(),
		PagingInfo:       pagingInfo,
		KeyValuePairs:    corestr.Empty.Hashmap(),
		AnyKeyValuePairs: coredynamic.EmptyMapAnyItems(),
		DynamicPayloads:  jsonResult.SafeBytes(),
	}
}

func (it *newAttributesCreator) All(
	keyValues *corestr.Hashmap,
	anyKeyValues *coredynamic.MapAnyItems,
	pagingInfo *PagingInfo,
	dynamicPayloads []byte,
	err error,
) *Attributes {
	return &Attributes{
		ErrorMessage:     errcore.ToString(err),
		PagingInfo:       pagingInfo,
		KeyValuePairs:    keyValues,
		AnyKeyValuePairs: anyKeyValues,
		DynamicPayloads:  dynamicPayloads,
	}
}

func (it *newAttributesCreator) UsingDynamicAny(
	anyItem interface{},
) *Attributes {
	jsonResult := corejson.
		Serialize.
		UsingAny(anyItem)

	return &Attributes{
		ErrorMessage:     jsonResult.MeaningfulErrorMessage(),
		KeyValuePairs:    corestr.Empty.Hashmap(),
		AnyKeyValuePairs: coredynamic.EmptyMapAnyItems(),
		DynamicPayloads:  jsonResult.SafeBytes(),
	}
}

func (it *newAttributesCreator) UsingKeyValues(
	keyValues *corestr.Hashmap,
) *Attributes {
	return &Attributes{
		KeyValuePairs:    keyValues,
		AnyKeyValuePairs: coredynamic.EmptyMapAnyItems(),
		DynamicPayloads:  []byte{},
	}
}

func (it *newAttributesCreator) UsingKeyValuesPlusDynamic(
	keyValues *corestr.Hashmap,
	dynamicPayloads []byte,
) *Attributes {
	return &Attributes{
		KeyValuePairs:    keyValues,
		AnyKeyValuePairs: coredynamic.EmptyMapAnyItems(),
		DynamicPayloads:  dynamicPayloads,
	}
}

func (it *newAttributesCreator) UsingAnyKeyValues(
	anyKeyValues *coredynamic.MapAnyItems,
) *Attributes {
	return it.UsingAnyKeyValuesPlusDynamic(
		anyKeyValues,
		[]byte{})
}

func (it *newAttributesCreator) UsingAnyKeyValuesPlusDynamic(
	anyKeyValues *coredynamic.MapAnyItems,
	dynamicPayloads []byte,
) *Attributes {
	return &Attributes{
		KeyValuePairs:    corestr.Empty.Hashmap(),
		AnyKeyValuePairs: anyKeyValues,
		DynamicPayloads:  dynamicPayloads,
	}
}

func (it *newAttributesCreator) UsingErrors(
	errorItems ...error,
) *Attributes {
	return &Attributes{
		ErrorMessage:     errcore.MergeErrorsToStringDefault(errorItems...),
		KeyValuePairs:    corestr.Empty.Hashmap(),
		AnyKeyValuePairs: coredynamic.EmptyMapAnyItems(),
		DynamicPayloads:  []byte{},
	}
}

func (it *newAttributesCreator) Empty() *Attributes {
	return &Attributes{
		KeyValuePairs:    corestr.Empty.Hashmap(),
		AnyKeyValuePairs: coredynamic.EmptyMapAnyItems(),
		DynamicPayloads:  []byte{},
	}
}
