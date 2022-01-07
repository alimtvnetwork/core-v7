package taskpriority

import (
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/core/simplewrap"
)

func New(priorityName string) (Variant, error) {
	if priorityName == "" {
		return Invalid, errcore.
			EmptyItemsType.
			ErrorNoRefs("empty string cannot be converted to task-priority type")
	}

	variant, has := nameToVariantMap[priorityName]

	if has {
		return variant, nil
	}

	message := simplewrap.WithCurly(priorityName) + " is not found in the map!"
	err := errcore.
		EmptyItemsType.
		ErrorNoRefs(message)

	return Invalid, err
}
