package strtype

import (
	"gitlab.com/evatix-go/core/corecomparator"
	"gitlab.com/evatix-go/core/msgtype"
)

type Variant string

func (v Variant) Value() string {
	return string(v)
}

func (v Variant) StringValue() string {
	return string(v)
}

// v + n
func (v Variant) Add(n string) Variant {
	return Variant(v.Value() + n)
}

func (v Variant) Is(n Variant) bool {
	return v.Value() == n.Value()
}

func (v Variant) IsEqual(n string) bool {
	return v.Value() == n
}

// v.Value() > n
func (v Variant) IsGreater(n string) bool {
	return v.Value() > n
}

// v.Value() >= n
func (v Variant) IsGreaterEqual(n string) bool {
	return v.Value() >= n
}

// v.Value() < n
func (v Variant) IsLess(n string) bool {
	return v.Value() < n
}

// v.Value() <= n
func (v Variant) IsLessEqual(n string) bool {
	return v.Value() <= n
}

// Here left is v, and right is `n`
func (v Variant) IsTrue(n string, compare corecomparator.Compare) bool {
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
