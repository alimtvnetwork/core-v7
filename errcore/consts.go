package errcore

import "gitlab.com/evatix-go/core/constants"

const (
	ReferenceStart                           = "Reference(s) ("
	ReferenceEnd                             = ")"
	ReferenceFormat                          = " Ref(s) { \"%v\" }"
	rangeWithRangeFormat                     = "Range must be in between %v and %v. Ranges must be one of these {%v}"
	rangeWithoutRangeFormat                  = "Range must be in between %v and %v."
	CannotConvertStringToByteForLessThanZero = "Cannot convert string to byte. String cannot be less than 0 for byte."
	CannotConvertStringToByteForMoreThan255  = "Cannot convert string to byte. String is a number " +
		"but larger than byte size. At max it could be 255."
	CannotConvertStringToByte = "Cannot convert string to byte."
	// expectingMessageFormat "%s - expecting (type:[%T]) : [\"%v\"], but received or
	// actual (type:[%T]) : [\"%v\"]"
	expectingMessageFormat = "%s - expecting (type:[%T]) : [\"%v\"], but received " +
		"or actual (type:[%T]) : [\"%v\"]"
	expectingSimpleMessageFormat                  = "%s - Expect (type:\"%T\")[\"%v\"] != [\"%v\"](type:\"%T\") Actual"
	expectingSimpleNoTypeMessageFormat            = "%s - Expect [\"%v\"] != [\"%v\"] Actual"
	expectingNotMatchingSimpleNoTypeMessageFormat = "%s - Expect [\"%v\"] Not Matching [\"%v\"] Actual"
	var2Format                                    = "(%s, %s) = (%v, %v)"
	var2WithTypeFormat                            = "(%s [t:%T], %s[t:%T]) = (%v, %v)"
	var3Format                                    = "(%s, %s, %s) = (%v, %v, %v)"
	keyValFormat                                  = constants.KeyValShortFormat
	var3WithTypeFormat                            = "(%s [t:%T], %s[t:%T], %s[t:%T]) = (%v, %v, %v)"
	messageVar2Format                             = "%s (%s, %s) = (%v, %v)"
	messageVar3Format                             = "%s (%s, %s, %s) = (%v, %v, %v)"
	messageMapFormat                              = constants.MessageReferenceWrapFormat
	messageWithTracesWithoutRefFormat             = "%s \n%s"
	refsWithoutQuotation                          = " Ref (s) { %v }"
	messageWithRefWithoutQuoteFormat              = "%s" + refsWithoutQuotation
	messageWithOtherMsgWithRefWithoutQuoteFormat  = "%s " + messageWithRefWithoutQuoteFormat
	messageWithTracesRefFormat                    = messageWithTracesWithoutRefFormat + refsWithoutQuotation
	PrefixStackTrace                              = constants.Hyphen + constants.Space
	PrefixStackTraceNewLine                       = constants.DefaultLine + PrefixStackTrace
	NewLineCodeStacksHeader                       = "\nCode Stacks :\n"
	CodeStacksHeaderNewLine                       = "Code Stacks :\n"
)
