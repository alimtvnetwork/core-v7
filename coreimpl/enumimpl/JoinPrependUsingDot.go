package enumimpl

import (
	"gitlab.com/evatix-go/core/constants"
)

func JoinPrependUsingDot(prepend interface{}, anyItems ...interface{}) string {
	return PrependJoin(constants.Dot, prepend, anyItems...)
}
