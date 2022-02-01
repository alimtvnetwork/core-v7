package coreinstruction

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
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
