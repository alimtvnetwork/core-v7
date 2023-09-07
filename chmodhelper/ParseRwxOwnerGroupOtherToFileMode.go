package chmodhelper

import (
	"os"

	"gitlab.com/auk-go/core/chmodhelper/chmodins"
	"gitlab.com/auk-go/core/errcore"
)

func ParseRwxOwnerGroupOtherToFileMode(
	rwxOwnerGroupOther *chmodins.RwxOwnerGroupOther,
) (os.FileMode, error) {
	varWrapper, err := ParseRwxOwnerGroupOtherToRwxVariableWrapper(
		rwxOwnerGroupOther)

	if err != nil {
		return 0, errcore.MeaningfulErrorWithData(
			errcore.FailedToConvertType,
			"ParseRwxOwnerGroupOtherToFileMode",
			err,
			rwxOwnerGroupOther)
	}

	return varWrapper.ToCompileFixedPtr().ToFileMode(), nil
}
