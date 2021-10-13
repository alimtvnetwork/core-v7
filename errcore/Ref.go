package errcore

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

func Ref(reference interface{}) string {
	if reference == nil {
		return ""
	}

	return fmt.Sprintf(
		constants.ReferenceWrapFormat,
		reference)

}
