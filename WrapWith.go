package constants

func WrapWith(start, end, source string) string {
	return start + source + end
}

func WrapWithPtr(start, end, source *string) *string {
	final := *start + *source + *end

	return &final
}
