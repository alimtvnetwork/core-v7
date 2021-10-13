package errcore

import (
	"errors"
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

func RefToError(reference interface{}) error {
	if reference == nil {
		return nil
	}

	return errors.New(fmt.Sprintf(
		constants.ReferenceWrapFormat,
		reference),
	)
}
