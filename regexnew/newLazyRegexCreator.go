package regexnew

type newLazyRegexCreator struct{}

// New
//
//	used to create as vars
func (it newLazyRegexCreator) New(
	pattern string,
) *LazyRegex {
	lazyRegex, _ := lazyRegexOnceMap.CreateOrExisting(
		pattern)

	return lazyRegex
}

// NewLock
//
//	used to generate inside method
func (it newLazyRegexCreator) NewLock(
	pattern string,
) *LazyRegex {
	lazyRegex, _ := lazyRegexOnceMap.CreateOrExistingLock(
		pattern)

	return lazyRegex
}

func (it newLazyRegexCreator) TwoLock(
	pattern, secondPattern string,
) (first, second *LazyRegex) {
	lazyRegexLock.Lock()
	defer lazyRegexLock.Unlock()

	first = it.New(pattern)
	second = it.New(secondPattern)

	return first, second
}

func (it newLazyRegexCreator) ManyUsingLock(
	patterns ...string,
) (patternsKeyAsMap map[string]*LazyRegex) {
	if len(patterns) == 0 {
		return map[string]*LazyRegex{}
	}

	lazyRegexLock.Lock()
	defer lazyRegexLock.Unlock()

	patternsKeyAsMap = make(
		map[string]*LazyRegex,
		len(patterns))

	for _, pattern := range patterns {
		patternsKeyAsMap[pattern] = it.New(pattern)
	}

	return patternsKeyAsMap
}

func (it newLazyRegexCreator) AllPatternsMap() map[string]*LazyRegex {
	lazyRegexLock.Lock()
	defer lazyRegexLock.Unlock()

	return lazyRegexOnceMap.items
}

// NewLockIf
//
//	used to generate inside method
//	lock must be performed when called from method.
func (it newLazyRegexCreator) NewLockIf(
	isLock bool,
	pattern string,
) *LazyRegex {
	if isLock {
		return it.NewLock(pattern)
	}

	return it.New(pattern)
}
