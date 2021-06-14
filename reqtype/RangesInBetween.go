package reqtype

func RangesInBetween(
	start, end Request,
) []Request {
	startVal := start.Value()
	endVal := end.Value()
	length := endVal - startVal + 1

	slice := make(
		[]Request,
		length)

	index := 0
	for i := startVal; i < endVal; i++ {
		slice[index] = Request(i)
		index++
	}

	return slice
}
