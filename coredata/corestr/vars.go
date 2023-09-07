package corestr

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/errcore"
)

//goland:noinspection ALL
var (
	New             = &newCreator{}
	Empty           = &emptyCreator{}
	StaticJsonError = errcore.EmptyResultCannotMakeJsonType.
			Error(constants.EmptyString, constants.EmptyString)
	ExpectingLengthForLeftRight      = constants.Two
	LeftRightExpectingLengthMessager = errcore.ExpectingFuture(
		"Expecting length at least",
		ExpectingLengthForLeftRight)
)
