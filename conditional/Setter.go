package conditional

import "gitlab.com/evatix-go/core/issetter"

func Setter(
	isTrue bool,
	trueValue, falseValue issetter.Value,
) issetter.Value {
	if isTrue {
		return trueValue
	}

	return falseValue
}
