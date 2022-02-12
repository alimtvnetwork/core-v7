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
	IsInteger() bool
	IsString() bool
	IsValidInvalidChecker
}
