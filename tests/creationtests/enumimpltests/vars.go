package enumimpltests

import (
	"gitlab.com/evatix-go/core/coretests"
	"gitlab.com/evatix-go/core/errcore"
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
