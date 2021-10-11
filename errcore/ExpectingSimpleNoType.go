package errcore

import "fmt"

// ExpectingSimpleNoType
//
// returns
//      "%s - Expect [\"%v\"] != [\"%v\"] Actual"
func ExpectingSimpleNoType(title, wasExpecting, actual interface{}) string {
	return fmt.Sprintf(
		expectingSimpleNoTypeMessageFormat,
		title,
		wasExpecting,
		actual)
}
