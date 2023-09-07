package regexnew

import (
	"gitlab.com/auk-go/core/regconsts"
)

var (
	WhitespaceFinderRegex             = New.Lazy(regconsts.AllWhitespaces)
	HashCommentWithSpaceOptionalRegex = New.Lazy(regconsts.HashCommentWithSpaceOptional)
	WhitespaceOrPipeFinderRegex       = New.Lazy(regconsts.AllWhitespacesOrPipe)
	DollarIdentifierRegex             = New.Lazy(regconsts.EachWordsWithDollarSymbolDefinition)
	PercentIdentifierRegex            = New.Lazy(regconsts.EachWordsWithinPercentSymbolDefinition)
	PrettyNameRegex                   = New.Lazy(regconsts.PrettyName)
	ExactIdFieldMatchingRegex         = New.Lazy(regconsts.ExactIdFieldMatching)
	ExactVersionIdFieldMatchingRegex  = New.Lazy(regconsts.ExactVersionIdFieldMatching)
	UbuntuNameCheckerRegex            = New.Lazy(regconsts.UbuntuNameChecker)
	CentOsNameCheckerRegex            = New.Lazy(regconsts.CentOsNameChecker)
	RedHatNameCheckerRegex            = New.Lazy(regconsts.RedHatNameChecker)
	FirstNumberAnyWhereCheckerRegex   = New.Lazy(regconsts.FirstNumberAnyWhere)
	WindowsVersionNumberCheckerRegex  = FirstNumberAnyWhereCheckerRegex
)
