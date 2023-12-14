package regexnew

import "regexp"

type newCreator struct {
	LazyRegex newLazyRegexCreator
}

// Lazy
//
//	used to create from vars
func (it newCreator) Lazy(pattern string) *LazyRegex {
	return it.LazyRegex.New(pattern)
}

// LazyLock
//
//	used from method inside to create
func (it newCreator) LazyLock(pattern string) *LazyRegex {
	return it.LazyRegex.NewLock(pattern)
}

func (it newCreator) Default(pattern string) (*regexp.Regexp, error) {
	return Create(pattern)
}

func (it newCreator) DefaultLock(pattern string) (*regexp.Regexp, error) {
	return CreateLock(pattern)
}

func (it newCreator) DefaultLockIf(
	isLock bool,
	pattern string,
) (*regexp.Regexp, error) {
	return CreateLockIf(
		isLock,
		pattern)
}

func (it newCreator) DefaultApplicableLock(pattern string) (
	regEx *regexp.Regexp,
	err error,
	isApplicable bool,
) {
	return CreateApplicableLock(pattern)
}
