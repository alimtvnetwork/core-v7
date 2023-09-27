package enumimpl

import (
	"fmt"
	"reflect"

	"gitlab.com/auk-go/core/constants"
)

type differCheckerImpl struct{}

func (it *differCheckerImpl) GetSingleDiffResult(isLeft bool, l, r interface{}) interface{} {
	if isLeft {
		return l
	}

	return r
}

func (it *differCheckerImpl) GetResultOnKeyMissingInRightExistInLeft(lKey string, lVal interface{}) interface{} {
	return lVal
}
func (it *differCheckerImpl) IsEqual(isRegardlessType bool, l, r interface{}) bool {
	if isRegardlessType {
		leftString := fmt.Sprintf(
			constants.SprintPropertyNameValueFormat,
			l)
		rightString := fmt.Sprintf(
			constants.SprintPropertyNameValueFormat,
			r)

		return leftString == rightString
	}

	return reflect.DeepEqual(l, r)
}

func (it differCheckerImpl) AsDifferChecker() DifferChecker {
	return &it
}
