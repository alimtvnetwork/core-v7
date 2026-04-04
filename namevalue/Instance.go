package namevalue

import (
	"encoding/json"
	"fmt"

	"github.com/alimtvnetwork/core/constants"
)

type Instance[K comparable, V any] struct {
	Name  K
	Value V
}

func (it *Instance[K, V]) IsNull() bool {
	return it == nil
}

func (it *Instance[K, V]) String() string {
	if it.IsNull() {
		return constants.EmptyString
	}

	return fmt.Sprintf(
		"%v = %v",
		it.Name,
		it.Value)
}

func (it *Instance[K, V]) JsonString() string {
	if it.IsNull() {
		return constants.EmptyString
	}

	rawBytes, err := json.Marshal(it)

	if err != nil || rawBytes == nil {
		return constants.EmptyString
	}

	return string(rawBytes)
}

func (it *Instance[K, V]) Dispose() {
	if it == nil {
		return
	}

	var zeroK K
	var zeroV V

	it.Name = zeroK
	it.Value = zeroV
}
