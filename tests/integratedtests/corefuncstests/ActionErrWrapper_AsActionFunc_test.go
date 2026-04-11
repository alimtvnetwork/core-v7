package corefuncstests

import (
	"errors"
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/corefuncs"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================================================
// ActionReturnsErrorFuncWrapper.AsActionFunc
// ==========================================================================

var covActionErrAsActionFuncTestCases = []coretestcases.CaseV1{
	{
		Title:         "AsActionFunc executes without panic on success",
		ArrangeInput:  args.Map{},
		ExpectedInput: args.Map{"called": true},
	},
}

func Test_ActionErrWrapper_AsActionFunc_Coverage(t *testing.T) {
	for caseIndex, testCase := range covActionErrAsActionFuncTestCases {
		// Arrange
		called := false
		wrapper := corefuncs.New.ActionErr("test", func() error {
			called = true
			return nil
		})

		actionFunc := wrapper.AsActionFunc()
		actionFunc()

		// Act
		actual := args.Map{"called": called}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// InOutErrFuncWrapperOf.AsActionReturnsErrorFunc
// ==========================================================================

var covInOutErrOfAsErrFuncTestCases = []coretestcases.CaseV1{
	{
		Title:         "AsActionReturnsErrorFunc success returns nil",
		ArrangeInput:  args.Map{"scenario": "success"},
		ExpectedInput: args.Map{"isNil": true},
	},
	{
		Title:         "AsActionReturnsErrorFunc failure returns error",
		ArrangeInput:  args.Map{"scenario": "failure"},
		ExpectedInput: args.Map{"isNil": false},
	},
}

func Test_InOutErrWrapperOf_AsErrFunc_Coverage(t *testing.T) {
	for caseIndex, testCase := range covInOutErrOfAsErrFuncTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		scenario, _ := input.GetAsString("scenario")

		var errFunc corefuncs.ActionReturnsErrorFunc

		switch scenario {
		case "success":
			wrapper := corefuncs.NewInOutErrWrapper[string, int](
				"ok", func(s string) (int, error) { return 1, nil },
			)
			errFunc = wrapper.AsActionReturnsErrorFunc("test")
		case "failure":
			wrapper := corefuncs.NewInOutErrWrapper[string, int](
				"fail", func(s string) (int, error) { return 0, errors.New("fail") },
			)
			errFunc = wrapper.AsActionReturnsErrorFunc("test")
		}

		err := errFunc()

		// Act
		actual := args.Map{"isNil": err == nil}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// InOutFuncWrapperOf.AsActionFunc / AsActionReturnsErrorFunc
// ==========================================================================

var covInOutFuncOfTestCases = []coretestcases.CaseV1{
	{
		Title:         "InOutFuncWrapperOf.AsActionFunc calls action",
		ArrangeInput:  args.Map{"method": "AsActionFunc"},
		ExpectedInput: args.Map{"called": true},
	},
	{
		Title:         "InOutFuncWrapperOf.AsActionReturnsErrorFunc returns nil",
		ArrangeInput:  args.Map{"method": "AsActionReturnsErrorFunc"},
		ExpectedInput: args.Map{"isNil": true},
	},
}

func Test_InOutFuncWrapperOf_Coverage(t *testing.T) {
	for caseIndex, testCase := range covInOutFuncOfTestCases {
		input := testCase.ArrangeInput.(args.Map)
		method, _ := input.GetAsString("method")

		called := false
		wrapper := corefuncs.NewInOutWrapper[string, int](
			"test", func(s string) int {
				called = true
				return len(s)
			},
		)

		var actual args.Map

		switch method {
		case "AsActionFunc":
			actionFunc := wrapper.AsActionFunc("hello")
			actionFunc()
			actual = args.Map{"called": called}
		case "AsActionReturnsErrorFunc":
			errFunc := wrapper.AsActionReturnsErrorFunc("hello")
			err := errFunc()
			actual = args.Map{"isNil": err == nil}
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// InActionReturnsErrFuncWrapperOf.AsActionReturnsErrorFunc
// ==========================================================================

var covInActionErrOfAsErrFuncTestCases = []coretestcases.CaseV1{
	{
		Title:         "AsActionReturnsErrorFunc success returns nil",
		ArrangeInput:  args.Map{"scenario": "success"},
		ExpectedInput: args.Map{"isNil": true},
	},
	{
		Title:         "AsActionReturnsErrorFunc failure returns error",
		ArrangeInput:  args.Map{"scenario": "failure"},
		ExpectedInput: args.Map{"isNil": false},
	},
}

func Test_InActionErrWrapperOf_AsErrFunc_Coverage(t *testing.T) {
	for caseIndex, testCase := range covInActionErrOfAsErrFuncTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		scenario, _ := input.GetAsString("scenario")

		var errFunc corefuncs.ActionReturnsErrorFunc

		switch scenario {
		case "success":
			wrapper := corefuncs.NewInActionErrWrapper[string](
				"ok", func(s string) error { return nil },
			)
			errFunc = wrapper.AsActionReturnsErrorFunc("test")
		case "failure":
			wrapper := corefuncs.NewInActionErrWrapper[string](
				"fail", func(s string) error { return errors.New("fail") },
			)
			errFunc = wrapper.AsActionReturnsErrorFunc("test")
		}

		err := errFunc()

		// Act
		actual := args.Map{"isNil": err == nil}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// ResultDelegatingFuncWrapperOf.AsActionReturnsErrorFunc
// ==========================================================================

var covResultDelegatingOfAsErrFuncTestCases = []coretestcases.CaseV1{
	{
		Title:         "AsActionReturnsErrorFunc success returns nil",
		ArrangeInput:  args.Map{"scenario": "success"},
		ExpectedInput: args.Map{"isNil": true},
	},
	{
		Title:         "AsActionReturnsErrorFunc failure returns error",
		ArrangeInput:  args.Map{"scenario": "failure"},
		ExpectedInput: args.Map{"isNil": false},
	},
}

func Test_ResultDelegatingWrapperOf_AsErrFunc_Coverage(t *testing.T) {
	for caseIndex, testCase := range covResultDelegatingOfAsErrFuncTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		scenario, _ := input.GetAsString("scenario")

		var errFunc corefuncs.ActionReturnsErrorFunc

		switch scenario {
		case "success":
			wrapper := corefuncs.NewResultDelegatingWrapper[*string](
				"ok", func(t *string) error { return nil },
			)
			var s string
			errFunc = wrapper.AsActionReturnsErrorFunc(&s)
		case "failure":
			wrapper := corefuncs.NewResultDelegatingWrapper[*string](
				"fail", func(t *string) error { return errors.New("fail") },
			)
			var s string
			errFunc = wrapper.AsActionReturnsErrorFunc(&s)
		}

		err := errFunc()

		// Act
		actual := args.Map{"isNil": err == nil}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// ResultDelegatingFuncWrapper.AsActionFunc
// ==========================================================================

var covResultDelegatingAsActionFuncTestCases = []coretestcases.CaseV1{
	{
		Title:         "AsActionFunc executes without panic on success",
		ArrangeInput:  args.Map{},
		ExpectedInput: args.Map{"called": true},
	},
}

func Test_ResultDelegatingWrapper_AsActionFunc_Coverage(t *testing.T) {
	for caseIndex, testCase := range covResultDelegatingAsActionFuncTestCases {
		// Arrange
		called := false
		wrapper := corefuncs.New.LegacyResultDelegating("test", func(target any) error {
			called = true
			return nil
		})

		actionFunc := wrapper.AsActionFunc("target")
		actionFunc()

		// Act
		actual := args.Map{"called": called}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// ResultDelegatingFuncWrapperOf.AsActionFunc
// ==========================================================================

var covResultDelegatingOfAsActionFuncTestCases = []coretestcases.CaseV1{
	{
		Title:         "ResultDelegatingFuncWrapperOf.AsActionFunc executes",
		ArrangeInput:  args.Map{},
		ExpectedInput: args.Map{"called": true},
	},
}

func Test_ResultDelegatingWrapperOf_AsActionFunc_Coverage(t *testing.T) {
	for caseIndex, testCase := range covResultDelegatingOfAsActionFuncTestCases {
		// Arrange
		called := false
		wrapper := corefuncs.NewResultDelegatingWrapper[*string](
			"test", func(t *string) error {
				called = true
				return nil
			},
		)

		var s string
		actionFunc := wrapper.AsActionFunc(&s)
		actionFunc()

		// Act
		actual := args.Map{"called": called}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// SerializeOutputFuncWrapperOf.AsActionReturnsErrorFunc (error path)
// ==========================================================================

var covSerializeErrPathTestCases = []coretestcases.CaseV1{
	{
		Title:         "SerializeOutputFuncWrapperOf error path returns error",
		ArrangeInput:  args.Map{},
		ExpectedInput: args.Map{"hasError": true},
	},
}

func Test_SerializeWrapper_AsErrFunc_Error_Coverage(t *testing.T) {
	for caseIndex, testCase := range covSerializeErrPathTestCases {
		// Arrange
		wrapper := corefuncs.NewSerializeWrapper[string](
			"fail-serialize",
			func(s string) ([]byte, error) { return nil, errors.New("serialize failed") },
		)

		errFunc := wrapper.AsActionReturnsErrorFunc("test")
		err := errFunc()

		// Act
		actual := args.Map{"hasError": err != nil}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// InOutErrFuncWrapper.AsActionFunc (legacy)
// ==========================================================================

var covLegacyInOutErrAsActionFuncTestCases = []coretestcases.CaseV1{
	{
		Title:         "LegacyInOutErr.AsActionFunc executes without panic",
		ArrangeInput:  args.Map{},
		ExpectedInput: args.Map{"called": true},
	},
}

func Test_LegacyInOutErr_AsActionFunc_Coverage(t *testing.T) {
	for caseIndex, testCase := range covLegacyInOutErrAsActionFuncTestCases {
		// Arrange
		called := false
		wrapper := corefuncs.New.LegacyInOutErr("test", func(input any) (any, error) {
			called = true
			return nil, nil
		})

		actionFunc := wrapper.AsActionFunc("input")
		actionFunc()

		// Act
		actual := args.Map{"called": called}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// NewCreator — LegacyInOutErr, LegacyResultDelegating
// ==========================================================================

var covNewCreatorTestCases = []coretestcases.CaseV1{
	{
		Title:         "New.LegacyInOutErr creates wrapper with name",
		ArrangeInput:  args.Map{"method": "LegacyInOutErr"},
		ExpectedInput: args.Map{"name": "my-legacy"},
	},
	{
		Title:         "New.LegacyResultDelegating creates wrapper with name",
		ArrangeInput:  args.Map{"method": "LegacyResultDelegating"},
		ExpectedInput: args.Map{"name": "my-legacy"},
	},
	{
		Title:         "New.NamedAction creates wrapper with name",
		ArrangeInput:  args.Map{"method": "NamedAction"},
		ExpectedInput: args.Map{"name": "my-legacy"},
	},
}

func Test_NewCreator_Coverage(t *testing.T) {
	for caseIndex, testCase := range covNewCreatorTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		method, _ := input.GetAsString("method")

		var name string

		switch method {
		case "LegacyInOutErr":
			w := corefuncs.New.LegacyInOutErr("my-legacy", func(i any) (any, error) { return nil, nil })
			name = w.Name
		case "LegacyResultDelegating":
			w := corefuncs.New.LegacyResultDelegating("my-legacy", func(any) error { return nil })
			name = w.Name
		case "NamedAction":
			w := corefuncs.New.NamedAction("my-legacy", func(string) {})
			name = w.Name
		}

		// Act
		actual := args.Map{"name": name}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// Generic constructors
// ==========================================================================

var covGenericConstructorTestCases = []coretestcases.CaseV1{
	{
		Title:         "NewInOutErrWrapper has correct name",
		ArrangeInput:  args.Map{"constructor": "InOutErr"},
		ExpectedInput: args.Map{"name": "cov-test"},
	},
	{
		Title:         "NewInOutWrapper has correct name",
		ArrangeInput:  args.Map{"constructor": "InOut"},
		ExpectedInput: args.Map{"name": "cov-test"},
	},
	{
		Title:         "NewInActionErrWrapper has correct name",
		ArrangeInput:  args.Map{"constructor": "InActionErr"},
		ExpectedInput: args.Map{"name": "cov-test"},
	},
	{
		Title:         "NewResultDelegatingWrapper has correct name",
		ArrangeInput:  args.Map{"constructor": "ResultDelegating"},
		ExpectedInput: args.Map{"name": "cov-test"},
	},
	{
		Title:         "NewSerializeWrapper has correct name",
		ArrangeInput:  args.Map{"constructor": "Serialize"},
		ExpectedInput: args.Map{"name": "cov-test"},
	},
}

func Test_GenericConstructors_Coverage(t *testing.T) {
	for caseIndex, testCase := range covGenericConstructorTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		constructor, _ := input.GetAsString("constructor")

		var name string

		switch constructor {
		case "InOutErr":
			w := corefuncs.NewInOutErrWrapper[string, int]("cov-test", func(s string) (int, error) { return 0, nil })
			name = w.Name
		case "InOut":
			w := corefuncs.NewInOutWrapper[string, int]("cov-test", func(s string) int { return 0 })
			name = w.Name
		case "InActionErr":
			w := corefuncs.NewInActionErrWrapper[string]("cov-test", func(s string) error { return nil })
			name = w.Name
		case "ResultDelegating":
			w := corefuncs.NewResultDelegatingWrapper[*string]("cov-test", func(t *string) error { return nil })
			name = w.Name
		case "Serialize":
			w := corefuncs.NewSerializeWrapper[string]("cov-test", func(s string) ([]byte, error) { return nil, nil })
			name = w.Name
		}

		// Act
		actual := args.Map{"name": name}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ==========================================================================
// SerializeOutputFuncWrapperOf.Exec
// ==========================================================================

var covSerializeExecTestCases = []coretestcases.CaseV1{
	{
		Title:         "SerializeOutputFuncWrapperOf.Exec returns bytes",
		ArrangeInput:  args.Map{},
		ExpectedInput: args.Map{"length": fmt.Sprintf("%d", 4)},
	},
}

func Test_SerializeWrapper_Exec_Coverage(t *testing.T) {
	for caseIndex, testCase := range covSerializeExecTestCases {
		// Arrange
		wrapper := corefuncs.NewSerializeWrapper[string](
			"serialize",
			func(s string) ([]byte, error) { return []byte(s), nil },
		)

		bytes, _ := wrapper.Exec("test")

		// Act
		actual := args.Map{"length": fmt.Sprintf("%d", len(bytes))}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
