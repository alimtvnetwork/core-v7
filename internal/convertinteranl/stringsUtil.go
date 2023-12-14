package convertinteranl

type stringsUtil struct{}

func (it stringsUtil) PrependWithSpaces(
	spaceCountLines int,
	existingLines []string,
	prependingLinesSpaceCount int,
	prependingLines ...string,
) []string {
	var newSlice []string

	if prependingLinesSpaceCount > 0 {
		prependingLines = StringsTo.
			WithSpaces(
				prependingLinesSpaceCount,
				prependingLines...,
			)
	}

	newSlice = append(
		newSlice,
		prependingLines...,
	)

	if spaceCountLines > 0 {
		existingLines = StringsTo.
			WithSpaces(
				spaceCountLines,
				existingLines...,
			)
	}

	newSlice = append(newSlice, existingLines...)

	return newSlice
}
