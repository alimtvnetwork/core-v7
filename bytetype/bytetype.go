package bytetype

import (
	"gitlab.com/evatix-go/core/coreinterface"
	"gitlab.com/evatix-go/core/msgtype"
)

type Variant byte

func (v Variant) MarshalJSON() ([]byte, error) {
	return basicEnumImpl.ToEnumJsonBytes(v.Value()), nil
}

func (v Variant) UnmarshalJSON(data []byte) error {
	panic(msgtype.NotImplemented.ErrorNoRefs("UnmarshalJSON not implemented for bytetype."))
}

func (v Variant) String() string {
	return basicEnumImpl.ToEnumString(v.Value())
}

func (v Variant) StringJson() (jsonString string, err error) {
	return basicEnumImpl.StringJson(v.Value())
}

func (v Variant) StringJsonMust() string {
	return basicEnumImpl.StringJsonMust(v.Value())
}

func (v Variant) StringRangesPtr() *[]string {
	return basicEnumImpl.StringRangesPtr()
}

func (v Variant) StringRanges() []string {
	return basicEnumImpl.StringRanges()
}

func (v Variant) RangesInvalidMessage() string {
	return basicEnumImpl.RangesInvalidMessage()
}

func (v Variant) RangesInvalidErr() error {
	return basicEnumImpl.RangesInvalidErr()
}

func (v Variant) IsValidRange() bool {
	return basicEnumImpl.IsValidRange(v.Value())
}

func (v Variant) IsInvalidRange() bool {
	return !v.IsValidRange()
}

func (v Variant) Value() byte {
	return byte(v)
}

func (v Variant) StringValue() string {
	return string(v)
}

// Add v + n
func (v Variant) Add(n byte) Variant {
	return Variant(v.Value() + n)
}

// Subtract v - n
func (v Variant) Subtract(n byte) Variant {
	return Variant(v.Value() - n)
}

func (v Variant) Is(n Variant) bool {
	return v.Value() == n.Value()
}

// IsBetween val >= start &&  val <= end
func (v Variant) IsBetween(start, end byte) bool {
	val := v.Value()

	return val >= start && val <= end
}

// IsBetweenInt val >= start &&  val <= end
func (v Variant) IsBetweenInt(start, end int) bool {
	val := v.Value()

	return val >= byte(start) && val <= byte(end)
}

func (v Variant) IsEqual(n byte) bool {
	return v.Value() == n
}

// IsGreater v.Value() > n
func (v Variant) IsGreater(n byte) bool {
	return v.Value() > n
}

// IsGreaterEqual v.Value() >= n
func (v Variant) IsGreaterEqual(n byte) bool {
	return v.Value() >= n
}

// IsLess v.Value() < n
func (v Variant) IsLess(n byte) bool {
	return v.Value() < n
}

// IsLessEqual v.Value() <= n
func (v Variant) IsLessEqual(n byte) bool {
	return v.Value() <= n
}

func (v Variant) IsEqualInt(n int) bool {
	return v.Value() == byte(n)
}

// IsGreaterInt v.Value() > n
func (v Variant) IsGreaterInt(n int) bool {
	return v.Value() > byte(n)
}

// IsGreaterEqualInt v.Value() >= n
func (v Variant) IsGreaterEqualInt(n int) bool {
	return v.Value() >= byte(n)
}

// IsLessInt v.Value() < n
func (v Variant) IsLessInt(n int) bool {
	return v.Value() < byte(n)
}

// IsLessEqualInt v.Value() <= n
func (v Variant) IsLessEqualInt(n int) bool {
	return v.Value() <= byte(n)
}

func (v Variant) AsBasicEnumContractsBinder() coreinterface.BasicEnumContractsBinder {
	return v
}
