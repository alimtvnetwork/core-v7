package errcore

import "gitlab.com/evatix-go/core/internal/utilstringinternal"

func ErrorToSplitNonEmptyLines(err error) []string {
	lines := ErrorToSplitLines(err)

	return utilstringinternal.NonEmptySlice(lines)
}
