package regexnewtests

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/regexnew"
)

func Test_CreateMust_Verification(t *testing.T) {
	for caseIndex, testCase := range createMustTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)
		compareInput, _ := input.GetAsString(params.compareInput)

		// Act
		regex := regexnew.CreateMust(pattern)

		actual := args.Map{
			params.regexNotNil: regex != nil,
			params.isMatch:     regex.MatchString(compareInput),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_CreateMust_PanicsOnInvalidPattern(t *testing.T) {
	// Arrange
	pattern := "[invalid"

	// Act & Assert
	defer func() {
		r := recover()
		actual := args.Map{"result": r == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "CreateMust should panic on invalid pattern", actual)
	}()

	regexnew.CreateMust(pattern)
}

func Test_CreateMustLockIf_Verification(t *testing.T) {
	for caseIndex, testCase := range createMustLockIfTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)
		compareInput, _ := input.GetAsString(params.compareInput)
		isLock, _ := input.GetAsBool(params.isLock)

		// Act
		regex := regexnew.CreateMustLockIf(isLock, pattern)

		actual := args.Map{
			params.regexNotNil: regex != nil,
			params.isMatch:     regex.MatchString(compareInput),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_CreateMustLockIf_PanicsOnInvalidPattern(t *testing.T) {
	// Arrange
	pattern := "[invalid"

	// Act & Assert — with lock
	defer func() {
		r := recover()
		actual := args.Map{"result": r == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "CreateMustLockIf should panic on invalid pattern", actual)
	}()

	regexnew.CreateMustLockIf(true, pattern)
}

func Test_CreateLockIf_Verification(t *testing.T) {
	for caseIndex, testCase := range createLockIfTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)
		isLock, _ := input.GetAsBool(params.isLock)

		// Act
		regex, err := regexnew.CreateLockIf(isLock, pattern)

		actual := args.Map{
			params.regexNotNil: regex != nil,
			params.hasError:    err != nil,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_CreateApplicableLock_Verification(t *testing.T) {
	for caseIndex, testCase := range createApplicableLockTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)

		// Act
		regex, err, isApplicable := regexnew.CreateApplicableLock(pattern)

		actual := args.Map{
			params.regexNotNil:  regex != nil,
			params.hasError:     err != nil,
			params.isApplicable: isApplicable,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_NewMustLock_Verification(t *testing.T) {
	for caseIndex, testCase := range newMustLockTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)
		compareInput, _ := input.GetAsString(params.compareInput)

		// Act
		regex := regexnew.NewMustLock(pattern)

		actual := args.Map{
			params.regexNotNil: regex != nil,
			params.isMatch:     regex.MatchString(compareInput),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_NewMustLock_PanicsOnInvalidPattern(t *testing.T) {
	// Arrange
	pattern := "[invalid"

	// Act & Assert
	defer func() {
		r := recover()
		actual := args.Map{"result": r == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "NewMustLock should panic on invalid pattern", actual)
	}()

	regexnew.NewMustLock(pattern)
}

func Test_MatchUsingFuncErrorLock_Verification(t *testing.T) {
	matchFunc := regexnew.RegexValidationFunc(
		func(re *regexp.Regexp, term string) bool {
			return re.MatchString(term)
		},
	)

	for caseIndex, testCase := range matchUsingFuncErrorLockTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)
		compareInput, _ := input.GetAsString(params.compareInput)

		// Act
		err := regexnew.MatchUsingFuncErrorLock(
			pattern,
			compareInput,
			matchFunc,
		)

		actual := args.Map{
			params.isNoError: err == nil,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_MatchUsingCustomizeErrorFuncLock_Verification(t *testing.T) {
	matchFunc := regexnew.RegexValidationFunc(
		func(re *regexp.Regexp, term string) bool {
			return re.MatchString(term)
		},
	)

	customErrFunc := regexnew.CustomizeErr(
		func(regexPattern, matchLookingTerm string, err error, re *regexp.Regexp) error {
			return fmt.Errorf("CUSTOM: pattern %s failed on %s", regexPattern, matchLookingTerm)
		},
	)

	for caseIndex, testCase := range matchUsingCustomizeErrorFuncLockTestCases {
		// Arrange
		input := testCase.Input
		pattern, _ := input.GetAsString(params.pattern)
		compareInput, _ := input.GetAsString(params.compareInput)
		customizerVal, _ := input.GetAsString(params.customizer)

		var errFunc regexnew.CustomizeErr
		if customizerVal == "custom" {
			errFunc = customErrFunc
		}

		// Act
		err := regexnew.MatchUsingCustomizeErrorFuncLock(
			pattern,
			compareInput,
			matchFunc,
			errFunc,
		)

		actual := args.Map{
			params.isNoError:     err == nil,
			params.isCustomError: err != nil && strings.HasPrefix(err.Error(), "CUSTOM:"),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
