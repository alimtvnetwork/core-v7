package corevalidator

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/enums/stringcompareas"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/internal/strutilinternal"
)

// SliceValidator
//
// Use this only for one time verification only.
//
// If IsUsedAlready, don't mutate the ActualLines or ExpectedLines
// it will not work.
type SliceValidator struct {
	Condition
	CompareAs stringcompareas.Variant
	// ActualLines considered to be actual
	// ExpectedLines considered to be expected
	ActualLines, ExpectedLines []string
	comparingValidators        *TextValidators // lazy
}

func NewSliceValidatorUsingErr(
	errActual error,
	compareLinesContentAsExpected string,
	isTrimLineCompare,
	isNonEmptyWhitespace,
	isSortStringsBySpace bool,
	compareAs stringcompareas.Variant,
) *SliceValidator {
	inputLines := errcore.ErrorToSplitLines(errActual)
	compareLines := strings.Split(
		compareLinesContentAsExpected,
		constants.NewLineUnix,
	)

	return &SliceValidator{
		ActualLines:   inputLines,
		ExpectedLines: compareLines,
		Condition: Condition{
			IsTrimCompare:        isTrimLineCompare,
			IsNonEmptyWhitespace: isNonEmptyWhitespace,
			IsSortStringsBySpace: isSortStringsBySpace,
		},
		CompareAs:           compareAs,
		comparingValidators: nil,
	}
}

func NewSliceValidatorUsingAny(
	anyValActual interface{},
	compareLinesContentExpected string,
	isTrimLineCompare,
	isNonEmptyWhitespace,
	isSortStringsBySpace bool,
	compareAs stringcompareas.Variant,
) *SliceValidator {
	anyToString := strutilinternal.AnyToString(anyValActual)
	splitLines := strings.Split(anyToString, constants.NewLineUnix)
	compareLines := strings.Split(
		compareLinesContentExpected,
		constants.NewLineUnix,
	)

	return &SliceValidator{
		ActualLines:   splitLines,
		ExpectedLines: compareLines,
		Condition: Condition{
			IsTrimCompare:        isTrimLineCompare,
			IsNonEmptyWhitespace: isNonEmptyWhitespace,
			IsSortStringsBySpace: isSortStringsBySpace,
		},
		CompareAs:           compareAs,
		comparingValidators: nil,
	}
}

func (it *SliceValidator) IsUsedAlready() bool {
	if it == nil {
		return false
	}

	return it.comparingValidators != nil
}

func (it *SliceValidator) ActualLinesLength() int {
	if it == nil {
		return 0
	}

	return len(it.ActualLines)
}

func (it *SliceValidator) MethodName() string {
	return it.CompareAs.Name()
}

func (it *SliceValidator) IsValidOtherLines(
	isCaseSensitive bool,
	otherActualLines []string,
) bool {
	return it.
		isValidLines(
			isCaseSensitive,
			otherActualLines,
		)
}

func (it *SliceValidator) SetActual(
	actual []string,
) *SliceValidator {
	it.ActualLines = actual

	return it
}

func (it *SliceValidator) SetActualVsExpected(
	actual, expected []string,
) *SliceValidator {
	it.ActualLines = actual
	it.ExpectedLines = expected

	return it
}

func (it *SliceValidator) ActualLinesString() string {
	if it == nil {
		return constants.EmptyString
	}

	return errcore.StringLinesToQuoteLinesToSingle(
		it.ActualLines,
	)
}

func (it *SliceValidator) ExpectingLinesString() string {
	if it == nil {
		return constants.EmptyString
	}

	return errcore.StringLinesToQuoteLinesToSingle(
		it.ExpectedLines,
	)
}

func (it *SliceValidator) ExpectingLinesLength() int {
	if it == nil {
		return 0
	}

	return len(it.ExpectedLines)
}

func (it *SliceValidator) ComparingValidators() *TextValidators {
	if it.comparingValidators != nil {
		return it.comparingValidators
	}

	validators := NewTextValidators(it.ExpectingLinesLength())

	for _, line := range it.ExpectedLines {
		validators.Add(
			TextValidator{
				Search:    line,
				Condition: it.Condition,
				SearchAs:  it.CompareAs,
			},
		)
	}

	it.comparingValidators = validators

	return it.comparingValidators
}

func (it *SliceValidator) IsValid(isCaseSensitive bool) bool {
	if it == nil {
		return true
	}

	return it.isValidLines(
		isCaseSensitive,
		it.ActualLines,
	)
}

func (it *SliceValidator) isValidLines(
	isCaseSensitive bool,
	lines []string,
) bool {
	if it == nil && lines == nil {
		return true
	}

	if lines == nil && it.ExpectedLines == nil {
		return true
	}

	if lines == nil || it.ExpectedLines == nil {
		return false
	}

	inputLength := len(lines)
	comparingLength := len(it.ExpectedLines)

	if inputLength != comparingLength {
		return false
	}

	validators := it.ComparingValidators()

	for i, validator := range validators.Items {
		isNotMatch := !validator.IsMatch(
			lines[i],
			isCaseSensitive,
		)

		if isNotMatch {
			return false
		}
	}

	return true
}

func (it *SliceValidator) VerifyFirstError(
	parameter *Parameter,
) error {
	if it == nil {
		return nil
	}

	return it.VerifyFirstLengthUptoError(
		parameter,
		it.ExpectingLinesLength(),
	)
}

func (it *SliceValidator) VerifyFirstLengthUptoError(
	params *Parameter,
	lengthUpTo int,
) error {
	if it == nil {
		return nil
	}

	return it.AllVerifyErrorUptoLength(
		true,
		params,
		lengthUpTo,
	)
}

func (it *SliceValidator) AssertAllQuick(
	t *testing.T,
	caseIndex int,
	header string,
	actualElements ...string,
) {
	if it == nil {
		return
	}

	toErr := it.AllVerifyErrorQuick(
		caseIndex,
		header,
		actualElements...,
	)

	convey.Convey(
		header, t, func() {
			convey.So(
				toErr,
				should.BeNil,
			)
		},
	)
}

func (it *SliceValidator) AllVerifyErrorQuick(
	caseIndex int,
	header string,
	actualElements ...string,
) error {
	if it == nil {
		return nil
	}

	var params = Parameter{
		CaseIndex:                  caseIndex,
		Header:                     header,
		IsSkipCompareOnActualEmpty: true,
		IsAttachUserInputs:         true,
		IsCaseSensitive:            true,
	}

	it.SetActual(actualElements)

	return it.AllVerifyErrorUptoLength(
		false,
		&params,
		it.ExpectingLinesLength(),
	)
}

func (it *SliceValidator) AllVerifyError(
	params *Parameter,
) error {
	if it == nil {
		return nil
	}

	return it.AllVerifyErrorUptoLength(
		false,
		params,
		it.ExpectingLinesLength(),
	)
}

func (it *SliceValidator) AllVerifyErrorTestCase(
	caseIndex int,
	header string,
	isCaseSensitive bool,
) error {
	if it == nil {
		return nil
	}

	params := Parameter{
		CaseIndex:                  caseIndex,
		Header:                     header,
		IsSkipCompareOnActualEmpty: false,
		IsAttachUserInputs:         true,
		IsCaseSensitive:            isCaseSensitive,
	}

	err := it.AllVerifyErrorUptoLength(
		false,
		&params,
		it.ExpectingLinesLength(),
	)

	errcore.PrintErrorWithTestIndex(caseIndex, header, err)

	return err
}

// AllVerifyErrorExceptLast
//
// Verify up to the second last item.
func (it *SliceValidator) AllVerifyErrorExceptLast(
	params *Parameter,
) error {
	if it == nil {
		return nil
	}

	return it.AllVerifyErrorUptoLength(
		false,
		params,
		it.ExpectingLinesLength()-1,
	)
}

func (it *SliceValidator) AllVerifyErrorUptoLength(
	isFirstOnly bool,
	params *Parameter,
	lengthUpto int,
) error {
	if it == nil {
		return nil
	}

	if it.isEmptyIgnoreCase(params) {
		return nil
	}

	initialVerifyErr := it.initialVerifyErrorWithMerged(
		params,
		lengthUpto,
	)

	if initialVerifyErr != nil {
		return initialVerifyErr
	}

	lengthErr := it.lengthVerifyError(params, lengthUpto)
	if lengthErr != nil {
		return lengthErr
	}

	validators := it.ComparingValidators()
	var sliceErr []string
	for i, validator := range validators.Items[:lengthUpto] {
		err := validator.VerifySimpleError(
			i,
			params,
			it.ActualLines[i],
		)

		if err != nil {
			sliceErr = append(
				sliceErr,
				err.Error(),
			)
		}

		if isFirstOnly && err != nil {
			break
		}
	}

	if params.IsAttachUserInputs && len(sliceErr) > constants.Zero {
		sliceErr = append(
			sliceErr,
			it.ActualInputWithExpectingMessage(
				params.CaseIndex,
				params.Header,
			),
		)
	}

	return errcore.SliceToError(sliceErr)
}

func (it *SliceValidator) lengthVerifyError(
	params *Parameter,
	lengthUpto int,
) error {
	hasLengthUpto := lengthUpto > constants.InvalidValue
	comparingLength := it.ExpectingLinesLength()

	var comparingLengthError error
	if hasLengthUpto && lengthUpto > comparingLength {
		comparingLengthError = errcore.OutOfRangeLengthType.Error(
			"Asked comparingLength is out of range!",
			comparingLength,
		)
	}

	if comparingLengthError != nil {
		return it.UserInputsMergeWithError(
			params,
			comparingLengthError,
		)
	}

	var inputLengthErr error
	if it.ActualLinesLength() > 0 && comparingLength == 0 {
		inputLengthErr = errcore.LengthIssueType.Error(
			"Input comparison has some text but comparing length is 0! Must set comparing text!",
			comparingLength,
		)
	}

	if inputLengthErr != nil {
		return it.UserInputsMergeWithError(
			params,
			inputLengthErr,
		)
	}

	return nil
}

// initialVerifyError, verifyLengthUpto less than 0 will check actual length
func (it *SliceValidator) initialVerifyError(
	lengthUpto int,
) error {
	if it.ActualLines == nil && it.ExpectedLines == nil {
		return nil
	}

	isAnyNilCase := it.ActualLines == nil ||
		it.ExpectedLines == nil

	if isAnyNilCase {
		return errcore.ExpectingErrorSimpleNoTypeNewLineEnds(
			"ActualLines, ExpectedLines any is nil and other is not.",
			it.ActualLines,
			it.ExpectedLines,
		)
	}

	if !it.isLengthOkay(lengthUpto) {
		return errcore.ExpectingErrorSimpleNoTypeNewLineEnds(
			"ActualLines, ExpectedLines Length is not equal.",
			len(it.ActualLines),
			len(it.ExpectedLines),
		)
	}

	return nil
}

func (it *SliceValidator) isLengthOkay(lengthUpto int) bool {
	inputLength := len(it.ActualLines)
	comparingLength := len(it.ExpectedLines)
	isLengthCheckUpto := lengthUpto > constants.InvalidValue
	var isMinLengthMeet bool

	if isLengthCheckUpto {
		remainingInputLength := inputLength - lengthUpto
		remainingCompareLength := comparingLength - lengthUpto

		isMinLengthMeet = remainingInputLength == remainingCompareLength
	}

	isLengthOkay := isMinLengthMeet ||
		inputLength == comparingLength

	return isLengthOkay
}

func (it *SliceValidator) initialVerifyErrorWithMerged(
	params *Parameter,
	lengthUpto int,
) error {
	initialVerifyErr := it.initialVerifyError(
		lengthUpto,
	)

	if initialVerifyErr != nil {
		return it.UserInputsMergeWithError(
			params,
			initialVerifyErr,
		)
	}

	return nil
}

func (it *SliceValidator) ActualInputWithExpectingMessage(caseIndex int, header string) string {
	actualInputMessage := it.ActualInputMessage(
		caseIndex,
		header,
	)
	userExpectingMessage := it.UserExpectingMessage(
		caseIndex,
		header,
	)
	finalMessage := actualInputMessage + constants.NewLineUnix + userExpectingMessage

	return finalMessage
}

func (it *SliceValidator) ActualInputMessage(
	caseIndex int,
	header string,
) string {
	finalHeader := fmt.Sprintf(
		actualUserInputsV2MessageFormat,
		caseIndex,
		header,
	)

	return errcore.MsgHeaderPlusEnding(
		finalHeader,
		it.ActualLinesString(),
	)
}

func (it *SliceValidator) UserExpectingMessage(
	caseIndex int,
	header string,
) string {
	finalHeader := fmt.Sprintf(
		expectingLinesV2MessageFormat,
		caseIndex,
		header,
	)

	return errcore.MsgHeaderPlusEnding(
		finalHeader,
		it.ExpectingLinesString(),
	)
}

func (it *SliceValidator) isEmptyIgnoreCase(
	params *Parameter,
) bool {
	return params.IsSkipCompareOnActualEmpty &&
		len(it.ActualLines) == 0
}

// UserInputsMergeWithError
//
//   - Returns a combine error of actual and expecting inputs.
//   - If all validation successful then no error.
func (it *SliceValidator) UserInputsMergeWithError(
	parameter *Parameter,
	err error,
) error {
	if !parameter.IsAttachUserInputs {
		return err
	}

	toStr := it.ActualInputWithExpectingMessage(
		parameter.CaseIndex,
		parameter.Header,
	)

	if err == nil && len(toStr) == 0 {
		return nil
	}

	if err == nil && len(toStr) >= 0 {
		return errors.New(toStr)
	}

	msg := err.Error() + toStr

	return errors.New(msg)
}

func (it *SliceValidator) Dispose() {
	if it == nil {
		return
	}

	it.ActualLines = nil
	it.ExpectedLines = nil
	it.comparingValidators.Dispose()
	it.comparingValidators = nil
}
