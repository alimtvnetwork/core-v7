package conditional

import "gitlab.com/auk-go/core/constants"

func NilDefByte(
	valuePointer *byte,
) byte {
	if valuePointer == nil {
		return constants.Zero
	}

	return *valuePointer
}

func NilDefBytePtr(
	valuePointer *byte,
) *byte {
	if valuePointer == nil {
		return constants.ZeroBytePtr
	}

	return valuePointer
}

func NilByteVal(
	valuePointer *byte,
	defVal byte,
) byte {
	if valuePointer == nil {
		return defVal
	}

	return *valuePointer
}

func NilByteValPtr(
	valuePointer *byte,
	defVal byte,
) *byte {
	if valuePointer == nil {
		return &defVal
	}

	return valuePointer
}
