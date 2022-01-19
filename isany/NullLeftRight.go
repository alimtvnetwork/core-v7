package isany

func NullLeftRight(
	leftAnyItem,
	rightAnyItem interface{},
) (
	isLeftNull, isRightNull bool,
) {
	return Null(leftAnyItem),
		Null(rightAnyItem)
}
