package chmodhelper

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/errcore"
)

type RwxMatchingStatus struct {
	RwxMismatchInfos         []*RwxMismatchInfo
	MissingOrPathsWithIssues []string
	IsAllMatching            bool
	Error                    error
}

func InvalidRwxMatchingStatus(err error) *RwxMatchingStatus {
	return &RwxMatchingStatus{
		RwxMismatchInfos:         []*RwxMismatchInfo{},
		MissingOrPathsWithIssues: []string{},
		IsAllMatching:            false,
		Error:                    err,
	}
}

func EmptyRwxMatchingStatus() *RwxMatchingStatus {
	return &RwxMatchingStatus{
		RwxMismatchInfos:         []*RwxMismatchInfo{},
		MissingOrPathsWithIssues: []string{},
		IsAllMatching:            false,
		Error:                    nil,
	}
}

func (it *RwxMatchingStatus) MissingFilesToString() string {
	return strings.Join(it.MissingOrPathsWithIssues, constants.CommaSpace)
}

func (it *RwxMatchingStatus) CreateErrFinalError() error {
	if it.IsAllMatching && it.Error == nil {
		return nil
	}

	length := len(it.RwxMismatchInfos) +
		constants.Capacity2

	sliceErr := make([]string, 0, length)

	for _, info := range it.RwxMismatchInfos {
		expectingMessage := errcore.ExpectingSimpleNoType(
			info.FilePath,
			info.Expecting,
			info.Actual)

		sliceErr = append(
			sliceErr,
			expectingMessage)
	}

	if it.Error != nil {
		sliceErr = append(
			sliceErr,
			it.Error.Error())
	}

	return errcore.SliceToError(sliceErr)
}
