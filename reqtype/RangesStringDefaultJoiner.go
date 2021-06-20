package reqtype

func RangesStringDefaultJoiner(requests ...Request) string {
	return RangesString(defaultRangesJoiner, requests...)
}
