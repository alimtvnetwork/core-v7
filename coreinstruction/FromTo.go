package coreinstruction

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

type FromTo struct {
	From string `json:"From,omitempty"`
	To   string `json:"To,omitempty"`
}

func (it *FromTo) IsNull() bool {
	return it == nil
}

func (it *FromTo) IsFromEmpty() bool {
	return it == nil || it.From == ""
}

func (it *FromTo) IsToEmpty() bool {
	return it == nil || it.To == ""
}

func (it FromTo) String() string {
	if it.IsNull() {
		return "FromTo null!"
	}

	return fmt.Sprintf(
		constants.FromToFormat,
		it.From,
		it.To)
}

func (it FromTo) FromName() string {
	return it.From
}

func (it FromTo) ToName() string {
	return it.To
}

func (it *FromTo) SetFromName(form string) {
	it.From = form
}

func (it *FromTo) SetToName(to string) {
	it.To = to
}

func (it *FromTo) SourceDestination() *SourceDestination {
	if it == nil {
		return nil
	}

	return &SourceDestination{
		Source:      it.From,
		Destination: it.To,
	}
}

func (it *FromTo) Rename() *Rename {
	if it == nil {
		return nil
	}

	return &Rename{
		Existing: it.From,
		New:      it.To,
	}
}

func (it FromTo) Clone() FromTo {
	return FromTo{
		From: it.From,
		To:   it.To,
	}
}

func (it *FromTo) ClonePtr() *FromTo {
	if it == nil {
		return nil
	}

	return &FromTo{
		From: it.From,
		To:   it.To,
	}
}
