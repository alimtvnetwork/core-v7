package issetter

import (
	"errors"

	"gitlab.com/evatix-go/core/corecomparator"
	"gitlab.com/evatix-go/core/msgtype"
)

var values = []string{"Uninitialized", "True", "False", "Unset", "Set"}

type Value byte

const (
	Uninitialized Value = 0
	True          Value = 1
	False         Value = 2
	Unset         Value = 3
	Set           Value = 4
)

func GetSet(
	isCondition bool,
	trueValue Value,
	falseValue Value,
) Value {
	if isCondition {
		return trueValue
	}

	return falseValue
}

func GetSetByte(
	isCondition bool,
	trueValue byte,
	falseValue byte,
) Value {
	if isCondition {
		return Value(trueValue)
	}

	return Value(falseValue)
}

func (v Value) Value() byte {
	return byte(v)
}

func (v Value) StringValue() string {
	return string(v)
}

func (v Value) String() string {
	return values[v]
}

// v == True
func (v Value) IsTrue() bool {
	return v == True
}

// v == False
func (v Value) IsFalse() bool {
	return v == False
}

// v == Set
func (v Value) IsSet() bool {
	return v == Set
}

// v == Unset
func (v Value) IsUnset() bool {
	return v == Unset
}

func (v Value) HasInitialized() bool {
	return v != Uninitialized
}

func (v Value) HasInitializedAndSet() bool {
	return v == Set
}

func (v Value) HasInitializedAndTrue() bool {
	return v == True
}

// v == Uninitialized
func (v Value) IsUninitialized() bool {
	return v == Uninitialized
}

// v == Uninitialized || v == Unset
func (v Value) IsUnSetOrUninitialized() bool {
	return v == Uninitialized || v == Unset
}

// v == Uninitialized || v == Unset || v == False
func (v Value) IsNegative() bool {
	return v == Uninitialized || v == Unset || v == False
}

// v == True || v == Set
func (v Value) IsPositive() bool {
	return v == True || v == Set
}

// val >= start &&  val <= end
func (v Value) IsBetween(start, end byte) bool {
	val := v.Value()

	return val >= start && val <= end
}

// val >= start &&  val <= end
func (v Value) IsBetweenInt(start, end int) bool {
	val := v.Value()

	return val >= byte(start) && val <= byte(end)
}

// v + n
func (v Value) Add(n byte) Value {
	return Value(v.Value() + n)
}

func (v Value) Is(n Value) bool {
	return v.Value() == n.Value()
}

func (v Value) IsEqual(n byte) bool {
	return v.Value() == n
}

// v.Value() > n
func (v Value) IsGreater(n byte) bool {
	return v.Value() > n
}

// v.Value() >= n
func (v Value) IsGreaterEqual(n byte) bool {
	return v.Value() >= n
}

// v.Value() < n
func (v Value) IsLess(n byte) bool {
	return v.Value() < n
}

// v.Value() <= n
func (v Value) IsLessEqual(n byte) bool {
	return v.Value() <= n
}

func (v Value) IsEqualInt(n int) bool {
	return v.Value() == byte(n)
}

// v.Value() > n
func (v Value) IsGreaterInt(n int) bool {
	return v.Value() > byte(n)
}

// v.Value() >= n
func (v Value) IsGreaterEqualInt(n int) bool {
	return v.Value() >= byte(n)
}

// v.Value() < n
func (v Value) IsLessInt(n int) bool {
	return v.Value() < byte(n)
}

// v.Value() <= n
func (v Value) IsLessEqualInt(n int) bool {
	return v.Value() <= byte(n)
}

// n < Uninitialized.Value() || n > Set.Value()
func IsOutOfRange(n byte) bool {
	return n < Uninitialized.Value() || n > Set.Value()
}

func (v Value) PanicOnOutOfRange(n byte, msg string) {
	if IsOutOfRange(n) {
		panic(msg)
	}
}

func (v Value) GetErrorOnOutOfRange(n byte, msg string) error {
	if IsOutOfRange(n) {
		return errors.New(msg)
	}

	return nil
}

// Here left is v, and right is `n`
func (v Value) IsCompareResult(n byte, compare corecomparator.Compare) bool {
	switch compare {
	case corecomparator.Equal:
		return v.IsEqual(n)
	case corecomparator.LeftGreater:
		return v.IsGreater(n)
	case corecomparator.LeftGreaterEqual:
		return v.IsGreaterEqual(n)
	case corecomparator.LeftLess:
		return v.IsLess(n)
	case corecomparator.LeftLessEqual:
		return v.IsLessEqual(n)
	default:
		msg := msgtype.RangeNotMeet(
			msgtype.ComparatorShouldBeWithinRanghe.String(),
			corecomparator.Min(),
			corecomparator.Max(),
			compare.Ranges())

		panic(msg)
	}
}

func Min() Value {
	return Uninitialized
}

func Max() Value {
	return Set
}

func MinByte() byte {
	return Uninitialized.Value()
}

func MaxByte() byte {
	return Set.Value()
}
