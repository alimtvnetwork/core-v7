package enumimpl

type valueByter interface {
	Value() byte
}

type exactValueByter interface {
	ValueByte() byte
}

type valueInter interface {
	Value() int
}

type exactValueInter interface {
	ValueInt() int
}

type valueInt8er interface {
	Value() int8
}

type exactValueInt8er interface {
	ValueInt8() int8
}

type valueUInt16er interface {
	Value() uint16
}

type exactValueUInt16er interface {
	ValueUInt16() uint16
}

type formatter interface {
	TypeName() string
	Name() string
	ValueString() string
}

type DifferChecker interface {
	GetSingleDiffResult(isLeft bool, l, r interface{}) interface{}
	GetResultOnKeyMissingInRightExistInLeft(lKey string, lVal interface{}) interface{}
	IsEqual(isRegardless bool, l, r interface{}) bool
}
