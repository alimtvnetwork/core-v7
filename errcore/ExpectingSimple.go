package errcore

import "fmt"

// ExpectingSimple
//
// returns
//
//	"%s - Expect (type:\"%T\")[\"%v\"] != [\"%v\"](type:\"%T\") Actual"
func ExpectingSimple(title, wasExpecting, actual interface{}) string {
	return fmt.Sprintf(
		expectingSimpleMessageFormat,
		title,
		wasExpecting, wasExpecting,
		actual, actual)
}
