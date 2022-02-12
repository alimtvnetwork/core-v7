package enuminf

type EnumTyper interface {
	EnumTypeChecker
	enumNameStinger
	nameValuer
	IsNameEqualer
	IsAnyNameOfChecker
	ValueByte() byte
	Value() byte
	ToNumberStringer
}
