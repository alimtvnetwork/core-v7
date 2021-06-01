package corecomparator

type Compare byte

var compares = []string{"Equal", "LeftGreater", "LeftGreaterEqual", "LeftLess", "LeftLessEqual", "NotEqual"}

const (
	Equal Compare = iota
	LeftGreater
	LeftGreaterEqual
	LeftLess
	LeftLessEqual
	NotEqual
)

func (compare Compare) Is(other Compare) bool {
	return compare == other
}

func (compare Compare) IsEqual() bool {
	return compare == Equal
}

func (compare Compare) IsLeftGreater() bool {
	return compare == LeftGreater
}

func (compare Compare) IsLeftGreaterEqual() bool {
	return compare == LeftGreaterEqual
}

func (compare Compare) IsLeftLess() bool {
	return compare == LeftLess
}

func (compare Compare) IsLeftLessEqual() bool {
	return compare == LeftLessEqual
}

func (compare Compare) IsNotEqual() bool {
	return compare == NotEqual
}

func (compare Compare) Value() byte {
	return byte(compare)
}

func (compare Compare) StringValue() string {
	return string(compare)
}

func (compare Compare) String() string {
	return compares[compare]
}

func (compare Compare) Ranges() []string {
	return compares
}
