package issetter

import (
	"errors"
)

var values = []string{"Uninitialized", "True", "False", "Unset", "Set", "Wildcard"}

type Value byte

const (
	Uninitialized Value = 0
	True          Value = 1
	False         Value = 2
	Unset         Value = 3
	Set           Value = 4
	Wildcard      Value = 5
)

func (v Value) Value() byte {
	return byte(v)
}

func (v Value) StringValue() string {
	return string(v)
}

func (v Value) String() string {
	return values[v]
}

// IsTrue v == True
func (v Value) IsTrue() bool {
	return v == True
}

// IsFalse v == False
func (v Value) IsFalse() bool {
	return v == False
}

// IsSet v == Set
func (v Value) IsSet() bool {
	return v == Set
}

// IsUnset v == Unset
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

func (v Value) IsWildcard() bool {
	return v == Wildcard
}

// WildcardApply
//
// if IsWildcard() || IsUnSetOrUninitialized() then
//
//      return inputVal
// else
//
//      return v. IsTrue()
func (v Value) WildcardApply(inputBool bool) bool {
	if v.IsWildcard() || v.IsUnSetOrUninitialized() {
		return inputBool
	}

	return v.IsTrue()
}

// WildcardValueApply
//
// if IsWildcard() || IsUnSetOrUninitialized() then
//
//      return inputVal
// else
//
//      return v. IsTrue()
func (v Value) WildcardValueApply(inputVal Value) bool {
	if v.IsWildcard() || v.IsUnSetOrUninitialized() {
		return inputVal.IsTrue()
	}

	return v.IsTrue()
}

// OrBool
//
// if IsWildcard() || IsUnSetOrUninitialized() then
//
//      return inputBool
// else
//
//      return v. IsTrue() || inputBool
func (v Value) OrBool(inputBool bool) bool {
	if v.IsWildcard() || v.IsUnSetOrUninitialized() {
		return inputBool
	}

	return v.IsTrue() || inputBool
}

// OrValue
//
// if IsWildcard() || IsUnSetOrUninitialized() then
//
//      return inputVal
// else
//
//      return v. IsTrue() || inputVal. IsTrue()
func (v Value) OrValue(inputVal Value) bool {
	if v.IsWildcard() || v.IsUnSetOrUninitialized() {
		return inputVal.IsTrue()
	}

	return v.IsTrue() || inputVal.IsTrue()
}

// AndBool
//
// if IsWildcard() || IsUnSetOrUninitialized() then
//
//      return inputVal
// else
//
//      return v. IsTrue() && inputBool
func (v Value) AndBool(inputBool bool) bool {
	if v.IsWildcard() || v.IsUnSetOrUninitialized() {
		return inputBool
	}

	return v.IsTrue() && inputBool
}

// And
//
// if IsWildcard() || IsUnSetOrUninitialized() then
//
//      return inputVal
// else
//
//      return GetBool(v. IsTrue() && inputVal. IsTrue())
func (v Value) And(inputVal Value) Value {
	if v.IsWildcard() || v.IsUnSetOrUninitialized() {
		return inputVal
	}

	return GetBool(v.IsTrue() && inputVal.IsTrue())
}

// IsUninitialized v == Uninitialized
func (v Value) IsUninitialized() bool {
	return v == Uninitialized
}

// IsUnSetOrUninitialized v == Uninitialized || v == Unset
func (v Value) IsUnSetOrUninitialized() bool {
	return v == Uninitialized || v == Unset
}

// IsNegative v == Uninitialized || v == Unset || v == False
func (v Value) IsNegative() bool {
	return v == Uninitialized || v == Unset || v == False
}

// IsPositive v == True || v == Set
func (v Value) IsPositive() bool {
	return v == True || v == Set
}

// IsBetween val >= start &&  val <= end
func (v Value) IsBetween(start, end byte) bool {
	val := v.Value()

	return val >= start && val <= end
}

// IsBetweenInt val >= start &&  val <= end
func (v Value) IsBetweenInt(start, end int) bool {
	val := v.Value()

	return val >= byte(start) && val <= byte(end)
}

// Add v + n
func (v Value) Add(n byte) Value {
	return Value(v.Value() + n)
}

func (v Value) Is(n Value) bool {
	return v.Value() == n.Value()
}

func (v Value) IsEqual(n byte) bool {
	return v.Value() == n
}

// IsGreater v.Value() > n
func (v Value) IsGreater(n byte) bool {
	return v.Value() > n
}

// IsGreaterEqual v.Value() >= n
func (v Value) IsGreaterEqual(n byte) bool {
	return v.Value() >= n
}

// IsLess v.Value() < n
func (v Value) IsLess(n byte) bool {
	return v.Value() < n
}

// IsLessEqual v.Value() <= n
func (v Value) IsLessEqual(n byte) bool {
	return v.Value() <= n
}

func (v Value) IsEqualInt(n int) bool {
	return v.Value() == byte(n)
}

// IsGreaterInt v.Value() > n
func (v Value) IsGreaterInt(n int) bool {
	return v.Value() > byte(n)
}

// IsGreaterEqualInt v.Value() >= n
func (v Value) IsGreaterEqualInt(n int) bool {
	return v.Value() >= byte(n)
}

// IsLessInt v.Value() < n
func (v Value) IsLessInt(n int) bool {
	return v.Value() < byte(n)
}

// IsLessEqualInt v.Value() <= n
func (v Value) IsLessEqualInt(n int) bool {
	return v.Value() <= byte(n)
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
