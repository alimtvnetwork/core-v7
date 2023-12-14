package coreversiontests

const (
	defaultCreationFmt          = "%d : %s (compact: %s, display: %s)"
	defaultInvalidV1CreationFmt = "%d : invalid - %s (raw: %s)"
	defaultInvalidV2CreationFmt = "%d : invalid - %s "

	comparisonFmt        = "%d : [ %s ] %s [ %s ] | Expect: %s - %t"
	comparisonMethodFmt  = "    %d : .%s(%d, %d) -> %t | %t - expected"
	comparisonMethod3Fmt = "    %d : .%s(%d, %d, %d) -> %t | %t - expected"
	jsonFmt              = "%d : [ %s ] - %s"
)
