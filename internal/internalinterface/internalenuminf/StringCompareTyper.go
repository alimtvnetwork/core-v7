package internalenuminf

type StringCompareTyper interface {
	BasicEnumer
	ByteValuePlusEqualer
	InternalByteEnumer

	IsEqual() bool
	IsStartsWith() bool
	IsEndsWith() bool
	IsAnywhere() bool
	IsContains() bool
	IsAnyChars() bool
	IsRegex() bool
	IsNegativeCondition() bool
	IsNotEqual() bool
	IsNotStartsWith() bool
	IsNotEndsWith() bool
	IsNotContains() bool
	IsNotMatchRegex() bool

	IsCompareSuccess(
		isIgnoreCase bool,
		content,
		search string,
	) bool
	VerifyMessage(
		isIgnoreCase bool,
		content,
		search string,
	) string
	VerifyError(
		isIgnoreCase bool,
		content,
		search string,
	) error

	VerifyMessageCaseSensitive(
		content,
		search string,
	) string
	IsCompareSuccessCaseSensitive(content, search string) bool
	IsCompareSuccessNonCaseSensitive(content, search string) bool

	IsAnyMethod(methodNames ...string) bool
}
