package namevalue

import (
	"encoding/json"
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

type Instance struct {
	Name  string
	Value interface{}
}

func (it *Instance) IsNull() bool {
	return it == nil
}

func (it Instance) String() string {
	if it.IsNull() {
		return constants.EmptyString
	}

	return fmt.Sprintf(
		constants.KeyValShortFormat,
		it.Name,
		it.Value)
}

func (it Instance) JsonString() string {
	if it.IsNull() {
		return constants.EmptyString
	}

	rawBytes, err := json.Marshal(it)

	if err != nil || rawBytes == nil {
		return constants.EmptyString
	}

	return string(rawBytes)
}

func (it *Instance) Dispose() {
	if it == nil {
		return
	}

	it.Name = constants.EmptyString
	it.Value = nil
}
