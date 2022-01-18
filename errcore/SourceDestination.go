package errcore

import "gitlab.com/evatix-go/core/constants"

func SourceDestination(
	isIncludeType bool,
	srcVal,
	destinationVal interface{},
) string {
	return VarTwo(
		isIncludeType,
		constants.SourceLower,
		srcVal,
		constants.DestinationLower,
		destinationVal)
}
