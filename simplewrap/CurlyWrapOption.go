package simplewrap

import (
	"gitlab.com/auk-go/core/internal/convertinteranl"
)

func CurlyWrapOption(
	isSkipIfExists bool,
	source interface{},
) string {
	toStr := convertinteranl.
		AnyTo.
		SmartString(source)
	
	if isSkipIfExists {
		return ConditionalWrapWith(
			'{',
			toStr,
			'}')
	}
	
	return CurlyWrap(source)
}
