package enumimpl

import "fmt"

type leftRightDiffCheckerImpl struct{}

func (it leftRightDiffCheckerImpl) GetSingleDiffResult(isLeft bool, l, r interface{}) interface{} {
	diff := DiffLeftRight{
		Left:  l,
		Right: r,
	}

	return diff.DiffString()
}

func (it leftRightDiffCheckerImpl) GetResultOnKeyMissingInRightExistInLeft(lKey string, lVal interface{}) interface{} {
	return fmt.Sprintf("%+v (type:%T) - left - key is missing!", lVal, lVal)
}

func (it leftRightDiffCheckerImpl) IsEqual(isRegardless bool, l, r interface{}) bool {
	return DefaultDiffCheckerImpl.IsEqual(isRegardless, l, r)
}

func (it leftRightDiffCheckerImpl) AsChecker() DifferChecker {
	return &it
}
