package chmodhelper

func MergeRwxWildcardWithFixedRwx(rwxWildcard, rwxFixed string) (fixedAttribute *Attribute, err error) {
	length := len(rwxWildcard)

	if length != SupportedLength {
		return nil, GetRwxLengthError(rwxWildcard)
	}

	length2 := len(rwxFixed)

	if length2 != SupportedLength {
		return nil, GetRwxLengthError(rwxFixed)
	}

	varAttr, err := ParseRwxToVarAttribute(rwxWildcard)

	if err != nil {
		return nil, err
	}

	attr := NewAttributeUsingRwx(rwxFixed)
	fixedAttr := varAttr.ToCompileAttr(&attr)

	return &fixedAttr, nil
}
