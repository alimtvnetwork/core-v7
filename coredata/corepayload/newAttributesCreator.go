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

func (it *newAttributesCreator) Deserialize(
	rawBytes []byte,
) (*Attributes, error) {
	empty := &Attributes{}
	err := corejson.
		Deserialize.
		UsingBytes(rawBytes, empty)

	if err == nil {
		return empty, nil
	}

	// has error
	return nil, err
}

func (it *newAttributesCreator) DeserializeMany(
	rawBytes []byte,
) (attrSlice []*Attributes, err error) {
	err = corejson.
		Deserialize.
		UsingBytes(rawBytes, &attrSlice)

	if err == nil {
		return attrSlice, nil
	}

	// has error
	return nil, err
}

func (it *newAttributesCreator) DeserializeUsingJsonResult(
	jsonResult *corejson.Result,
) (*Attributes, error) {
	empty := &Attributes{}
	err := corejson.
		Deserialize.
		UsingResult(jsonResult, empty)

	if err == nil {
		return empty, nil
	}

	// has error
	return nil, err
}

func (it *newAttributesCreator) Create(
	err error,
	authInfo *AuthInfo,
	dynamicPayloads []byte,
) *Attributes {
	return &Attributes{
		ErrorMessage:     errcore.ToString(err),
		AuthInfo:         authInfo,
		KeyValuePairs:    corestr.Empty.Hashmap(),
		AnyKeyValuePairs: coredynamic.EmptyMapAnyItems(),
		DynamicPayloads:  dynamicPayloads,
	}
}

func (it *newAttributesCreator) UsingAuthInfoJsonResult(
	authInfo *AuthInfo,
	jsonResult *corejson.Result,
) *Attributes {
	return &Attributes{
		ErrorMessage:     jsonResult.MeaningfulErrorMessage(),
		AuthInfo:         authInfo,
		KeyValuePairs:    corestr.Empty.Hashmap(),
		AnyKeyValuePairs: coredynamic.EmptyMapAnyItems(),
		DynamicPayloads:  jsonResult.Bytes,
	}
}

func (it *newAttributesCreator) UsingAuthInfoDynamicBytes(
	authInfo *AuthInfo,
	dynamicPayloads []byte,
) *Attributes {
	return &Attributes{
		AuthInfo:         authInfo,
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
	authInfo *AuthInfo,
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
		AuthInfo:         authInfo,
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
	authInfo *AuthInfo,
	keyValues *corestr.Hashmap,
	anyKeyValues *coredynamic.MapAnyItems,
	pagingInfo *PagingInfo,
	dynamicPayloads []byte,
	err error,
) *Attributes {
	return &Attributes{
		ErrorMessage:     errcore.ToString(err),
		AuthInfo:         authInfo,
		PagingInfo:       pagingInfo,
		KeyValuePairs:    keyValues,
		AnyKeyValuePairs: anyKeyValues,
		DynamicPayloads:  dynamicPayloads,
	}
}

func (it *newAttributesCreator) UsingAuthInfo(
	authInfo *AuthInfo,
) *Attributes {
	return &Attributes{
		AuthInfo:         authInfo,
		KeyValuePairs:    corestr.Empty.Hashmap(),
		AnyKeyValuePairs: coredynamic.EmptyMapAnyItems(),
	}
}

func (it *newAttributesCreator) UsingDynamicPayloadAny(
	authInfo *AuthInfo,
	anyItem interface{},
) *Attributes {
	jsonResult := corejson.
		Serialize.
		UsingAny(anyItem)

	return &Attributes{
		ErrorMessage:     jsonResult.MeaningfulErrorMessage(),
		AuthInfo:         authInfo,
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

func (it *newAttributesCreator) UsingAuthInfoKeyValues(
	authInfo *AuthInfo,
	keyValues *corestr.Hashmap,
) *Attributes {
	return &Attributes{
		AuthInfo:         authInfo,
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

func (it *newAttributesCreator) UsingAuthInfoAnyKeyValues(
	authInfo *AuthInfo,
	anyKeyValues *coredynamic.MapAnyItems,
) *Attributes {
	return &Attributes{
		AuthInfo:         authInfo,
		KeyValuePairs:    corestr.Empty.Hashmap(),
		AnyKeyValuePairs: anyKeyValues,
		DynamicPayloads:  []byte{},
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
