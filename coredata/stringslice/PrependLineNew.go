package stringslice

func PrependLineNew(
	line string,
	slice []string,
) []string {
	mergedNew := PrependNew(slice, line)

	return *mergedNew
}
