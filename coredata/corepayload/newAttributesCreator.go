package corepayload

import (
	"gitlab.com/evatix-go/core/coredata/coredynamic"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/coreinstruction"
	"gitlab.com/evatix-go/core/coreinterface/errcoreinf"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/core/isany"
)

type newAttributesCreator struct{}

func (it newAttributesCreator) CastOrDeserializeFrom(
	anyItem interface{},
) (*Attributes, error) {
	if isany.Null(anyItem) {
		return nil, errcore.
			CannotBeNilOrEmptyType.
			ErrorNoRefs(
				"given any item is nil failed to convert to attributes")
	}

	toAttributes := &Attributes{}
	err := corejson.CastAny.FromToDefault(
		anyItem,
		toAttributes)

	return toAttributes, err
}

func (it newAttributesCreator) Deserialize(
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

func (it newAttributesCreator) DeserializeMany(
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

func (it newAttributesCreator) DeserializeUsingJsonResult(
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

func (it newAttributesCreator) Create(
	basicErrWrapper errcoreinf.BasicErrWrapper,
	authInfo *AuthInfo,
	dynamicPayloads []byte,
) *Attributes {
	return &Attributes{
		BasicErrWrapper:  basicErrWrapper,
		AuthInfo:         authInfo,
		KeyValuePairs:    corestr.Empty.Hashmap(),
		AnyKeyValuePairs: coredynamic.EmptyMapAnyItems(),
		DynamicPayloads:  dynamicPayloads,
	}
}

func (it newAttributesCreator) ErrFromTo(
	basicErrWrapper errcoreinf.BasicErrWrapper,
	fromTo *coreinstruction.FromTo,
	dynamicPayloads []byte,
) *Attributes {
	return &Attributes{
		BasicErrWrapper:  basicErrWrapper,
		KeyValuePairs:    corestr.Empty.Hashmap(),
		AnyKeyValuePairs: coredynamic.EmptyMapAnyItems(),
		FromTo:           fromTo,
		DynamicPayloads:  dynamicPayloads,
	}
}

func (it newAttributesCreator) UsingAuthInfoJsonResult(
	authInfo *AuthInfo,
	jsonResult *corejson.Result,
) (*Attributes, error) {
	return &Attributes{
			AuthInfo:         authInfo,
			KeyValuePairs:    corestr.Empty.Hashmap(),
			AnyKeyValuePairs: coredynamic.EmptyMapAnyItems(),
			DynamicPayloads:  jsonResult.Bytes,
		},
		jsonResult.MeaningfulError()
}

func (it newAttributesCreator) UsingAuthInfoDynamicBytes(
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

func (it newAttributesCreator) UsingDynamicPayloadBytes(
	dynamicPayloads []byte,
) *Attributes {
	return &Attributes{
		KeyValuePairs:    corestr.Empty.Hashmap(),
		AnyKeyValuePairs: coredynamic.EmptyMapAnyItems(),
		DynamicPayloads:  dynamicPayloads,
	}
}

func (it newAttributesCreator) AllAny(
	authInfo *AuthInfo,
	keyValues *corestr.Hashmap,
	anyKeyValues *coredynamic.MapAnyItems,
	pagingInfo *PagingInfo,
	anyItem interface{},
) (*Attributes, error) {
	jsonResult := corejson.
		Serialize.
		UsingAny(anyItem)

	return &Attributes{
		AuthInfo:         authInfo,
		PagingInfo:       pagingInfo,
		KeyValuePairs:    keyValues,
		AnyKeyValuePairs: anyKeyValues,
		DynamicPayloads:  jsonResult.SafeBytes(),
	}, jsonResult.MeaningfulError()
}

func (it newAttributesCreator) PageInfoAny(
	pagingInfo *PagingInfo,
	anyItem interface{},
) (*Attributes, error) {
	jsonResult := corejson.
		Serialize.
		UsingAny(anyItem)

	return &Attributes{
		PagingInfo:       pagingInfo,
		KeyValuePairs:    corestr.Empty.Hashmap(),
		AnyKeyValuePairs: coredynamic.EmptyMapAnyItems(),
		DynamicPayloads:  jsonResult.SafeBytes(),
	}, jsonResult.MeaningfulError()
}

func (it newAttributesCreator) All(
	authInfo *AuthInfo,
	keyValues *corestr.Hashmap,
	anyKeyValues *coredynamic.MapAnyItems,
	pagingInfo *PagingInfo,
	dynamicPayloads []byte,
	fromTo *coreinstruction.FromTo,
	basicErr errcoreinf.BasicErrWrapper,
) *Attributes {
	return &Attributes{
		BasicErrWrapper:  basicErr,
		AuthInfo:         authInfo,
		PagingInfo:       pagingInfo,
		KeyValuePairs:    keyValues,
		AnyKeyValuePairs: anyKeyValues,
		FromTo:           fromTo,
		DynamicPayloads:  dynamicPayloads,
	}
}

func (it newAttributesCreator) UsingAuthInfo(
	authInfo *AuthInfo,
) *Attributes {
	return &Attributes{
		AuthInfo:         authInfo,
		KeyValuePairs:    corestr.Empty.Hashmap(),
		AnyKeyValuePairs: coredynamic.EmptyMapAnyItems(),
	}
}

func (it newAttributesCreator) UsingDynamicPayloadAny(
	authInfo *AuthInfo,
	anyItem interface{},
) (*Attributes, error) {
	jsonResult := corejson.
		Serialize.
		UsingAny(anyItem)

	attr := &Attributes{
		AuthInfo:         authInfo,
		KeyValuePairs:    corestr.Empty.Hashmap(),
		AnyKeyValuePairs: coredynamic.EmptyMapAnyItems(),
		DynamicPayloads:  jsonResult.SafeBytes(),
	}

	return attr, jsonResult.MeaningfulError()
}

func (it newAttributesCreator) UsingKeyValues(
	keyValues *corestr.Hashmap,
) *Attributes {
	return &Attributes{
		KeyValuePairs:    keyValues,
		AnyKeyValuePairs: coredynamic.EmptyMapAnyItems(),
		DynamicPayloads:  []byte{},
	}
}

func (it newAttributesCreator) UsingAuthInfoKeyValues(
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

func (it newAttributesCreator) UsingKeyValuesPlusDynamic(
	keyValues *corestr.Hashmap,
	dynamicPayloads []byte,
) *Attributes {
	return &Attributes{
		KeyValuePairs:    keyValues,
		AnyKeyValuePairs: coredynamic.EmptyMapAnyItems(),
		DynamicPayloads:  dynamicPayloads,
	}
}

func (it newAttributesCreator) UsingAuthInfoAnyKeyValues(
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

func (it newAttributesCreator) UsingAnyKeyValues(
	anyKeyValues *coredynamic.MapAnyItems,
) *Attributes {
	return it.UsingAnyKeyValuesPlusDynamic(
		anyKeyValues,
		[]byte{})
}

func (it newAttributesCreator) UsingAnyKeyValuesPlusDynamic(
	anyKeyValues *coredynamic.MapAnyItems,
	dynamicPayloads []byte,
) *Attributes {
	return &Attributes{
		KeyValuePairs:    corestr.Empty.Hashmap(),
		AnyKeyValuePairs: anyKeyValues,
		DynamicPayloads:  dynamicPayloads,
	}
}

func (it newAttributesCreator) UsingBasicError(
	basicErrWrapper errcoreinf.BasicErrWrapper,
) *Attributes {
	return &Attributes{
		BasicErrWrapper:  basicErrWrapper,
		KeyValuePairs:    corestr.Empty.Hashmap(),
		AnyKeyValuePairs: coredynamic.EmptyMapAnyItems(),
		DynamicPayloads:  []byte{},
	}
}

func (it newAttributesCreator) Empty() *Attributes {
	return &Attributes{
		KeyValuePairs:    corestr.Empty.Hashmap(),
		AnyKeyValuePairs: coredynamic.EmptyMapAnyItems(),
		DynamicPayloads:  []byte{},
	}
}
