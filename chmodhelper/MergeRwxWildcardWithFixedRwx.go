package chmodhelper

// MergeRwxWildcardWithFixedRwx
//
//  - existingRwx : Usually refers to fixed rwx values like "rwx", "--x", "-w-" etc.
//  - rwxWildcardInput : Usually refers to fixed rwx values like "rw*", "*-x", "-w-" etc.
//      Wildcard means keep the existing value as is.
func MergeRwxWildcardWithFixedRwx(
	existingRwx,
	rwxWildcardInput string,
) (
	fixedAttribute *Attribute,
	err error,
) {
	length := len(rwxWildcardInput)

	if length != SingleRwxLength {
		return nil, GetRwxLengthError(rwxWildcardInput)
	}

	length2 := len(existingRwx)

	if length2 != SingleRwxLength {
		return nil, GetRwxLengthError(existingRwx)
	}

	varAttr, err := ParseRwxToVarAttribute(rwxWildcardInput)

	if err != nil {
		return nil, err
	}

	attr := NewAttributeUsingRwx(existingRwx)
	fixedAttr := varAttr.ToCompileAttr(&attr)

	return &fixedAttr, nil
}
