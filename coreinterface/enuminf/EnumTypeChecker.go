package enuminf

type EnumTypeChecker interface {
	IsBoolean() bool
	IsByte() bool
	IsUnsignedInteger16() bool
	IsUnsignedInteger32() bool
	IsUnsignedInteger64() bool
	IsInteger8() bool
	IsInteger16() bool
	IsInteger32() bool
	IsInteger64() bool
	IsInteger() bool
	IsString() bool
	// IsNumber
	//
	// Any number returns true
	IsNumber() bool
	IsAnyInteger() bool
	IsAnyUnsignedNumber() bool
	IsValidInvalidChecker
}
