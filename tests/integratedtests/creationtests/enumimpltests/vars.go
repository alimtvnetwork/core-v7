package enumimpltests

import (
	"reflect"

	"gitlab.com/auk-go/core/coreimpl/enumimpl"
	"gitlab.com/auk-go/core/coretests"
)

var (
	checker1                                     = enumimpl.LeftRightDiffCheckerImpl
	typeVerifyOfForDynamicMapSimpleDiffTestCases = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf(LeftRightDynamicMapWithDefaultChecker{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}

	simpleDiffTestCases = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf(LeftRightDynamicMapWithDefaultChecker{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf(""),
	}

	typeVerifyOfForDynamicMapDiffTestCases = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf(LeftRightDynamicMap{}),
		ActualInput:   reflect.TypeOf(enumimpl.DynamicMap{}),
		ExpectedInput: reflect.TypeOf(enumimpl.DynamicMap{}),
	}

	typeVerifyOfDynamicMapDiffMessageTestCases = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf(LeftRightDynamicMap{}),
		ActualInput:   reflect.TypeOf(""),
		ExpectedInput: reflect.TypeOf(""),
	}
)
