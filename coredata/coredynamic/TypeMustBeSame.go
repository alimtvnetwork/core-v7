package coredynamic

func TypeMustBeSame(
	left, right interface{},
) {
	err := TypeNotEqualErr(left, right)

	if err != nil {
		panic(err)
	}
}
