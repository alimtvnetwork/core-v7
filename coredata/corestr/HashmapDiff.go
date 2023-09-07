package corestr

import (
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/internal/mapdiffinternal"
)

type HashmapDiff map[string]string

func (it *HashmapDiff) Length() int {
	if it == nil {
		return 0
	}

	return len(*it)
}

func (it HashmapDiff) IsEmpty() bool {
	return it.Length() == 0
}

func (it HashmapDiff) HasAnyItem() bool {
	return it.Length() > 0
}

func (it HashmapDiff) LastIndex() int {
	return it.Length() - 1
}

func (it HashmapDiff) AllKeysSorted() []string {
	return mapdiffinternal.HashmapDiff(it.Raw()).AllKeysSorted()
}

func (it *HashmapDiff) MapAnyItems() map[string]interface{} {
	if it == nil || len(*it) == 0 {
		return map[string]interface{}{}
	}

	newMap := make(
		map[string]interface{},
		it.Length()+1)

	for name, value := range *it {
		newMap[name] = value
	}

	return newMap
}

func (it *HashmapDiff) HasAnyChanges(
	rightMap map[string]string,
) bool {
	return !it.IsRawEqual(
		rightMap)
}

func (it *HashmapDiff) RawMapStringAnyDiff() mapdiffinternal.MapStringAnyDiff {
	return it.MapAnyItems()
}

func (it *HashmapDiff) IsRawEqual(
	rightMap map[string]string,
) bool {
	differ := mapdiffinternal.
		HashmapDiff(it.Raw())

	return differ.
		IsRawEqual(rightMap)
}

func (it *HashmapDiff) HashmapDiffUsingRaw(
	rightMap map[string]string,
) HashmapDiff {
	diffMap := it.DiffRaw(
		rightMap)

	if len(diffMap) == 0 {
		return map[string]string{}
	}

	return diffMap
}

func (it *HashmapDiff) DiffRaw(
	rightMap map[string]string,
) map[string]string {
	differ := mapdiffinternal.
		HashmapDiff(it.Raw())

	return differ.
		DiffRaw(rightMap)
}

func (it *HashmapDiff) DiffJsonMessage(
	rightMap map[string]string,
) string {
	differ := mapdiffinternal.
		HashmapDiff(it.Raw())

	return differ.DiffJsonMessage(
		rightMap)
}

func (it *HashmapDiff) ToStringsSliceOfDiffMap(
	diffMap map[string]string,
) (diffSlice []string) {
	differ := mapdiffinternal.
		HashmapDiff(it.Raw())

	return differ.ToStringsSliceOfDiffMap(
		diffMap)
}

func (it *HashmapDiff) ShouldDiffMessage(
	title string,
	rightMap map[string]string,
) string {
	differ := mapdiffinternal.
		HashmapDiff(it.Raw())

	return differ.ShouldDiffMessage(
		title,
		rightMap)
}

func (it *HashmapDiff) LogShouldDiffMessage(
	title string,
	rightMap map[string]string,
) (diffMessage string) {
	differ := mapdiffinternal.
		HashmapDiff(it.Raw())

	return differ.LogShouldDiffMessage(
		title,
		rightMap)
}

func (it *HashmapDiff) Raw() map[string]string {
	if it == nil {
		return map[string]string{}
	}

	return *it
}

func (it *HashmapDiff) Serialize() ([]byte, error) {
	return corejson.Serialize.Raw(it.Raw())
}

func (it *HashmapDiff) Deserialize(toPtr interface{}) (parsingErr error) {
	return corejson.NewPtr(it.Raw()).Deserialize(toPtr)
}
