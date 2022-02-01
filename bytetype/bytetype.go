package bytetype

import (
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coreinterface"
)

type Variant byte

func (it Variant) IsValid() bool {
	return it != 0
}

func (it Variant) IsInvalid() bool {
	return it == 0
}

func (it Variant) NameValue() string {
	return BasicEnumImpl.NameWithValue(it)
}

func (it Variant) ToNumberString() string {
	return BasicEnumImpl.ToNumberString(it.Value())
}

func (it Variant) Name() string {
	return BasicEnumImpl.ToEnumString(it.Value())
}

func (it Variant) UnmarshallToValue(jsonUnmarshallingValue []byte) (byte, error) {
	newEmpty := Variant(0)
	err := corejson.
		Deserialize.
		UsingBytes(
			jsonUnmarshallingValue, &newEmpty)

	if err != nil {
		return 0, err
	}

	return newEmpty.Value(), nil
}

func (it Variant) MarshalJSON() ([]byte, error) {
	return BasicEnumImpl.ToEnumJsonBytes(it.Value()), nil
}

func (it *Variant) UnmarshalJSON(data []byte) error {
	newEmpty := Variant(0)
	err := corejson.
		Deserialize.
		UsingBytes(
			data, &newEmpty)

	if err == nil {
		*it = newEmpty
	}

	return err
}

func (it Variant) String() string {
	return BasicEnumImpl.ToEnumString(it.Value())
}

func (it Variant) JsonString() string {
	return BasicEnumImpl.JsonString(it)
}

func (it Variant) StringRangesPtr() *[]string {
	return BasicEnumImpl.StringRangesPtr()
}

func (it Variant) StringRanges() []string {
	return BasicEnumImpl.StringRanges()
}

func (it Variant) RangesInvalidMessage() string {
	return BasicEnumImpl.RangesInvalidMessage()
}

func (it Variant) RangesInvalidErr() error {
	return BasicEnumImpl.RangesInvalidErr()
}

func (it Variant) IsValidRange() bool {
	return BasicEnumImpl.IsValidRange(it.Value())
}

func (it Variant) IsInvalidRange() bool {
	return !it.IsValidRange()
}

func (it Variant) Value() byte {
	return byte(it)
}

func (it Variant) StringValue() string {
	return string(it)
}

// Add v + n
func (it Variant) Add(n byte) Variant {
	return Variant(it.Value() + n)
}

// Subtract v - n
func (it Variant) Subtract(n byte) Variant {
	return Variant(it.Value() - n)
}

func (it Variant) Is(n Variant) bool {
	return it.Value() == n.Value()
}

// IsBetween val >= start &&  val <= end
func (it Variant) IsBetween(start, end byte) bool {
	val := it.Value()

	return val >= start && val <= end
}

// IsBetweenInt val >= start &&  val <= end
func (it Variant) IsBetweenInt(start, end int) bool {
	val := it.Value()

	return val >= byte(start) && val <= byte(end)
}

func (it Variant) IsEqual(n byte) bool {
	return it.Value() == n
}

// IsGreater v.Value() > n
func (it Variant) IsGreater(n byte) bool {
	return it.Value() > n
}

// IsGreaterEqual v.Value() >= n
func (it Variant) IsGreaterEqual(n byte) bool {
	return it.Value() >= n
}

// IsLess v.Value() < n
func (it Variant) IsLess(n byte) bool {
	return it.Value() < n
}

// IsLessEqual v.Value() <= n
func (it Variant) IsLessEqual(n byte) bool {
	return it.Value() <= n
}

func (it Variant) IsEqualInt(n int) bool {
	return it.Value() == byte(n)
}

// IsGreaterInt v.Value() > n
func (it Variant) IsGreaterInt(n int) bool {
	return it.Value() > byte(n)
}

// IsGreaterEqualInt v.Value() >= n
func (it Variant) IsGreaterEqualInt(n int) bool {
	return it.Value() >= byte(n)
}

// IsLessInt v.Value() < n
func (it Variant) IsLessInt(n int) bool {
	return it.Value() < byte(n)
}

// IsLessEqualInt v.Value() <= n
func (it Variant) IsLessEqualInt(n int) bool {
	return it.Value() <= byte(n)
}

func (it Variant) RangeNamesCsv() string {
	return BasicEnumImpl.RangeNamesCsv()
}

func (it Variant) TypeName() string {
	return BasicEnumImpl.TypeName()
}

func (it Variant) AsBasicEnumContractsBinder() coreinterface.BasicEnumContractsBinder {
	return &it
}
