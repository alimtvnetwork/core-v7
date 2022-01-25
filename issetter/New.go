package issetter

import "errors"

func New(name string) (Value, error) {
	val, has := jsonValuesMap[name]

	if !has {
		return Uninitialized, errors.New(
			"cannot convert to issetter based on given value - " + name)
	}

	return Value(val.Value()), nil
}
