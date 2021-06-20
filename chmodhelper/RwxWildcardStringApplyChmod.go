package chmodhelper

import "gitlab.com/evatix-go/core/chmodhelper/chmodins"

// RwxWildcardStringApplyChmod
// rwxPartial can be any length in
// between 0-10 (rest will be fixed by wildcard)
//
// rwxPartial:
//  - "rwx" will be "-rwx******"
//  - "rwxr-x" will be "-rwxr-x***"
//  - "-rwxr-x" will be "-rwxr-x***"
func RwxWildcardStringApplyChmod(
	rwxPartial string,
	condition *chmodins.Condition,
	locations ...string,
) error {
	if len(locations) == 0 {
		return nil
	}

	rwxInstructionExecutor, err := RwxWildcardStringToInstructionExecutor(
		rwxPartial,
		condition)

	if err != nil {
		return err
	}

	return rwxInstructionExecutor.
		ApplyOnPathsPtr(&locations)
}
