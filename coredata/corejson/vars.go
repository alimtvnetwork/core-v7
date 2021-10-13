package corejson

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/errcore"
)

var (
	StaticJsonError = errcore.EmptyResultCannotMakeJsonType.
		Error(constants.EmptyString, constants.EmptyString)
)
