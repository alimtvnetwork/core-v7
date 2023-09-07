package corecomparator

import "gitlab.com/auk-go/core/internal/csvinternal"

func RangeNamesCsv() string {
	return csvinternal.RangeNamesWithValuesIndexesCsvString(
		CompareNames[:]...)
}
