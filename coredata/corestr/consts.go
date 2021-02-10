package corestr

import "gitlab.com/evatix-go/core/constants"

const (
	charCollectionDefaultCapacity               = constants.ArbitraryCapacity10
	emptyChar                              byte = 0
	defaultEachCollectionCapacity               = constants.ArbitraryCapacity30
	defaultHashsetItems                         = 10
	RegularCollectionEfficiencyLimit            = 1000
	DoubleLimit                                 = RegularCollectionEfficiencyLimit * 3
	commonJoiner                                = "\n\t- "
	NoElements                                  = " {No Element}"
	charCollectionMapLengthFormat               = "\n## Items of `%s`"
	charHashsetMapLengthFormat                  = charCollectionMapLengthFormat
	charCollectionMapSingleItemFormat           = "\n\t%d - `%s` has `%d` items."
	charHashsetMapSingleItemFormat              = charCollectionMapSingleItemFormat
	summaryOfCharCollectionMapLengthFormat      = "# Summary of `%T`, Length (\"%d\") - Sequence `%d`"
	summaryOfCharHashsetMapLengthFormat         = summaryOfCharCollectionMapLengthFormat
	nodesCannotBeNull                           = "node cannot be nil."
	currentNodeCannotBeNull                     = "CurrentNode cannot be nil."
	linkedListCollectionCompareHeaderLeft       = "\n=================================================\n-------------------------------------------------\n# %s - Length : %d\n%s\n\n"
	linkedListCollectionCompareHeaderRight      = "\n-------------------------------------------------\n# %s - Length : %d\n%s\n----------------\nIs Equals: %+v\nLength: %d | %d\n"
)
