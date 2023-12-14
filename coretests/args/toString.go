package args

import "gitlab.com/auk-go/core/internal/convertinteranl"

func toString(i interface{}) string {
	return convertinteranl.AnyTo.SmartString(i)
}
