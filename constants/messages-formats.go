package constants

const (
	SprintValueFormat                            = "%v"
	SprintValueDoubleQuotationFormat             = "\"%v\""
	SprintNumberFormat                           = "%d"
	SprintFullPropertyNameValueFormat            = "%#v"
	SprintPropertyNameValueFormat                = "%+v"
	SprintTypeFormat                             = "%T"
	SprintTypeInParenthesisFormat                = "(type : %T)"
	SprintNilValueTypeInParenthesisFormat        = "<nil> (type : %T)"
	SprintValueWithTypeFormat                    = "%v " + SprintTypeInParenthesisFormat
	SprintDoubleQuoteFormat                      = "%q"
	SprintSingleQuoteFormat                      = "'%s'"
	SprintStringFormat                           = "%s"
	SprintThirdBracketQuoteFormat                = "[\"%v\"]"
	KeyValuePariSimpleFormat                     = "{ Key (Type - %T): %v} - { Value (Type - %T) : %v  }"
	SprintFormatNumberWithColon                  = "%d:%d"
	SprintFormatAnyValueWithColon                = "%v:%v"
	CurlyTitleWrapFormat                         = "%v: {%v}"        // Title, Value
	QuotationTitleWrapFormat                     = "%v: \"%v\""      // Title, Value
	QuotationTitleMetaWrapFormat                 = "%v: \"%v\" (%v)" // Title, Value, Meta
	CurlyTitleMetaWrapFormat                     = "%v: {%v} (%v)"   // Title, Value, Meta
	SquareTitleWrapFormat                        = "%v: [%v]"        // Title, Value
	SquareTitleMetaWrapFormat                    = "%v: [%v] (%v)"   // Title, Value, Meta
	SprintFormatAnyValueWithComma                = "%v,%v"
	SprintFormatWithNewLine                      = "%v\n%v"
	SprintFormatAnyValueWithPipe                 = "%v|%v"
	SprintFormatAnyNameValueWithColon            = "%#v:%#v"
	SprintFormatAnyNameValueWithPipe             = "%#v|%#v"
	SprintFormatNumberWithHyphen                 = "%d-%d"
	SprintFormatNumberWithPipe                   = "%d|%d"
	ThreeValueNewLineJoin                        = "%v\n%v\n%v"
	ThreeValueNewLineSpaceJoin                   = " %v\n %v\n %v"
	BracketWrapFormat                            = "[%v]"
	BracketQuotationWrapFormat                   = "[\"%v\"]"
	CurlyWrapFormat                              = "{%v}"
	SquareWrapFormat                             = "[%v]"
	ParenthesisWrapFormat                        = "(%v)"
	CurlyQuotationWrapFormat                     = "{\"%v\"}"
	ParenthesisQuotationWrap                     = "(\"%v\")"
	ReferenceWrapFormat                          = "Ref (s) { %v }"
	MessageReferenceWrapFormat                   = "%s Ref (s) { %v }"
	StringWithBracketWrapNumberFormat            = "%s[%d]"
	DoubleQuoteStringWithBracketWrapNumberFormat = "\"%s\"[%d]"
	SpaceHyphenAngelBracketSpaceRefWrapFormat    = " -> Ref(%v)"
	ValueWithDoubleQuoteFormat                   = "\"%v\""
	ValueWithSingleQuoteFormat                   = "'%v'"
	StringWithDoubleQuoteFormat                  = "\"%s\""
	StringWithSingleQuoteFormat                  = "'%s'"
	MessageWrapMessageFormat                     = "%s (%s)"
	ValueWrapValueFormat                         = "%v (%v)"
	FilePathEmpty                                = "File path was empty(\"\")."
	EnumOnlySupportedFormat                      = "enum: %T, " +
		"not supported (\"%s\") | only supported { %s }" // enumSelf, enumSelf, csv-support
	EnumOnlySupportedWithMessageFormat = "enum: %T, " +
		"not supported (\"%s\") | %s | only supported { %s }" // enumSelf, enumSelf, message, csv-support
)
