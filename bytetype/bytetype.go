package bytetype

import (
	"gitlab.com/evatix-go/core/corecomparator"
	"gitlab.com/evatix-go/core/msgtype"
)

type Variant byte

func (v Variant) Value() byte {
	return byte(v)
}

func (v Variant) StringValue() string {
	return string(v)
}

// v + n
func (v Variant) Add(n byte) Variant {
	return Variant(v.Value() + n)
}

// v - n
func (v Variant) Subtract(n byte) Variant {
	return Variant(v.Value() - n)
}

func (v Variant) Is(n Variant) bool {
	return v.Value() == n.Value()
}

func (v Variant) IsEqual(n byte) bool {
	return v.Value() == n
}

// v.Value() > n
func (v Variant) IsGreater(n byte) bool {
	return v.Value() > n
}

// v.Value() >= n
func (v Variant) IsGreaterEqual(n byte) bool {
	return v.Value() >= n
}

// v.Value() < n
func (v Variant) IsLess(n byte) bool {
	return v.Value() < n
}

// v.Value() <= n
func (v Variant) IsLessEqual(n byte) bool {
	return v.Value() <= n
}

// Here left is v, and right is `n`
func (v Variant) IsTrue(n byte, compare corecomparator.Compare) bool {
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
