package keymk

import "gitlab.com/auk-go/core/constants"

const (
	DefaultJoiner            = constants.Hyphen
	DefaultCap               = constants.Capacity16
	cannotModifyErrorMessage = "cannot insert items to complete or finalized object."
	root                     = "root"
	relative                 = "relative"
	id                       = "id"
	prefix                   = "prefix"
	LegendChainSample        = "{root}-{package}-{group}-{state}-{user}-{item}" // item refers to the final element like id or file name
)
