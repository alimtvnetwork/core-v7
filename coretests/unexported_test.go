package coretests

import (
	"testing"
)

// These tests use unexported symbols and must remain in source package.

func Test_MessagePrinter_FailedExpected(t *testing.T) {
	p := printMessage{}
	// Arrange + Act: not failed
	p.FailedExpected(false, "when", "actual", "expected", 0)
	// Arrange + Act: failed
	p.FailedExpected(true, "when", "actual", "expected", 1)
}

func Test_MessagePrinter_NameValueAndValue(t *testing.T) {
	p := printMessage{}
	// Act
	p.NameValue("header", map[string]int{"a": 1})
	p.Value("header2", []int{1, 2, 3})
}

func Test_SkipOnUnix(t *testing.T) {
	t.Run("sub", func(st *testing.T) {
		SkipOnUnix(st)
	})
}
