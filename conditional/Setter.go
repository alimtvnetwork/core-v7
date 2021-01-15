package conditional

import "gitlab.com/evatix-go/core/issetter"

func Setter(
	isTrue bool,
	value, defaultVal issetter.Value,
) interface{} {
	if isTrue {
		return value
	}

	return defaultVal
}
