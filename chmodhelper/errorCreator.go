package chmodhelper

import (
	"errors"
	"os"
)

type errorCreator struct{}

func (it errorCreator) dirError(dirPath string, err error) error {
	err2 := it.notDirError(dirPath)
	if err2 != nil {
		return err2
	}

	// has err
	return errors.New(
		"dir : " + dirPath +
			", applyChmod :" + dirDefaultChmod.String() +
			", " + err.Error(),
	)
}

func (it errorCreator) notDirError(dirPath string) error {
	if IsPathInvalid(dirPath) {
		return nil
	}

	// exist

	if !IsDirectory(dirPath) {
		return errors.New(
			"dir : " + dirPath +
				", applyChmod :" + dirDefaultChmod.String() +
				", path exist but it is not a dir.",
		)
	}

	return nil
}

func (it errorCreator) pathError(
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
		err,
	)

	return errors.New(compiledMessage)
}

func (it errorCreator) pathErrorWithDirValidate(
	message string,
	applyChmod os.FileMode,
	location string,
	err error,
) error {
	notDirErr := it.notDirError(location)

	if notDirErr != nil {
		return notDirErr
	}

	if err == nil {
		return nil
	}

	compiledMessage := pathErrorMessage(
		message,
		applyChmod,
		location,
		err,
	)

	return errors.New(compiledMessage)
}

func (it errorCreator) chmodApplyFailed(
	applyChmod os.FileMode,
	location string,
	err error,
) error {
	if err == nil {
		return nil
	}

	compiledMessage := pathErrorMessage(
		"chmod apply failed",
		applyChmod,
		location,
		err,
	)

	return errors.New(compiledMessage)
}
