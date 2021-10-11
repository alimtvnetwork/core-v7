package corevalidator

import (
	"errors"
	"fmt"
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreutils/stringutil"
	"gitlab.com/evatix-go/core/enums/stringcompareas"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/core/internal/msgformats"
)

type TextValidator struct {
	Search   string `json:"Search,omitempty"`
	SearchAs stringcompareas.Variant
	ValidatorCoreCondition
	searchTextFinalized *string
}

func (it *TextValidator) ToString(isSingleLine bool) string {
	if isSingleLine {
		return fmt.Sprintf(
			msgformats.TextValidatorSingleLineFormat,
			it.Search,
			it.SearchAs.Name(),
			it.IsTrimCompare,
			it.IsSplitByWhitespace(),
			it.IsUniqueWordOnly,
			it.IsNonEmptyWhitespace,
			it.IsSortStringsBySpace)
	}

	return fmt.Sprintf(
		msgformats.TextValidatorMultiLineFormat,
		it.Search,
		it.SearchAs.Name(),
		it.IsTrimCompare,
		it.IsSplitByWhitespace(),
		it.IsUniqueWordOnly,
		it.IsNonEmptyWhitespace,
		it.IsSortStringsBySpace)
}

func (it *TextValidator) String() string {
	return it.ToString(true)
}

func (it *TextValidator) SearchTextFinalized() string {
	return *it.SearchTextFinalizedPtr()
}

func (it *TextValidator) SearchTextFinalizedPtr() *string {
	if it.searchTextFinalized != nil {
		return it.searchTextFinalized
	}

	searchTerm := it.GetCompiledTermBasedOnConditions(
		it.Search,
		it.IsUniqueWordOnly) // for unique word, use lowercase

	it.searchTextFinalized = &searchTerm

	return it.searchTextFinalized
}

func (it *TextValidator) GetCompiledTermBasedOnConditions(
	input string,
	isCaseSensitive bool,
) string {
	searchTerm := input

	if it.IsTrimCompare {
		searchTerm = strings.TrimSpace(searchTerm)
	}

	if it.IsSplitByWhitespace() {
		compiledStringSplits := stringutil.SplitContentsByWhitespaceConditions(
			searchTerm,
			it.IsTrimCompare,
			it.IsNonEmptyWhitespace,
			it.IsSortStringsBySpace,
			it.IsUniqueWordOnly,
			!isCaseSensitive)

		return strings.Join(
			compiledStringSplits,
			constants.Space)
	}

	return searchTerm
}

func (it *TextValidator) IsMatch(
	content string,
	isCaseSensitive bool,
) bool {
	search := it.SearchTextFinalized()
	processedContent := it.GetCompiledTermBasedOnConditions(
		content,
		isCaseSensitive)

	isIgnoreCase := !isCaseSensitive

	return it.SearchAs.IsCompareSuccess(
		isIgnoreCase,
		processedContent,
		search,
	)
}

func (it *TextValidator) IsMatchMany(
	isSkipOnContentsEmpty,
	isCaseSensitive bool,
	contents ...string,
) bool {
	if it == nil {
		return true
	}

	if len(contents) == 0 && isSkipOnContentsEmpty {
		return true
	}

	for _, content := range contents {
		if !it.IsMatch(content, isCaseSensitive) {
			return false
		}
	}

	return true
}

func (it *TextValidator) VerifyDetailError(
	params *ValidatorParamsBase,
	content string,
) error {
	if it == nil {
		return nil
	}

	return it.verifyDetailErrorUsingLineProcessing(
		constants.InvalidValue,
		params,
		content)
}

func (it *TextValidator) verifyDetailErrorUsingLineProcessing(
	lineProcessingIndex int,
	params *ValidatorParamsBase,
	content string,
) error {
	if it == nil {
		return nil
	}

	processedSearch := it.SearchTextFinalized()
	processedContent := it.GetCompiledTermBasedOnConditions(
		content,
		params.IsCaseSensitive)

	isMatch := it.SearchAs.IsCompareSuccess(
		params.IsIgnoreCase(),
		processedContent,
		processedSearch,
	)

	if isMatch {
		return nil
	}

	expectationMethod := it.SearchAs.Name()

	msg := errcore.GetSearchTermExpectationMessage(
		params.CaseIndex,
		expectationMethod,
		lineProcessingIndex,
		processedContent,
		processedSearch,
		it.String())

	return errors.New(msg)
}

func (it *TextValidator) MethodName() string {
	return it.SearchAs.Name()
}

func (it *TextValidator) VerifySimpleError(
	processingIndex int,
	params *ValidatorParamsBase,
	content string,
) error {
	if it == nil {
		return nil
	}

	processedSearch := it.SearchTextFinalized()
	processedContent := it.GetCompiledTermBasedOnConditions(
		content,
		params.IsCaseSensitive)

	isMatch := it.SearchAs.IsCompareSuccess(
		params.IsIgnoreCase(),
		processedContent,
		processedSearch,
	)

	if isMatch {
		return nil
	}

	method := it.SearchAs.Name()

	msg := errcore.GetSearchTermExpectationSimpleMessage(
		params.CaseIndex,
		method,
		processingIndex,
		processedContent,
		processedSearch)

	return errors.New(msg)
}

func (it *TextValidator) VerifyMany(
	isContinueOnError bool,
	params *ValidatorParamsBase,
	contents ...string,
) error {
	if isContinueOnError {
		return it.AllVerifyError(
			params,
			contents...)
	}

	return it.VerifyFirstError(
		params,
		contents...)
}

func (it *TextValidator) VerifyFirstError(
	params *ValidatorParamsBase,
	contents ...string,
) error {
	if it == nil {
		return nil
	}

	length := len(contents)
	if length == 0 && params.IsIgnoreCompareOnActualInputEmpty {
		return nil
	}

	for i, content := range contents {
		err := it.verifyDetailErrorUsingLineProcessing(
			i,
			params,
			content,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func (it *TextValidator) AllVerifyError(
	params *ValidatorParamsBase,
	contents ...string,
) error {
	if it == nil {
		return nil
	}

	length := len(contents)
	if length == 0 && params.IsIgnoreCompareOnActualInputEmpty {
		return nil
	}

	var sliceErr []string

	for i, content := range contents {
		err := it.verifyDetailErrorUsingLineProcessing(
			i,
			params,
			content,
		)

		if err != nil {
			sliceErr = append(
				sliceErr,
				err.Error())
		}
	}

	return errcore.SliceToError(
		sliceErr)
}
