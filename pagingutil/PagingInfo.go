package pagingutil

type PagingInfo struct {
	PageIndex, SkipItems, EndingLength int
	IsPagingPossible                   bool
}
