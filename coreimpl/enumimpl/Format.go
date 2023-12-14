package enumimpl

import "gitlab.com/auk-go/core/internal/strutilinternal"

// Format
//
//	Outputs name and
//	value by given format.
//
// sample-format :
//   - "Enum of {type-name} - {name} - {value}"
//
// sample-format-output :
//   - "Enum of EnumFullName - Invalid - 0"
//
// Key-Meaning :
//   - {type-name} : represents type-name string
//   - {name}      : represents name string
//   - {value}     : represents value string
func Format(
	typeName,
	enumName,
	valueString string,
	format string,
) string {
	replacerMap := map[string]string{
		typeNameTemplateKey: typeName,
		nameKey:             enumName,
		valueKey:            valueString,
	}

	return strutilinternal.ReplaceTemplateMap(
		true,
		format,
		replacerMap)
}
