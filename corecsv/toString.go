package corecsv

import "gitlab.com/auk-go/core/internal/convertinteranl"

func toString(
	source interface{},
) string {
	return convertinteranl.AnyTo.SmartString(source)
}
