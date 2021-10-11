package pagingutil

import "gitlab.com/evatix-go/core/errcore"

func GetPagingInfo(request PagingRequest) PagingInfo {
	length := request.Length

	if length < request.EachPageSize {
		return PagingInfo{
			PageIndex:        request.PageIndex,
			SkipItems:        0,
			EndingLength:     request.EachPageSize,
			IsPagingPossible: false,
		}
	}

	/**
	 * eachPageItems = 10
	 * pageIndex = 4
	 * skipItems = 10 * (4 - 1) = 30
	 */
	skipItems := request.EachPageSize * (request.PageIndex - 1)
	if skipItems < 0 {
		errcore.
			CannotBeNegativeIndex.
			HandleUsingPanic(
				"pageIndex cannot be negative or zero.",
				request.PageIndex)
	}

	endingIndex := skipItems + request.EachPageSize

	if endingIndex > length {
		endingIndex = length
	}

	return PagingInfo{
		PageIndex:        request.PageIndex,
		SkipItems:        skipItems,
		EndingLength:     endingIndex,
		IsPagingPossible: true,
	}
}
