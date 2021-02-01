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
	NoElements                                  = " {No Elements}"
	charCollectionMapLengthFormat               = "\n## Items of `%s`"
	charHashsetMapLengthFormat                  = charCollectionMapLengthFormat
	charCollectionMapSingleItemFormat           = "\n\t- `%s` has `%d` items."
	charHashsetMapSingleItemFormat              = charCollectionMapSingleItemFormat
	summaryOfCharCollectionMapLengthFormat      = "# Summary of `%T`, Length (\"%d\")"
	summaryOfCharHashsetMapLengthFormat         = summaryOfCharCollectionMapLengthFormat
)
