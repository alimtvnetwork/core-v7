package errcore

import "gitlab.com/evatix-go/core/constants"

func SourceDestinationNoType(
	srcVal,
	destinationVal interface{},
) string {
	return VarTwo(
		false,
		constants.SourceLower,
		srcVal,
		constants.DestinationLower,
		destinationVal)
}
