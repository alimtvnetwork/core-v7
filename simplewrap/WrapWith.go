package simplewrap

// With
//
// Combines start + source + end .
func With(start, source, end string) string {
	return start + source + end
}

// WithPtr
//
// Combines start + source + end with nil safety check.
func WithPtr(start, source, end *string) *string {
	var nStart, nEnd, nSource string
	
	if start != nil {
		nStart = *start
	}
	if end != nil {
		nEnd = *end
	}
	
	if source != nil {
		nSource = *source
	}
	
	final := nStart + nSource + nEnd
	
	return &final
}
