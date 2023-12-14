package reflectinternal

import "gitlab.com/auk-go/core/constants"

const (
	gitHubDotCom                = "github.com"
	gitlabDotCom                = "gitlab.com"
	defaultPointerReduction     = 3
	invalid                     = -1
	funcPrintFormat             = "func %s (...) line (%d):\n\r\tFile %s:%d"
	shortStringFormat           = "%s (%d) -> %s:%d"
	fileWithLineFormat          = "%s:%d"
	defaultStackCount           = 12
	defaultMaxLevelOfReflection = constants.DefaultMaxLevel

	defaultInternalSkip = constants.One
	SkipNone            = constants.Zero
	Skip1               = constants.One
	Skip2               = constants.Two
	Skip3               = constants.Three
	Skip4               = constants.Four
	Skip5               = constants.Five
)
