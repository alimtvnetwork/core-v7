package coreinstruction

import (
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coredata/corestr"
)

type NameList struct {
	Name string               `json:"Name,omitempty"`
	List *corestr.SimpleSlice `json:"List,omitempty"`
}

func (it *NameList) IsNull() bool {
	return it == nil
}

func (it *NameList) IsAnyNull() bool {
	return it == nil || it.List == nil
}

func (it *NameList) IsNameEmpty() bool {
	return it == nil || it.Name == ""
}

func (it *NameList) HasName() bool {
	return it != nil && it.Name != ""
}

func (it NameList) String() string {
	if it.IsNull() {
		return ""
	}

	return corejson.
		Serialize.
		ToString(it)
}

func (it *NameList) Clone(
	isDeepClone bool,
) *NameList {
	if it == nil {
		return nil
	}

	return &NameList{
		Name: it.Name,
		List: it.
			List.
			ClonePtr(isDeepClone),
	}
}

func (it *NameList) DeepClone() *NameList {
	return it.Clone(true)
}
