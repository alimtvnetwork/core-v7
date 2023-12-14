package enuminf

type IsValidChecker interface {
	// IsValid similar or alias for IsSuccessChecker
	IsValid() bool
}

type IsInvalidChecker interface {
	IsInvalid() bool
}

type IsValidInvalidChecker interface {
	IsValidChecker
	IsInvalidChecker
}

type IsNameEqualer interface {
	IsNameEqual(name string) bool
}

// IsAnyNameOfChecker
//
//	Returns true if any of the name matches.
type IsAnyNameOfChecker interface {
	// IsAnyNamesOf
	//
	//  Returns true if any of the name matches.
	IsAnyNamesOf(names ...string) bool
}

type RangeValidateChecker interface {
	// RangesInvalidMessage get invalid message
	RangesInvalidMessage() string
	// RangesInvalidErr get invalid message error
	RangesInvalidErr() error
	// IsValidRange Is with in the range as expected.
	IsValidRange() bool
	// IsInvalidRange Is out of the ranges expected.
	IsInvalidRange() bool
}

type IsStartChecker interface {
	IsStart() bool
}

type IsEndChecker interface {
	IsEnd() bool
}

type IsStartEndChecker interface {
	IsStartChecker
	IsEndChecker
}
