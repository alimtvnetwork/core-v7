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
		"%d )\tWhen:%#v\n\t\t" +
		"  Actual:`%#v` ,\n\t\t" +
		"Expected:`%#v`"

	// QuickIndexTitleInputActualExpectedMessageFormat
	//
	// Index, Title, Input, Actual, Expected
	QuickIndexTitleInputActualExpectedMessageFormat = "----------------------\n" +
		"%d )\tTitle:%#v\n\t\t" +
		"   Input:`%#v` ,\n\t\t" +
		"  Actual:`%#v` ,\n\t\t" +
		"Expected:`%#v`"

	PrintWhenActualAndExpectedProcessedFormat = "" +
		"\n%d )" +
		"   When: %#v\n  " +
		"    Func:`%#v` ,\n  " +
		"  Actual:`%#v` ,\n  " +
		"Expected:`%#v`\n  " +
		"  Actual-Processed:`%#v`,\n  " +
		"Expected-Processed:`%#v`,\n  " +
		"    TestCase:`%#v` ,\n  "

	PrintActualAndExpectedProcessedFormat = "----------------------" +
		"\n%d )\t" +
		"  Actual:`%#v` ,\n\t\t" +
		"Expected:`%#v`\n\t\t" +
		"  Actual-Processed:`%#v` ,\n\t\t" +
		"Expected-Processed:`%#v`"

	SearchTermExpectedFormat = `Expecting (left) TextValidator %s ~= %s search term (right), method %s`

	PrintHeaderForSearchWithActualAndExpectedProcessedFormat = "" +
		"%d )\t" +
		"  Expectation failed: Using CompareMethod `%#v`, Line Index: %d\n  " +
		"   Content-Processed:`%#v`\n  " +
		"SearchTerm-Processed:`%#v`\n  " +
		"          Additional:`%#v`"

	PrintHeaderForSearchWithActualAndExpectedProcessedWithoutAdditionalFormat = "" +
		"%d )\t" +
		"         Expectation:`%s`, Line Index: %d\n  " +
		"   Content-Processed:`%#v`\n  " +
		"SearchTerm-Processed:`%#v`\n  "

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
