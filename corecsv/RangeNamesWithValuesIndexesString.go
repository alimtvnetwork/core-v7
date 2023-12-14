package corecsv

import "strings"

// RangeNamesWithValuesIndexesString
//
//	Returns a new slice to joined
//	string using RangeNamesWithValuesIndexes
//
//	format
//	 - `name[ValueIndex]` + joiner
//	example
//	 - `SomeName[1]` + joiner
func RangeNamesWithValuesIndexesString(
	joiner string,
	rangedItems ...string,
) string {
	return strings.Join(
		RangeNamesWithValuesIndexes(rangedItems...),
		joiner)
}
