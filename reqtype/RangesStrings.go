package reqtype

func RangesStrings(requests ...Request) []string {
	slice := make([]string, len(requests))

	if len(requests) == 0 {
		return slice
	}

	for i, request := range requests {
		slice[i] = request.Name()
	}

	return slice
}
