package coredynamic

import "gitlab.com/auk-go/core/constants"

type DynamicStatus struct {
	Dynamic
	Index   int
	Message string
}

func InvalidDynamicStatusNoMessage() *DynamicStatus {
	return InvalidDynamicStatus(constants.EmptyString)
}

func InvalidDynamicStatus(message string) *DynamicStatus {
	return &DynamicStatus{
		Dynamic: NewDynamic(nil, false),
		Index:   constants.InvalidNotFoundCase,
		Message: message,
	}
}

// Clone Warning: Cannot clone dynamic data or interface properly but set it again
//
// If it is a pointer one needs to copy it manually.
func (it DynamicStatus) Clone() DynamicStatus {
	return DynamicStatus{
		Dynamic: it.Dynamic.Clone(),
		Index:   constants.InvalidNotFoundCase,
		Message: it.Message,
	}
}

func (it *DynamicStatus) ClonePtr() *DynamicStatus {
	if it == nil {
		return nil
	}

	return &DynamicStatus{
		Dynamic: it.Dynamic.Clone(),
		Index:   constants.InvalidNotFoundCase,
		Message: it.Message,
	}
}
