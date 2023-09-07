package conditional

import "gitlab.com/auk-go/core/issetter"

func Setter(
	isTrue bool,
	trueValue, falseValue issetter.Value,
) issetter.Value {
	if isTrue {
		return trueValue
	}

	return falseValue
}
