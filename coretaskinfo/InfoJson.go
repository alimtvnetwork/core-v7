package coretaskinfo

import (
	"strings"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/coredata/corejson"
)

func (it Info) Json() corejson.Result {
	return corejson.New(it)
}

func (it Info) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it *Info) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	return jsonResult.Deserialize(it)
}

func (it *Info) JsonString() string {
	if it.IsNull() {
		return ""
	}

	return it.JsonPtr().JsonString()
}

func (it *Info) PrettyJsonStringWithPayloads(
	payloads []byte,
) string {
	return corejson.
		NewPtr(it.MapWithPayload(payloads)).
		PrettyJsonString()
}

func (it *Info) PrettyJsonString() string {
	if it.IsNull() {
		return ""
	}

	return corejson.
		NewPtr(it).
		PrettyJsonString()
}

func (it *Info) LazyMapPrettyJsonString() string {
	lazyMap := it.LazyMap()

	return corejson.
		NewPtr(lazyMap).
		PrettyJsonString()
}

func (it Info) JsonStringMust() string {
	jsonResult := it.Json()
	jsonResult.MustBeSafe()

	return jsonResult.JsonString()
}

func (it *Info) Serialize() ([]byte, error) {
	return corejson.Serialize.Raw(it)
}

func (it *Info) Deserialize(toPtr any) (parsingErr error) {
	return it.JsonPtr().Deserialize(toPtr)
}

func (it *Info) ExamplesAsString() (compiledString string) {
	if it.IsNull() {
		return ""
	}

	return strings.Join(
		it.Examples,
		constants.CommaSpace)
}

func (it *Info) String() string {
	if it == nil {
		return ""
	}

	return it.PrettyJsonString()
}

func (it Info) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return &it
}
