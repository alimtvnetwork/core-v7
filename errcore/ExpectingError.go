package errcore

import (
	"errors"
)

func ExpectingErrorSimpleNoType(
	title,
	wasExpecting,
	actual interface{},
) error {
	msg := ExpectingSimpleNoType(
		title,
		wasExpecting,
		actual)

	return errors.New(msg)
}
