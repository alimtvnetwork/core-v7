package inttype

import (
	"strconv"

	"gitlab.com/evatix-go/core/corecomparator"
	"gitlab.com/evatix-go/core/msgtype"
)

type Variant int

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

// v + n
func (v Variant) Add(n int) Variant {
	return Variant(v.Value() + n)
}

// v - n
func (v Variant) Subtract(n int) Variant {
	return Variant(v.Value() - n)
}

func (v Variant) Is(n Variant) bool {
	return v.Value() == n.Value()
}

func (v Variant) IsEqual(n int) bool {
	return v.Value() == n
}

// v.Value() > n
func (v Variant) IsGreater(n int) bool {
	return v.Value() > n
}

// v.Value() >= n
func (v Variant) IsGreaterEqual(n int) bool {
	return v.Value() >= n
}

// v.Value() < n
func (v Variant) IsLess(n int) bool {
	return v.Value() < n
}

// v.Value() <= n
func (v Variant) IsLessEqual(n int) bool {
	return v.Value() <= n
}

// Here left is v, and right is `n`
func (v Variant) IsTrue(n int, compare corecomparator.Compare) bool {
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
