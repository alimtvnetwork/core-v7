package corestr

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/msgtype"
)

//goland:noinspection ALL
var (
	StaticJsonError = msgtype.EmptyResultCannotMakeJson.
		Error(constants.EmptyString, constants.EmptyString)
)
