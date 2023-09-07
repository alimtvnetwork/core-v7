package chmodhelper

import (
	"errors"
	"os"
)

func pathError(
	message string,
	applyChmod os.FileMode,
	location string,
	err error,
) error {
	if err == nil {
		return nil
	}

	compiledMessage := pathErrorMessage(
		message,
		applyChmod,
		location,
		err)

	return errors.New(compiledMessage)
}
