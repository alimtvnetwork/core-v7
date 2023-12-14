package errcore

// StringLinesToQuoteLines
//
// Each line will be wrapped with "\"%s\", quotation and comma
func StringLinesToQuoteLines(lines []string) []string {
	if len(lines) == 0 {
		return []string{}
	}

	return LinesToDoubleQuoteLinesWithTabs(
		0,
		lines,
	)
}
