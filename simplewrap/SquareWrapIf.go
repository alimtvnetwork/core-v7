package simplewrap

func SquareWrapIf(
	isSquareWrap bool,
	source interface{},
) string {
	if !isSquareWrap {
		return toString(source)
	}
	
	return SquareWrap(source)
}
