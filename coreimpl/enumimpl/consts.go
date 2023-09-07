package enumimpl

const (
	errUnmappedMessage = "typename:[%s], value given : [\"%s\"], cannot find in the enum map. " +
		"reference values can be [%s], " +
		"brackets values can be used to unmarshal as well."
	typeNameTemplateKey                 = "type-name"
	nameKey                             = "name"
	valueKey                            = "value"
	diffBetweenMapShouldBeMessageFormat = "%s\n\nDifference Between Map:\n\n{%s}"                // title, diff string
	actualVsExpectingMessageFormat      = "%s :\n\nActual:\n%s\n\nExpecting:\n%s\n\n"            // title, actual, expecting
	curlyWrapFormat                     = "{\n\n%s\n\n}"                                         // jsonValueString
	currentValueNotFoundInJsonMapFormat = "current given value (%v) is not found in the map. %s" // anyValue, RangesInvalidMessage
)
