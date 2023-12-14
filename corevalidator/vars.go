package corevalidator

import "gitlab.com/auk-go/core/enums/stringcompareas"

var (
	DefaultDisabledCoreCondition = Condition{
		IsTrimCompare:        false,
		IsUniqueWordOnly:     false,
		IsNonEmptyWhitespace: false,
		IsSortStringsBySpace: false,
	}

	DefaultTrimCoreCondition = Condition{
		IsTrimCompare: true,
	}

	DefaultSortTrimCoreCondition = Condition{
		IsTrimCompare:        true,
		IsNonEmptyWhitespace: true,
		IsSortStringsBySpace: true,
	}

	DefaultUniqueWordsCoreCondition = Condition{
		IsTrimCompare:        true,
		IsUniqueWordOnly:     true,
		IsNonEmptyWhitespace: true,
		IsSortStringsBySpace: true,
	}

	EmptyValidator = TextValidator{
		Search:    "",
		SearchAs:  stringcompareas.Equal,
		Condition: DefaultTrimCoreCondition,
	}
)
