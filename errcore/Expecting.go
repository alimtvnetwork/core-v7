package errcore

import (
	"fmt"
)

// Expecting
//
// returns
//      "%s - expecting (type:[%T]) : [\"%v\"], but received or actual (type:[%T]) : [\"%v\"]"
func Expecting(title, wasExpecting, actual interface{}) string {
	return fmt.Sprintf(
		expectingMessageFormat,
		title,
		wasExpecting, wasExpecting,
		actual, actual)
}
