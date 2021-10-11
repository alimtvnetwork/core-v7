package errcore

import "fmt"

func ExpectingNotEqualSimpleNoType(
	title,
	wasExpecting,
	actual interface{},
) string {
	return fmt.Sprintf(
		expectingNotMatchingSimpleNoTypeMessageFormat,
		title,
		wasExpecting,
		actual)
}
