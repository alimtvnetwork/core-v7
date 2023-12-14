package chmodins

// ExpandRwxFullStringToOwnerGroupOtherByFixingFirst can be
//
//   - "-rwx" will be "-rwx******"
//   - "-rwxr-x" will be "-rwxr-x***"
//   - "-rwxr-x" will be "-rwxr-x***"
func ExpandRwxFullStringToOwnerGroupOtherByFixingFirst(
	rwxPartialRestWildcard string,
) (*RwxOwnerGroupOther, error) {
	fullRwxWithWildcard := FixRwxFullStringWithWildcards(rwxPartialRestWildcard)

	return ExpandRwxFullStringToOwnerGroupOther(
		fullRwxWithWildcard)
}
