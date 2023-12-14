package constants

const (
	SprintValueFormat                            = "%v"
	SprintValueDoubleQuotationFormat             = "\"%s\""
	SprintNumberFormat                           = "%d"
	SprintFullPropertyNameValueFormat            = "%#v"
	SprintPropertyNameValueFormat                = "%+v"
	SprintPropertyValueWithTypeFormat            = "%+v (%T)"
	SprintTypeFormat                             = "%T"
	SprintTypeInParenthesisFormat                = "(type : %T)"
	SprintNilValueTypeInParenthesisFormat        = "<nil> (type : %T)"
	SprintValueWithTypeFormat                    = "%v " + SprintTypeInParenthesisFormat
	SprintDoubleQuoteFormat                      = "%q"
	SprintStartStringEndCharFormat               = "%c%s%c"
	SprintSingleQuoteFormat                      = "'%s'"
	SprintStringFormat                           = "%s"
	SprintThirdBracketQuoteFormat                = "[\"%v\"]"
	KeyValuePariSimpleFormat                     = "{ Key (Type - %T): %v} - { Value (Type - %T) : %v  }"
	SprintFormatNumberWithColon                  = "%d:%d"
	SprintFormatAnyValueWithColon                = "%v:%v"
	TitleValueFormat                             = "%s : %v"
	CurlyTitleWrapFormat                         = "%s: {%s}"        // Title, Value
	QuotationTitleWrapFormat                     = "%v: \"%v\""      // Title, Value
	QuotationTitleMetaWrapFormat                 = "%s: \"%s\" (%v)" // Title, Value, Meta
	CurlyTitleMetaWrapFormat                     = "%s: {%s} (%v)"   // Title, Value, Meta
	SquareTitleWrapFormat                        = "%s: [%s]"        // Title, Value
	SquareTitleMetaWrapFormat                    = "%s: [%s] (%v)"   // Title, Value, Meta
	SprintFormatAnyValueWithComma                = "%v,%v"
	SprintFormatWithNewLine                      = "%v\n%v"
	SprintFormatAnyValueWithPipe                 = "%v|%v"
	SprintFormatAnyNameValueWithColon            = "%#v:%#v"
	SprintFormatAnyNameValueWithPipe             = "%#v|%#v"
	SprintFormatNumberWithHyphen                 = "%d-%d"
	SprintFormatNumberWithPipe                   = "%d|%d"
	ThreeValueNewLineJoin                        = "%v\n%v\n%v"
	ThreeValueNewLineSpaceJoin                   = " %v\n %v\n %v"
	BracketWrapFormat                            = "[%s]"
	BracketQuotationWrapFormat                   = "[\"%s\"]"
	CurlyWrapFormat                              = "{%s}"
	SquareWrapFormat                             = "[%s]"
	ParenthesisWrapFormat                        = "(%s)"
	CurlyQuotationWrapFormat                     = "{\"%s\"}"
	ParenthesisQuotationWrap                     = "(\"%s\")"
	ReferenceWrapFormat                          = "Ref (s) { %v }"
	MessageReferenceWrapFormat                   = "%s Ref (s) { %v }"
	StringWithBracketWrapNumberFormat            = "%s[%d]"
	DoubleQuoteStringWithBracketWrapNumberFormat = "\"%s\"[%d]"
	SpaceHyphenAngelBracketSpaceRefWrapFormat    = " -> Ref(%v)"
	ValueWithDoubleQuoteFormat                   = "\"%v\""
	ValueWithSingleQuoteFormat                   = "'%v'"
	StringWithDoubleQuoteFormat                  = "\"%s\""
	StringWithSingleQuoteFormat                  = "'%s'"
	StringTypeFormat                             = "'%s - %T'"
	TypeWithSingleQuoteFormat                    = "'%T'"
	TypeWithDoubleQuoteFormat                    = "\"%T\""
	MessageWrapMessageFormat                     = "%s (%s)"
	FromToFormat                                 = "{From : %q, To: %q}"            // From, To name
	SourceDestinationFormat                      = "{Source : %q, Destination: %q}" // source, destination
	RenameFormat                                 = "{Existing : %q, New: %q}"       // existing, new
	ValueWrapValueFormat                         = "%v (%v)"
	StringWrapValueFormat                        = "%s (%s)"
	FilePathEmpty                                = "File path was empty(\"\")."
	EnumOnlySupportedFormat                      = "enum: %T, " +
		"not supported (\"%s\") | only supported { %s }" // enumSelf, enumSelf, csv-support
	EnumOnlySupportedWithMessageFormat = "enum: %T, " +
		"not supported (\"%s\") | %s | only supported { %s }" // enumSelf, enumSelf, message, csv-support
)
