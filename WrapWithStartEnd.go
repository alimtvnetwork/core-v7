package constants

// wrapper + source + wrapper
func WrapWithStartEnd(wrapper, source string) string {
	return wrapper + source + wrapper
}

// wrapper + source + wrapper
func WrapWithStartEndPtr(wrapper, source *string) *string {
	final := *wrapper + *source + *wrapper

	return &final
}
