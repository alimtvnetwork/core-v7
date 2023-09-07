package corepayload

import (
	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coredata/corestr"
)

type emptyCreator struct{}

func (it *emptyCreator) Attributes() *Attributes {
	return &Attributes{}
}

func (it *emptyCreator) AttributesDefaults() *Attributes {
	return &Attributes{
		KeyValuePairs:    corestr.Empty.Hashmap(),
		AnyKeyValuePairs: coredynamic.EmptyMapAnyItems(),
		DynamicPayloads:  []byte{},
	}
}

func (it *emptyCreator) PayloadWrapper() *PayloadWrapper {
	return &PayloadWrapper{}
}

func (it *emptyCreator) PayloadsCollection() *PayloadsCollection {
	return &PayloadsCollection{
		Items: []*PayloadWrapper{},
	}
}
