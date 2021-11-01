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
	SprintFormatAnyValueWithComma                = "%v,%v"
	SprintFormatAnyValueWithPipe                 = "%v|%v"
	SprintFormatAnyNameValueWithColon            = "%#v:%#v"
	SprintFormatAnyNameValueWithPipe             = "%#v|%#v"
	SprintFormatNumberWithHyphen                 = "%d-%d"
	SprintFormatNumberWithPipe                   = "%d|%d"
	BracketWrapFormat                            = "[%v]"
	BracketQuotationWrapFormat                   = "[\"%v\"]"
	CurlyWrapFormat                              = "{%v}"
	CurlyQuotationWrapFormat                     = "{\"%v\"}"
	ParenthesisWrapFormat                        = "(%v)"
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
	EnumOnlySupportedFormat                      = "enum: %T, " +
		"not supported (\"%s\") | only supported { %s }" // enumSelf, enumSelf, csv-support
	EnumOnlySupportedWithMessageFormat = "enum: %T, " +
		"not supported (\"%s\") | %s | only supported { %s }" // enumSelf, enumSelf, message, csv-support
)
