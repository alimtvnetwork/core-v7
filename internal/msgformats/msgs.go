package msgformats

const (
	// LogFormat Contains name-value using %+v, %v for only value.
	//
	// Expectations : %+v
	// Actual: %+v
	LogFormat = "\n ====================================" +
		"Actual vs IsMatchesExpectation " +
		"====================================\n" +
		"\tExpectations : %+v\n" +
		"\tActual: %+v"
	PrintValuesFormat = "\nHeader:%s\n" +
		"\tType:%T\n" +
		"\tValue:%s\n"

	// QuickIndexInputActualExpectedMessageFormat
	//
	// Index, Input, Actual, Expected
	QuickIndexInputActualExpectedMessageFormat = "----------------------\n" +
		"%d )  When: %s,\n" +
		"   Actual: %s,\n" +
		" Expected: %s"

	// QuickIndexTitleInputActualExpectedMessageFormat
	//
	// Index, Title, Input, Actual, Expected
	QuickIndexTitleInputActualExpectedMessageFormat = "----------------------\n" +
		"%d )  Title:%#v\n" +
		"      Input:`%#v` ,\n\n" +
		"  Actual:\n  `%#v` ,\n\n" +
		"Expected:\n  `%#v`"

	QuickLinesFormat = "----------------------\n" +
		"%d )  Title: %s\n\n" +
		"=================>\n" +
		"%d ) Actual: %s\n" +
		"=================>\n" +
		"[]string{\n" +
		"%s\n},\n" +
		"=================>\n" +
		"%d ) Expected: %s\n" +
		"=================>\n" +
		"%s\n\n"

	PrintWhenActualAndExpectedProcessedFormat = "" +
		"\n%d )" +
		"   When: %#v\n  " +
		"    Func:`%#v` ,\n\n  " +
		"  Actual:\n  `%#v`,\n\n  " +
		"Expected:\n  `%#v`,\n\n\n  " +
		"  Actual-Processed:\n   `%#v`,\n\n" +
		"Expected-Processed:\n   `%#v`,\n\n" +
		"    TestCase:\n   %#v ,\n  "

	PrintActualAndExpectedProcessedFormat = "----------------------" +
		"\n\n%d )\t" +
		"  Actual:`%#v` ,\n\t\t" +
		"Expected:`%#v`\n\t\t" +
		"  Actual-Processed:`%#v` ,\n\t\t" +
		"Expected-Processed:`%#v`\n"

	SearchTermExpectedFormat = `Expecting (left) TextValidator %s ~= %s search term (right), method %s`

	PrintHeaderForSearchWithActualAndExpectedProcessedFormat = "" +
		"\n" +
		"%d )   Header: `%s`\n" +
		"----- Method: `%#v`, Line Index: %d\n\n" +
		"--------------- Actual:\n`%#v`\n\n" +
		"--- Expected or Search:\n`%#v`\n\n" +
		"Additional: `%v`"

	PrintHeaderForSearchWithActualAndExpectedProcessedWithoutAdditionalFormat = "" +
		"\n" +
		"%d ) Header: `%s`\n  " +
		"Expectation: `%s`, Line Index: %d\n  " +
		"     Actual: `%#v`\n\n" +
		"   Expected: `%#v`\n\n"

	PrintHeaderForSearchActualAndExpectedProcessedSimpleFormat = "%d )\t" +
		"ExpectationLines failed: Failed match method [%#v], Index : [%#v]\n  " +
		"   Actual-Processed: `%#v`\n  " +
		" Expected-Processed: `%#v`"

	PrintSearchLineNumberDidntMatchFormat = "----------------------" +
		"\n%d )\t" +
		"Line Number Failed to match: (left) Validator Line Number Expect [%d] != [%d] Actual Content Line Number \n  " +
		"        TextValidator:`%#v`\n  " +
		"           SearchTerm:`%#v`\n  " +
		"Line Number Expecting:`%#v`\n  " +
		" Line Number Received:`%#v`\n  " +
		"           Additional:`%#v`"

	SimpleGherkinsFormat = "----------------------" +
		"\n%d )\t" +
		"Feature: `%#v` , Index: [%d]\n  " +
		"  Given: `%#v`\n  " +
		"  When: `%#v`\n  " +
		"  Then: `%#v`\n  "

	SimpleGherkinsWithExpectationFormat = "----------------------" +
		"\n%d )\t" +
		"   Feature: `%#v` , Index: [%d]\n  " +
		"     Given: `%#v`\n  " +
		"      When: `%#v`\n  " +
		"      Then: `%#v`\n  " +
		"    Actual: `%#v`\n  " +
		"  Expected: `%#v`\n  "
	SimpleGherkinsExpectationFormat = "" +
		"    Actual: `%#v`\n  " +
		"  Expected: `%#v`\n  "

	TextValidatorSingleLineFormat = "" +
		"Search Input: [`%s`], " +
		"CompareMethod: [`%s`], " +
		"IsTrimCompare: [`%#v`], " +
		"IsSplitByWhitespace: [`%#v`], " +
		"IsUniqueWordOnly: [`%#v`], " +
		"IsNonEmptyWhitespace: [`%#v`], " +
		"IsSortStringsBySpace: [`%#v`]"

	TextValidatorMultiLineFormat = "" +
		"        Search Input: [`%s`],\n " +
		"              CompareMethod: [`%s`],\n " +
		"       IsTrimCompare: [`%#v`],\n " +
		" IsSplitByWhitespace: [`%#v`],\n " +
		"    IsUniqueWordOnly: [`%#v`],\n " +
		"IsNonEmptyWhitespace: [`%#v`],\n " +
		"IsSortStringsBySpace: [`%#v`]"

	MsgHeaderFormat = "\n============================>\n" +
		"`%#v`" +
		"\n============================>\n\n"
	EndingDashes              = "============================"
	MsgHeaderPlusEndingFormat = "\n============================>\n" +
		"%s" +
		"\n============================>\n" +
		"%s" +
		"\n============================>"

	LinePrinterFormat = "%#v,"
)
