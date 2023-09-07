package enumimpl

import (
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/internal/csvinternal"
)

func OnlySupportedErr(
	allNames []string,
	supportedNames ...string,
) error {
	if len(allNames) == 0 {
		return nil
	}

	unsupportedNames := UnsupportedNames(
		allNames,
		supportedNames...)

	if len(unsupportedNames) == 0 {
		return nil
	}

	supportedCsv := csvinternal.StringsToStringDefault(
		supportedNames...)

	unsupportedCsv := csvinternal.StringsToStringDefault(
		unsupportedNames...)

	supportedMsg := "Only supported (" + supportedCsv + ")"
	unsupportedMsg := "Unsupported (" + unsupportedCsv + ")"

	return errcore.
		RangesOnlySupportedType.
		ErrorNoRefs(
			supportedMsg + unsupportedMsg)
}
