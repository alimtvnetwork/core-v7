package errcore

import (
	"errors"

	"gitlab.com/evatix-go/core/constants"
)

func SourceDestinationErr(
	isIncludeType bool,
	srcVal,
	destinationVal interface{},
) error {
	message := VarTwo(
		isIncludeType,
		constants.SourceLower,
		srcVal,
		constants.DestinationLower,
		destinationVal)

	return errors.New(message)
}
