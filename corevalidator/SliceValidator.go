package corevalidator

import (
	"errors"
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/enums/stringcompareas"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/core/internal/utilstringinternal"
)

type SliceValidator struct {
	ValidatorCoreCondition
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
		constants.NewLineUnix)

	return &SliceValidator{
		ActualLines:   inputLines,
		ExpectedLines: compareLines,
		ValidatorCoreCondition: ValidatorCoreCondition{
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
	anyToString := utilstringinternal.AnyToString(anyValActual)
	splitLines := strings.Split(anyToString, constants.NewLineUnix)
	compareLines := strings.Split(
		compareLinesContentExpected,
		constants.NewLineUnix)

	return &SliceValidator{
		ActualLines:   splitLines,
		ExpectedLines: compareLines,
		ValidatorCoreCondition: ValidatorCoreCondition{
			IsTrimCompare:        isTrimLineCompare,
			IsNonEmptyWhitespace: isNonEmptyWhitespace,
			IsSortStringsBySpace: isSortStringsBySpace,
		},
		CompareAs:           compareAs,
		comparingValidators: nil,
	}
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
			otherActualLines)
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
		it.ActualLines)
}

func (it *SliceValidator) ExpectingLinesString() string {
	if it == nil {
		return constants.EmptyString
	}

	return errcore.StringLinesToQuoteLinesToSingle(
		it.ExpectedLines)
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
		validators.Add(TextValidator{
			Search:                 line,
			ValidatorCoreCondition: it.ValidatorCoreCondition,
			SearchAs:               it.CompareAs,
		})
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
		it.ActualLines)
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
			isCaseSensitive)

		if isNotMatch {
			return false
		}
	}

	return true
}

func (it *SliceValidator) VerifyFirstError(
	paramsBase *ValidatorParamsBase,
) error {
	if it == nil {
		return nil
	}

	return it.VerifyFirstLengthUptoError(
		paramsBase,
		it.ExpectingLinesLength())
}

func (it *SliceValidator) VerifyFirstLengthUptoError(
	params *ValidatorParamsBase,
	lengthUpTo int,
) error {
	if it == nil {
		return nil
	}

	return it.AllVerifyErrorUptoLength(
		true,
		params,
		lengthUpTo)
}

func (it *SliceValidator) AllVerifyError(
	params *ValidatorParamsBase,
) error {
	if it == nil {
		return nil
	}

	return it.AllVerifyErrorUptoLength(
		false,
		params,
		it.ExpectingLinesLength())
}

func (it *SliceValidator) AllVerifyErrorTestCase(
	caseIndex int,
	isCaseSensitive bool,
) error {
	if it == nil {
		return nil
	}

	params := ValidatorParamsBase{
		CaseIndex:                         caseIndex,
		IsIgnoreCompareOnActualInputEmpty: false,
		IsAttachUserInputs:                true,
		IsCaseSensitive:                   isCaseSensitive,
	}

	err := it.AllVerifyErrorUptoLength(
		false,
		&params,
		it.ExpectingLinesLength())

	errcore.ErrPrintWithTestIndex(caseIndex, err)

	return err
}

func (it *SliceValidator) AllVerifyErrorExceptLast(
	params *ValidatorParamsBase,
) error {
	if it == nil {
		return nil
	}

	return it.AllVerifyErrorUptoLength(
		false,
		params,
		it.ExpectingLinesLength()-1)
}

func (it *SliceValidator) AllVerifyErrorUptoLength(
	isFirstOnly bool,
	params *ValidatorParamsBase,
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
		lengthUpto)

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
			it.ActualLines[i])

		if err != nil {
			sliceErr = append(
				sliceErr,
				err.Error())
		}

		if isFirstOnly && err != nil {
			break
		}
	}

	if params.IsAttachUserInputs && len(sliceErr) > constants.Zero {
		sliceErr = append(
			sliceErr,
			it.ActualInputWithExpectingMessage(params.Header))
	}

	return errcore.SliceToError(sliceErr)
}

func (it *SliceValidator) lengthVerifyError(
	params *ValidatorParamsBase,
	lengthUpto int,
) error {
	hasLengthUpto := lengthUpto > constants.InvalidValue
	comparingLength := it.ExpectingLinesLength()

	var comparingLengthError error
	if hasLengthUpto && lengthUpto > comparingLength {
		comparingLengthError = errcore.OutOfRangeLength.Error(
			"Asked comparingLength is out of range!",
			comparingLength,
		)
	}

	if comparingLengthError != nil {
		return it.UserInputsMergeWithError(
			params,
			comparingLengthError)
	}

	var inputLengthErr error
	if it.ActualLinesLength() > 0 && comparingLength == 0 {
		inputLengthErr = errcore.LengthIssue.Error(
			"Input comparison has some text but comparing length is 0! Must set comparing text!",
			comparingLength,
		)
	}

	if inputLengthErr != nil {
		return it.UserInputsMergeWithError(
			params,
			inputLengthErr)
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
		return errcore.ExpectingErrorSimpleNoType(
			"ActualLines, ExpectedLines any is nil and other is not.",
			it.ActualLines,
			it.ExpectedLines,
		)
	}

	if !it.isLengthOkay(lengthUpto) {
		return errcore.ExpectingErrorSimpleNoType(
			"ActualLines, ExpectedLines Length is not equal",
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
	params *ValidatorParamsBase,
	lengthUpto int,
) error {
	initialVerifyErr := it.initialVerifyError(
		lengthUpto)

	if initialVerifyErr != nil {
		return it.UserInputsMergeWithError(
			params,
			initialVerifyErr)
	}

	return nil
}

func (it *SliceValidator) ActualInputWithExpectingMessage(header string) string {
	return it.ActualInputMessage(header) +
		it.UserExpectingMessage()
}

func (it *SliceValidator) ActualInputMessage(header string) string {
	return errcore.MsgHeaderPlusEnding(
		actualUserInputsMessage+header,
		it.ActualLinesString())
}

func (it *SliceValidator) UserExpectingMessage() string {
	return errcore.MsgHeaderPlusEnding(
		expectingLinesMessage,
		it.ExpectingLinesString())
}

func (it *SliceValidator) isEmptyIgnoreCase(
	params *ValidatorParamsBase,
) bool {
	return params.IsIgnoreCompareOnActualInputEmpty &&
		len(it.ActualLines) == 0
}

func (it *SliceValidator) UserInputsMergeWithError(
	paramsBase *ValidatorParamsBase,
	err error,
) error {
	if !paramsBase.IsAttachUserInputs {
		return err
	}

	if err == nil {
		return errors.New(it.ActualInputWithExpectingMessage(paramsBase.Header))
	}

	msg := err.Error() +
		it.ActualInputWithExpectingMessage(paramsBase.Header)

	return errors.New(msg)
}
