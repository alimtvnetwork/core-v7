package simplewrap

import "gitlab.com/evatix-go/core/internal/csvinternal"

// TitleSquareCsvMeta
//
//  Usages TitleSquareMeta to give the final output
//
// Example :
//  - Title : [Value] (csv meta items)
func TitleSquareCsvMeta(
	title string,
	value interface{},
	metaCsvItems ...interface{},
) string {
	csvString := csvinternal.AnyItemsToStringDefault(
		metaCsvItems...)

	return TitleSquareMeta(
		title,
		value,
		csvString)
}
