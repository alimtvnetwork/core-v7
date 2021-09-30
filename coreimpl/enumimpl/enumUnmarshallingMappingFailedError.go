package enumimpl

import (
	"fmt"
)

func enumUnmarshallingMappingFailedError(
	typeName string,
	valueGiven string,
	rangesAvailable string,
) error {
	//goland:noinspection GoErrorStringFormat
	return fmt.Errorf(
		errUnmappedMessage,
		typeName,
		valueGiven,
		rangesAvailable)
}
