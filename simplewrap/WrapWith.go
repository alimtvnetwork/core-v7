package simplewrap

func With(start, end, source string) string {
	return start + source + end
}

func WithPtr(start, end, source *string) *string {
	final := *start + *source + *end

	return &final
}
