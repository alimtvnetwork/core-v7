package chmodhelper

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/errcore"
)

func VerifyChmodPaths(
	locations *[]string,
	expectedHyphenedRwx string,
	isContinueOnError bool,
) error {
	if locations == nil || len(*locations) == 0 {
		return errcore.CannotBeNilOrEmptyMessage.
			Error(constants.EmptyString, nil)
	}

	if !isContinueOnError {
		for _, location := range *locations {
			err := VerifyChmod(location, expectedHyphenedRwx)

			if err != nil {
				return err
			}
		}
	}

	slice := corestr.NewCollection(constants.Zero)

	for _, location := range *locations {
		err := VerifyChmod(location, expectedHyphenedRwx)

		//goland:noinspection ALL
		slice.AddIf(err != nil, err.Error())
	}

	return errcore.SliceErrorDefault(slice.ListPtr())
}
