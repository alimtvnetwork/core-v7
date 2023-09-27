package enumimpltests

import (
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/errcore"
)

type EnumImplDynamicMapTestWrapper struct {
	coretests.BaseTestCase
}

func (it EnumImplDynamicMapTestWrapper) ArrangeAsLeftRightDynamicMap() LeftRightDynamicMap {
	casted, isSuccess := it.ArrangeInput.(LeftRightDynamicMap)

	if !isSuccess {
		errcore.HandleErrMessage("casting failed to LeftRightDynamicMap")
	}

	return casted
}

func (it EnumImplDynamicMapTestWrapper) ArrangeAsLeftRightDynamicMapWithDefaultChecker() LeftRightDynamicMapWithDefaultChecker {
	casted, isSuccess := it.ArrangeInput.(LeftRightDynamicMapWithDefaultChecker)

	if !isSuccess {
		errcore.HandleErrMessage("casting failed to LeftRightDynamicMap")
	}

	return casted
}
