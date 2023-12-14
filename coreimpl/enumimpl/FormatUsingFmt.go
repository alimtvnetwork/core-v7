package enumimpl

import "gitlab.com/auk-go/core/internal/strutilinternal"

// FormatUsingFmt
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
func FormatUsingFmt(
	fmt formatter,
	format string,
) string {
	replacerMap := map[string]string{
		typeNameTemplateKey: fmt.TypeName(),
		nameKey:             fmt.Name(),
		valueKey:            fmt.ValueString(),
	}

	return strutilinternal.ReplaceTemplateMap(
		true,
		format,
		replacerMap)
}
