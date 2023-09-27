package enumimpltests

import "gitlab.com/auk-go/core/coreimpl/enumimpl"

type LeftRightDynamicMapWithDefaultChecker struct {
	LeftRightDynamicMap
	enumimpl.DifferChecker
}
