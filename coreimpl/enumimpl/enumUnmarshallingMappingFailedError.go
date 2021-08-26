package enumimpl

import (
	"fmt"
)

func enumUnmarshallingMappingFailedError(
	typeName string,
	valueGiven string,
	rangesAvailable string,
) error {
	return fmt.Errorf(
		errUnmappedMessage,
		typeName,
		valueGiven,
		rangesAvailable)
}
