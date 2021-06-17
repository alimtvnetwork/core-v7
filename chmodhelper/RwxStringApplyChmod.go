package chmodhelper

import "gitlab.com/evatix-go/core/chmodhelper/chmodins"

// RwxStringApplyChmod rwxFullString 10 chars "-rwxrwxrwx"
func RwxStringApplyChmod(
	rwxFullString string, // rwxFullString 10 chars "-rwxrwxrwx"
	condition *chmodins.Condition,
	locations ...string,
) error {
	rwxWrapper, err := NewUsingHyphenedRwxFullString(rwxFullString)

	if err != nil {
		return err
	}

	return rwxWrapper.ApplyLinuxChmodOnMany(
		condition,
		locations...)
}
