package coreinstruction

import (
	"gitlab.com/auk-go/core/constants"
)

type IdentifiersWithGlobals struct {
	IdentifierWithIsGlobals []IdentifierWithIsGlobal `json:"IdentifierWithIsGlobals"`
}

func EmptyIdentifiersWithGlobals() *IdentifiersWithGlobals {
	return &IdentifiersWithGlobals{
		IdentifierWithIsGlobals: []IdentifierWithIsGlobal{},
	}
}

func NewIdentifiersWithGlobals(
	isGlobal bool,
	ids ...string,
) *IdentifiersWithGlobals {
	slice := make(
		[]IdentifierWithIsGlobal,
		len(ids))

	if len(ids) == 0 {
		return &IdentifiersWithGlobals{
			IdentifierWithIsGlobals: slice,
		}
	}

	for i, id := range ids {
		slice[i] = IdentifierWithIsGlobal{
			BaseIdentifier: BaseIdentifier{
				Id: id,
			},
			IsGlobal: isGlobal,
		}
	}

	return &IdentifiersWithGlobals{
		IdentifierWithIsGlobals: slice,
	}
}

func (receiver *IdentifiersWithGlobals) Length() int {
	return len(receiver.IdentifierWithIsGlobals)
}

func (receiver *IdentifiersWithGlobals) IsEmpty() bool {
	return receiver.Length() == 0
}

func (receiver *IdentifiersWithGlobals) IndexOf(id string) int {
	if id == constants.EmptyString || receiver.IsEmpty() {
		return constants.InvalidNotFoundCase
	}

	for index, identifierWithIsGlobal := range receiver.IdentifierWithIsGlobals {
		if identifierWithIsGlobal.Id == id {
			return index
		}
	}

	return constants.InvalidNotFoundCase
}

func (receiver *IdentifiersWithGlobals) GetById(id string) *IdentifierWithIsGlobal {
	if id == constants.EmptyString || receiver.IsEmpty() {
		return nil
	}

	for _, identifierWithIsGlobal := range receiver.IdentifierWithIsGlobals {
		if identifierWithIsGlobal.Id == id {
			return &identifierWithIsGlobal
		}
	}

	return nil
}

func (receiver *IdentifiersWithGlobals) Add(
	isGlobal bool,
	id string,
) *IdentifiersWithGlobals {
	if id == constants.EmptyString {
		return receiver
	}

	receiver.IdentifierWithIsGlobals = append(
		receiver.IdentifierWithIsGlobals,
		*NewIdentifierWithIsGlobal(id, isGlobal))

	return receiver
}

func (receiver *IdentifiersWithGlobals) Adds(
	isGlobal bool,
	ids ...string,
) *IdentifiersWithGlobals {
	if len(ids) == 0 {
		return receiver
	}

	for _, id := range ids {
		receiver.IdentifierWithIsGlobals = append(
			receiver.IdentifierWithIsGlobals,
			*NewIdentifierWithIsGlobal(id, isGlobal))
	}

	return receiver
}

func (receiver *IdentifiersWithGlobals) HasAnyItem() bool {
	return receiver.Length() > 0
}

func (receiver *IdentifiersWithGlobals) Clone() *IdentifiersWithGlobals {
	length := receiver.Length()

	slice := make(
		[]IdentifierWithIsGlobal,
		length)

	if length == 0 {
		return &IdentifiersWithGlobals{
			IdentifierWithIsGlobals: slice,
		}
	}

	for i, idWithIsGlobal := range receiver.IdentifierWithIsGlobals {
		slice[i] = *idWithIsGlobal.Clone()
	}

	return &IdentifiersWithGlobals{
		IdentifierWithIsGlobals: slice,
	}
}
