package coreinterface

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
