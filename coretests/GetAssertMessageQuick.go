package coretests

import (
	"fmt"

	"gitlab.com/auk-go/core/internal/msgformats"
)

// GetAssertMessageQuick
//
//  Gives generic and consistent test message using msgformats.QuickIndexInputActualExpectedMessageFormat
func GetAssertMessageQuick(
	when,
	actual,
	expected interface{},
	counter int,
) string {
	return fmt.Sprintf(
		msgformats.QuickIndexInputActualExpectedMessageFormat,
		counter,
		when,
		actual,
		expected,
	)
}
