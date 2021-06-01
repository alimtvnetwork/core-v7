package inttype

import (
	"strconv"
)

type Variant int

const (
	Uninitialized Variant = iota
)

func (v Variant) Value() int {
	return int(v)
}

//goland:noinspection ALL
func GetSet(
	isCondition bool,
	trueValue Variant,
	falseValue Variant,
) Variant {
	if isCondition {
		return trueValue
	}

	return falseValue
}

//goland:noinspection ALL
func GetSetVariant(
	isCondition bool,
	trueValue int,
	falseValue int,
) Variant {
	if isCondition {
		return Variant(trueValue)
	}

	return Variant(falseValue)
}

func (v Variant) StringValue() string {
	return strconv.Itoa(v.Value())
}

func (v Variant) IsUninitialized() bool {
	return v == Uninitialized
}

// Add v + n
func (v Variant) Add(n int) Variant {
	return Variant(v.Value() + n)
}

// Subtract v - n
func (v Variant) Subtract(n int) Variant {
	return Variant(v.Value() - n)
}

func (v Variant) Is(n Variant) bool {
	return v.Value() == n.Value()
}

func (v Variant) IsEqual(n int) bool {
	return v.Value() == n
}

// IsGreater v.Value() > n
func (v Variant) IsGreater(n int) bool {
	return v.Value() > n
}

// IsGreaterEqual v.Value() >= n
func (v Variant) IsGreaterEqual(n int) bool {
	return v.Value() >= n
}

// IsLess v.Value() < n
func (v Variant) IsLess(n int) bool {
	return v.Value() < n
}

// IsLessEqual v.Value() <= n
func (v Variant) IsLessEqual(n int) bool {
	return v.Value() <= n
}
