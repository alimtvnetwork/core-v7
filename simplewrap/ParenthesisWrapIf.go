package simplewrap

func ParenthesisWrapIf(
	isSquareWrap bool,
	source interface{},
) string {
	if !isSquareWrap {
		return toString(source)
	}
	
	return ParenthesisWrap(source)
}
