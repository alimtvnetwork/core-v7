package errcore

import (
	"strings"
)

func StackTracesCompiled(traces []string) string {
	tracesCompiled := CodeStacksHeaderNewLine +
		PrefixStackTrace +
		strings.Join(traces, PrefixStackTraceNewLine)

	return tracesCompiled
}
