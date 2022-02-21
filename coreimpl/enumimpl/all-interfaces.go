package enumimpl

type valueByter interface {
	Value() byte
}

type exactValueByter interface {
	ValueByte() byte
}
