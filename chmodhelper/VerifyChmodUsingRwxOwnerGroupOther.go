package chmodhelper

import (
	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
	"gitlab.com/evatix-go/core/errcore"
)

func VerifyChmodUsingRwxOwnerGroupOther(
	location string,
	rwx *chmodins.RwxOwnerGroupOther,
) error {
	if rwx == nil {
		return errcore.
			CannotBeNilOrEmptyMessage.
			Error("rwx is nil", location)
	}

	return VerifyChmod(
		location,
		rwx.String())
}
