package stringcomparator

type (
	IsLineCompareFunc func(
		content,
		searchComparingLine string,
		isCaseSensitive bool,
	) bool

	IsLineContainsFunc func(
		index int,
		contentLine string,
	) bool
)
