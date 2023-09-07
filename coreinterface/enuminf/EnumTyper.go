package enuminf

type EnumTyper interface {
	EnumTypeChecker
	enumNameStinger
	NameValuer
	IsNameEqualer
	IsAnyNameOfChecker
	ValueByte() byte
	Value() byte
	ToNumberStringer
}
