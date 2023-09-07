package errcore

import "gitlab.com/auk-go/core/internal/strutilinternal"

func ErrorToSplitNonEmptyLines(err error) []string {
	lines := ErrorToSplitLines(err)

	return strutilinternal.NonEmptySlice(lines)
}
