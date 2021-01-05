package simplewrap

// wrapper + source + wrapper
func WithStartEnd(wrapper, source string) string {
	return wrapper + source + wrapper
}

// wrapper + source + wrapper
func WithStartEndPtr(wrapper, source *string) *string {
	final := *wrapper + *source + *wrapper

	return &final
}
