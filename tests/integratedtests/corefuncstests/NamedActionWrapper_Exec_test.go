package corefuncstests

import (
	"errors"
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/corefuncs"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ============================================================================
// NamedActionFuncWrapper
// ============================================================================

func Test_NamedActionWrapper_Exec_Verification(t *testing.T) {
	for caseIndex, tc := range namedActionWrapperTestCases {
		// Arrange
		called := false
		receivedName := ""
		wrapper := corefuncs.New.NamedAction("my-named", func(name string) {
			called = true
			receivedName = name
		})

		// Act
		wrapper.Exec()

		// Assert
		actual := args.Map{
			"called": called,
			"name":   receivedName,
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_NamedActionWrapper_AsActionFunc_Verification(t *testing.T) {
	for caseIndex, tc := range namedActionAsActionFuncTestCases {
		// Arrange
		called := false
		wrapper := corefuncs.New.NamedAction("test", func(name string) {
			called = true
		})

		// Act
		actionFunc := wrapper.AsActionFunc()
		actionFunc()

		// Assert
		actual := args.Map{
			"called": called,
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_NamedActionWrapper_AsActionReturnsErrorFunc_Verification(t *testing.T) {
	for caseIndex, tc := range namedActionAsErrFuncTestCases {
		// Arrange
		wrapper := corefuncs.New.NamedAction("test", func(name string) {})

		// Act
		errFunc := wrapper.AsActionReturnsErrorFunc()
		err := errFunc()

		// Assert
		actual := args.Map{
			"isNil": err == nil,
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_NamedActionWrapper_Next_Verification(t *testing.T) {
	for caseIndex, tc := range namedActionNextTestCases {
		// Arrange
		firstCalled := false
		secondCalled := false
		first := corefuncs.New.NamedAction("first", func(name string) {
			firstCalled = true
		})
		second := corefuncs.New.NamedAction("second", func(name string) {
			secondCalled = true
		})

		// Act
		first.Next(&second)

		// Assert
		actual := args.Map{
			"firstCalled":  firstCalled,
			"secondCalled": secondCalled,
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ============================================================================
// InOutErrFuncWrapper (legacy)
// ============================================================================

func Test_LegacyInOutErr_Exec_Verification(t *testing.T) {
	for caseIndex, tc := range legacyInOutErrExecTestCases {
		// Arrange
		wrapper := corefuncs.New.LegacyInOutErr("strlen", func(input any) (any, error) {
			return len(input.(string)), nil
		})

		// Act
		result, err := wrapper.Exec("hello")

		// Assert
		actual := args.Map{
			"result": fmt.Sprintf("%v", result),
			"isNil":  err == nil,
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_LegacyInOutErr_AsErrFunc_Success_Verification(t *testing.T) {
	for caseIndex, tc := range legacyInOutErrAsErrFuncSuccessTestCases {
		// Arrange
		wrapper := corefuncs.New.LegacyInOutErr("ok", func(input any) (any, error) {
			return nil, nil
		})

		// Act
		errFunc := wrapper.AsActionReturnsErrorFunc("anything")
		err := errFunc()

		// Assert
		actual := args.Map{
			"isNil": err == nil,
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_LegacyInOutErr_AsErrFunc_Fail_Verification(t *testing.T) {
	for caseIndex, tc := range legacyInOutErrAsErrFuncFailTestCases {
		// Arrange
		wrapper := corefuncs.New.LegacyInOutErr("fail", func(input any) (any, error) {
			return nil, errors.New("failed")
		})

		// Act
		errFunc := wrapper.AsActionReturnsErrorFunc("anything")
		err := errFunc()

		// Assert
		actual := args.Map{
			"hasError": err != nil,
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ============================================================================
// ResultDelegatingFuncWrapper (legacy)
// ============================================================================

func Test_LegacyResultDelegating_Exec_Verification(t *testing.T) {
	for caseIndex, tc := range legacyResultDelegatingExecTestCases {
		// Arrange
		wrapper := corefuncs.New.LegacyResultDelegating("unmarshal", func(target any) error {
			ptr := target.(*string)
			*ptr = "delegated"
			return nil
		})

		// Act
		var result string
		err := wrapper.Exec(&result)

		// Assert
		actual := args.Map{
			"isNil":  err == nil,
			"result": result,
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_LegacyResultDelegating_AsErrFunc_Success_Verification(t *testing.T) {
	for caseIndex, tc := range legacyResultDelegatingAsErrFuncSuccessTestCases {
		// Arrange
		wrapper := corefuncs.New.LegacyResultDelegating("ok", func(target any) error {
			return nil
		})

		// Act
		errFunc := wrapper.AsActionReturnsErrorFunc("anything")
		err := errFunc()

		// Assert
		actual := args.Map{
			"isNil": err == nil,
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_LegacyResultDelegating_AsErrFunc_Fail_Verification(t *testing.T) {
	for caseIndex, tc := range legacyResultDelegatingAsErrFuncFailTestCases {
		// Arrange
		wrapper := corefuncs.New.LegacyResultDelegating("fail", func(target any) error {
			return errors.New("delegation failed")
		})

		// Act
		errFunc := wrapper.AsActionReturnsErrorFunc("anything")
		err := errFunc()

		// Assert
		actual := args.Map{
			"hasError": err != nil,
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ============================================================================
// GetFunc
// ============================================================================

func Test_GetFunc_Runtime_Verification(t *testing.T) {
	for caseIndex, tc := range getFuncRuntimeTestCases {
		// Act
		runtimeFunc := corefuncs.GetFunc(sampleFunc)

		// Assert
		actual := args.Map{
			"notNil": runtimeFunc != nil,
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ============================================================================
// Generic Wrappers - NewInOutWrapper
// ============================================================================

func Test_NewInOutWrapper_Exec_Verification(t *testing.T) {
	for caseIndex, tc := range newInOutWrapperTestCases {
		// Arrange
		wrapper := corefuncs.NewInOutWrapper[string, int](
			"strlen",
			func(s string) int { return len(s) },
		)

		// Act
		result := wrapper.Exec("test")

		// Assert
		actual := args.Map{
			"result": result,
			"name":   wrapper.Name,
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ============================================================================
// Generic Wrappers - NewInActionErrWrapper
// ============================================================================

func Test_NewInActionErrWrapper_Success_Verification(t *testing.T) {
	for caseIndex, tc := range newInActionErrWrapperSuccessTestCases {
		// Arrange
		wrapper := corefuncs.NewInActionErrWrapper[string](
			"validate",
			func(s string) error { return nil },
		)

		// Act
		err := wrapper.Exec("valid")

		// Assert
		actual := args.Map{
			"isNil": err == nil,
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_NewInActionErrWrapper_Fail_Verification(t *testing.T) {
	for caseIndex, tc := range newInActionErrWrapperFailTestCases {
		// Arrange
		wrapper := corefuncs.NewInActionErrWrapper[string](
			"validate",
			func(s string) error {
				if s == "" {
					return errors.New("empty")
				}
				return nil
			},
		)

		// Act
		err := wrapper.Exec("")

		// Assert
		actual := args.Map{
			"hasError": err != nil,
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ============================================================================
// Generic Wrappers - NewResultDelegatingWrapper
// ============================================================================

func Test_NewResultDelegatingWrapper_Verification(t *testing.T) {
	for caseIndex, tc := range newResultDelegatingWrapperTestCases {
		// Arrange
		wrapper := corefuncs.NewResultDelegatingWrapper[*string](
			"unmarshal",
			func(target *string) error {
				*target = "typed-delegated"
				return nil
			},
		)

		// Act
		var result string
		err := wrapper.Exec(&result)

		// Assert
		actual := args.Map{
			"isNil":  err == nil,
			"result": result,
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ============================================================================
// Generic Wrappers - NewSerializeWrapper
// ============================================================================

func Test_NewSerializeWrapper_Verification(t *testing.T) {
	for caseIndex, tc := range newSerializeWrapperTestCases {
		// Arrange
		wrapper := corefuncs.NewSerializeWrapper[string](
			"marshal",
			func(s string) ([]byte, error) {
				return []byte(s), nil
			},
		)

		// Act
		bytes, err := wrapper.Exec("test")

		// Assert
		actual := args.Map{
			"hasBytes": len(bytes) > 0,
			"isNil":    err == nil,
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ============================================================================
// IsSuccessFuncWrapper additional methods
// ============================================================================

func Test_IsSuccessWrapper_AsActionFunc_Verification(t *testing.T) {
	for caseIndex, tc := range isSuccessAsActionFuncTestCases {
		// Arrange
		called := false
		wrapper := corefuncs.New.IsSuccess("check", func() bool {
			called = true
			return true
		})

		// Act
		actionFunc := wrapper.AsActionFunc()
		actionFunc()

		// Assert
		actual := args.Map{
			"called": called,
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IsSuccessWrapper_AsErrFunc_Success_Verification(t *testing.T) {
	for caseIndex, tc := range isSuccessAsErrFuncSuccessTestCases {
		// Arrange
		wrapper := corefuncs.New.IsSuccess("check", func() bool { return true })

		// Act
		errFunc := wrapper.AsActionReturnsErrorFunc()
		err := errFunc()

		// Assert
		actual := args.Map{
			"isNil": err == nil,
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IsSuccessWrapper_AsErrFunc_Fail_Verification(t *testing.T) {
	for caseIndex, tc := range isSuccessAsErrFuncFailTestCases {
		// Arrange
		wrapper := corefuncs.New.IsSuccess("check", func() bool { return false })

		// Act
		errFunc := wrapper.AsActionReturnsErrorFunc()
		err := errFunc()

		// Assert
		actual := args.Map{
			"hasError": err != nil,
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ============================================================================
// ActionReturnsErrorFuncWrapper additional
// ============================================================================

func Test_ActionErrWrapper_AsErrFunc_Success_Verification(t *testing.T) {
	for caseIndex, tc := range actionErrAsActionFuncTestCases {
		// Arrange
		wrapper := corefuncs.New.ActionErr("cleanup", func() error { return nil })

		// Act
		errFunc := wrapper.AsActionReturnsErrorFunc()
		err := errFunc()

		// Assert
		actual := args.Map{
			"isNil": err == nil,
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_ActionErrWrapper_AsErrFunc_WithError_Verification(t *testing.T) {
	for caseIndex, tc := range actionErrAsErrFuncWithErrorTestCases {
		// Arrange
		wrapper := corefuncs.New.ActionErr("cleanup", func() error {
			return errors.New("cleanup failed")
		})

		// Act
		errFunc := wrapper.AsActionReturnsErrorFunc()
		err := errFunc()

		// Assert
		actual := args.Map{
			"hasError": err != nil,
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ============================================================================
// ToLegacy converters
// ============================================================================

func Test_InOutWrapper_ToLegacy_Verification(t *testing.T) {
	for caseIndex, tc := range inOutWrapperToLegacyTestCases {
		// Arrange
		wrapper := corefuncs.NewInOutWrapper[string, int](
			"legacy-conv",
			func(s string) int { return len(s) },
		)

		// Act
		legacy := wrapper.ToLegacy()
		result, _ := legacy.Exec("hello")

		// Assert
		actual := args.Map{
			"name":   legacy.Name,
			"result": fmt.Sprintf("%v", result),
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_InActionErr_ToLegacy_Verification(t *testing.T) {
	for caseIndex, tc := range inActionErrToLegacyTestCases {
		// Arrange
		wrapper := corefuncs.NewInActionErrWrapper[string](
			"legacy-validate",
			func(s string) error { return nil },
		)

		// Act
		legacy := wrapper.ToLegacy()

		// Assert
		actual := args.Map{
			"name": legacy.Name,
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_ResultDelegating_ToLegacy_Verification(t *testing.T) {
	for caseIndex, tc := range resultDelegatingToLegacyTestCases {
		// Arrange
		wrapper := corefuncs.NewResultDelegatingWrapper[*string](
			"legacy-unmarshal",
			func(target *string) error { return nil },
		)

		// Act
		legacy := wrapper.ToLegacy()

		// Assert
		actual := args.Map{
			"name": legacy.Name,
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
