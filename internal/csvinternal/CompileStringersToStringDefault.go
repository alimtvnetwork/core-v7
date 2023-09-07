package csvinternal

import "gitlab.com/auk-go/core/constants"

func CompileStringersToStringDefault(
	compileStringerFunctions ...func() string,
) string {
	return CompileStringersToString(
		constants.CommaSpace,
		true,
		false,
		compileStringerFunctions...)
}
