package coreinstruction

import (
	"gitlab.com/evatix-go/core/constants"
)

type Identifiers struct {
	IdentifierWithIsGlobals []IdentifierWithIsGlobal `json:"IdentifierWithIsGlobals"`
}

func EmptyIdentifiers() *Identifiers {
	return &Identifiers{
		IdentifierWithIsGlobals: []IdentifierWithIsGlobal{},
	}
}

func NewIdentifiers(
	isGlobal bool,
	ids ...string,
) *Identifiers {
	slice := make(
		[]IdentifierWithIsGlobal,
		len(ids))

	if len(ids) == 0 {
		return &Identifiers{
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

	return &Identifiers{
		IdentifierWithIsGlobals: slice,
	}
}

func (receiver *Identifiers) Length() int {
	return len(receiver.IdentifierWithIsGlobals)
}

func (receiver *Identifiers) IsEmpty() bool {
	return receiver.Length() == 0
}

func (receiver *Identifiers) IndexOf(id string) int {
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

func (receiver *Identifiers) GetById(id string) *IdentifierWithIsGlobal {
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

func (receiver *Identifiers) Add(
	isGlobal bool,
	id string,
) *Identifiers {
	if id == constants.EmptyString {
		return receiver
	}

	receiver.IdentifierWithIsGlobals = append(
		receiver.IdentifierWithIsGlobals,
		*NewIdentifierWithIsGlobal(id, isGlobal))

	return receiver
}

func (receiver *Identifiers) Adds(
	isGlobal bool,
	ids ...string,
) *Identifiers {
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

func (receiver *Identifiers) HasAnyItem() bool {
	return receiver.Length() > 0
}

func (receiver *Identifiers) Clone() *Identifiers {
	length := receiver.Length()

	slice := make(
		[]IdentifierWithIsGlobal,
		length)

	if length == 0 {
		return &Identifiers{
			IdentifierWithIsGlobals: slice,
		}
	}

	for i, idWithIsGlobal := range receiver.IdentifierWithIsGlobals {
		slice[i] = *idWithIsGlobal.Clone()
	}

	return &Identifiers{
		IdentifierWithIsGlobals: slice,
	}
}
