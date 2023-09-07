package coreinstruction

import "gitlab.com/auk-go/core/coredata/corejson"

type NameListCollection struct {
	NameLists []NameList
}

func (it *NameListCollection) IsNull() bool {
	return it == nil
}

func (it *NameListCollection) IsAnyNull() bool {
	return it == nil || it.NameLists == nil
}

func (it *NameListCollection) IsEmpty() bool {
	return it == nil || len(it.NameLists) == 0
}

func (it *NameListCollection) HasAnyItem() bool {
	return !it.IsEmpty()
}

func (it *NameListCollection) Length() int {
	if it == nil {
		return 0
	}

	return len(it.NameLists)
}

func (it NameListCollection) String() string {
	if it.IsNull() {
		return ""
	}

	return corejson.
		Serialize.
		ToString(it)
}
