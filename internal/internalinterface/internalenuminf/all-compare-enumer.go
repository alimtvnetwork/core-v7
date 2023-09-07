package internalenuminf

type CompareByteEnumer interface {
	IsEqual(comparingValue byte) bool
	IsLess(comparingValue byte) bool
	IsLessEqual(comparingValue byte) bool
	IsGreater(comparingValue byte) bool
	IsGreaterEqual(comparingValue byte) bool
	IsNotEqual(comparingValue byte) bool
}

type CompareInt8Enumer interface {
	IsEqual(comparingValue int8) bool
	IsLess(comparingValue int8) bool
	IsLessEqual(comparingValue int8) bool
	IsGreater(comparingValue int8) bool
	IsGreaterEqual(comparingValue int8) bool
	IsNotEqual(comparingValue int8) bool
}

type CompareInt16Enumer interface {
	IsEqual(comparingValue int16) bool
	IsLess(comparingValue int16) bool
	IsLessEqual(comparingValue int16) bool
	IsGreater(comparingValue int16) bool
	IsGreaterEqual(comparingValue int16) bool
	IsNotEqual(comparingValue int16) bool
}

type CompareInt32Enumer interface {
	IsEqual(comparingValue int32) bool
	IsLess(comparingValue int32) bool
	IsLessEqual(comparingValue int32) bool
	IsGreater(comparingValue int32) bool
	IsGreaterEqual(comparingValue int32) bool
	IsNotEqual(comparingValue int32) bool
}

type CompareIntegerEnumer interface {
	IsEqual(comparingValue int) bool
	IsLess(comparingValue int) bool
	IsLessEqual(comparingValue int) bool
	IsGreater(comparingValue int) bool
	IsGreaterEqual(comparingValue int) bool
	IsNotEqual(comparingValue int) bool
}

type CompareBasicEnumer interface {
	IsEqual(enum BasicEnumer) bool
	IsLess(enum BasicEnumer) bool
	IsLessEqual(enum BasicEnumer) bool
	IsGreater(enum BasicEnumer) bool
	IsGreaterEqual(enum BasicEnumer) bool
	IsNotEqual(enum BasicEnumer) bool
}
