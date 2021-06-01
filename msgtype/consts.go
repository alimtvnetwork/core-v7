package msgtype

const (
	ReferenceStart                           = " Reference ( "
	ReferenceEnd                             = " )"
	rangeWithRangeFormat                     = "Range must be in between %v and %v. Range {%#v}"
	rangeWithoutRangeFormat                  = "Range must be in between %v and %v."
	CannotConvertStringToByteForLessThanZero = "Cannot convert string to byte. String cannot be less than 0 for byte."
	CannotConvertStringToByteForMoreThan255  = "Cannot convert string to byte. String is a number but larger than byte size. At max it could be 255."
	CannotConvertStringToByte                = "Cannot convert string to byte."
)
