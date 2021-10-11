package corestr

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/errcore"
)

//goland:noinspection ALL
var (
	StaticJsonError = errcore.EmptyResultCannotMakeJson.
			Error(constants.EmptyString, constants.EmptyString)
	ExpectingLengthForLeftRight      = constants.Two
	LeftRightExpectingLengthMessager = errcore.ExpectingFuture(
		"Expecting length at least",
		ExpectingLengthForLeftRight)
)
