package coreinstruction

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

type Rename struct {
	Existing string `json:"Existing,omitempty"`
	New      string `json:"New,omitempty"`
}

func (it *Rename) IsNull() bool {
	return it == nil
}

func (it *Rename) IsExistingEmpty() bool {
	return it == nil || it.Existing == ""
}

func (it *Rename) IsNewEmpty() bool {
	return it == nil || it.New == ""
}

func (it Rename) String() string {
	if it.IsNull() {
		return "Rename null!"
	}

	return fmt.Sprintf(
		constants.RenameFormat,
		it.Existing,
		it.New)
}

func (it *Rename) SourceDestination() *SourceDestination {
	if it == nil {
		return nil
	}

	return &SourceDestination{
		Source:      it.Existing,
		Destination: it.New,
	}
}

func (it Rename) FromName() string {
	return it.Existing
}

func (it Rename) ExistingName() string {
	return it.Existing
}

func (it Rename) NewName() string {
	return it.New
}

func (it Rename) ToName() string {
	return it.New
}

func (it *Rename) SetFromName(form string) {
	it.Existing = form
}

func (it *Rename) SetToName(to string) {
	it.New = to
}

func (it *Rename) FromTo() *FromTo {
	if it == nil {
		return nil
	}

	return &FromTo{
		From: it.Existing,
		To:   it.New,
	}
}

func (it *Rename) Clone() *Rename {
	if it == nil {
		return nil
	}

	return &Rename{
		Existing: it.Existing,
		New:      it.New,
	}
}
