package conditional

import "gitlab.com/auk-go/core/issetter"

func SetterDefault(
	currentSetter issetter.Value,
	defVal issetter.Value,
) issetter.Value {
	if currentSetter.IsUnSetOrUninitialized() {
		return defVal
	}

	return currentSetter
}
