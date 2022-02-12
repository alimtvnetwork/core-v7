package internalinterface

type BaseRawErrorTyper interface {
	String() string
	NameWithNameEqualer
	Combine(
		otherMsg string, reference interface{},
	) string
	TypesAttach(
		otherMsg string,
		reflectionTypes ...interface{},
	) string
	TypesAttachErr(
		otherMsg string,
		reflectionTypes ...interface{},
	) error
	SrcDestination(
		otherMsg string,
		srcName string, srcValue interface{},
		destinationName string, destinationValue interface{},
	) string
	SrcDestinationErr(
		otherMsg string,
		srcName string, srcValue interface{},
		destinationName string, destinationValue interface{},
	) error
	Error(otherMsg string, reference interface{}) error
	MsgCsvRef(
		otherMsg string,
		csvReferenceItems ...interface{},
	) string
	MsgCsvRefError(
		otherMsg string,
		csvReferenceItems ...interface{},
	) error
	ErrorRefOnly(reference interface{}) error
	Expecting(expecting, actual interface{}) error
	NoRef(otherMsg string) string
	ErrorNoRefs(otherMsg string) error
	HandleUsingPanic(otherMsg string, reference interface{})
}
