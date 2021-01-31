package corejson

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/msgtype"
)

var (
	StaticJsonError = msgtype.EmptyResultCannotMakeJson.
		Error(constants.EmptyString, constants.EmptyString)
)
