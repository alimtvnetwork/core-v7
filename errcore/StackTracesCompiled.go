package errcore

import (
	"strings"
)

// StackTracesCompiled
//
// CodeStacksHeaderNewLine + PrefixStackTrace + JoinedWith(PrefixStackTraceNewLine)
//
// Example:
//   - "Code Stack :\n- JoinLinesWith(\n- )"
func StackTracesCompiled(traces []string) string {
	tracesCompiled := CodeStacksHeaderNewLine +
		PrefixStackTrace +
		strings.Join(traces, PrefixStackTraceNewLine)

	return tracesCompiled
}
