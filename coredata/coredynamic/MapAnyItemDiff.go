package coredynamic

import (
	"fmt"

	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/internal/mapdiffinternal"
)

type MapAnyItemDiff map[string]interface{}

func (it *MapAnyItemDiff) Length() int {
	if it == nil {
		return 0
	}

	return len(*it)
}

func (it MapAnyItemDiff) IsEmpty() bool {
	return it.Length() == 0
}

func (it MapAnyItemDiff) HasAnyItem() bool {
	return it.Length() > 0
}

func (it MapAnyItemDiff) LastIndex() int {
	return it.Length() - 1
}

func (it MapAnyItemDiff) AllKeysSorted() []string {
	return mapdiffinternal.MapStringAnyDiff(it.Raw()).AllKeysSorted()
}

func (it *MapAnyItemDiff) IsRawEqual(
	isRegardlessType bool,
	rightMap map[string]interface{},
) bool {
	differ := mapdiffinternal.
		MapStringAnyDiff(it.Raw())

	return differ.
		IsRawEqual(
			isRegardlessType,
			rightMap)
}

func (it *MapAnyItemDiff) HashmapDiffUsingRaw(
	isRegardlessType bool,
	rightMap map[string]interface{},
) MapAnyItemDiff {
	diffMap := it.DiffRaw(
		isRegardlessType,
		rightMap)

	if len(diffMap) == 0 {
		return map[string]interface{}{}
	}

	return diffMap
}

func (it *MapAnyItemDiff) MapAnyItems() *MapAnyItems {
	return &MapAnyItems{
		Items: it.Raw(),
	}
}

func (it *MapAnyItemDiff) DiffRaw(
	isRegardlessType bool,
	rightMap map[string]interface{},
) map[string]interface{} {
	differ := mapdiffinternal.
		MapStringAnyDiff(it.Raw())

	return differ.
		DiffRaw(
			isRegardlessType,
			rightMap)
}

func (it *MapAnyItemDiff) DiffJsonMessage(
	isRegardlessType bool,
	rightMap map[string]interface{},
) string {
	differ := mapdiffinternal.
		MapStringAnyDiff(it.Raw())

	return differ.DiffJsonMessage(
		isRegardlessType,
		rightMap)
}

func (it *MapAnyItemDiff) ToStringsSliceOfDiffMap(
	diffMap map[string]interface{},
) (diffSlice []string) {
	differ := mapdiffinternal.
		MapStringAnyDiff(it.Raw())

	return differ.ToStringsSliceOfDiffMap(
		diffMap)
}

func (it *MapAnyItemDiff) ShouldDiffMessage(
	isRegardlessType bool,
	title string,
	rightMap map[string]interface{},
) string {
	differ := mapdiffinternal.
		MapStringAnyDiff(it.Raw())

	return differ.ShouldDiffMessage(
		isRegardlessType,
		title,
		rightMap)
}

func (it *MapAnyItemDiff) LogShouldDiffMessage(
	isRegardlessType bool,
	title string,
	rightMap map[string]interface{},
) (diffMessage string) {
	differ := mapdiffinternal.
		MapStringAnyDiff(it.Raw())

	return differ.LogShouldDiffMessage(
		isRegardlessType,
		title,
		rightMap)
}

func (it *MapAnyItemDiff) Raw() map[string]interface{} {
	if it == nil {
		return map[string]interface{}{}
	}

	return *it
}

func (it *MapAnyItemDiff) Clear() MapAnyItemDiff {
	if it == nil {
		return map[string]interface{}{}
	}

	*it = map[string]interface{}{}

	return *it
}

func (it MapAnyItemDiff) Json() corejson.Result {
	return corejson.New(it)
}

func (it MapAnyItemDiff) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it MapAnyItemDiff) PrettyJsonString() string {
	return corejson.NewPtr(it).PrettyJsonString()
}

func (it MapAnyItemDiff) LogPrettyJsonString() {
	if it.IsEmpty() {
		fmt.Println("Empty Map")
	}

	prettyJson := it.Json().PrettyJsonString()

	fmt.Println(prettyJson)
}
