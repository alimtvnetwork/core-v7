package enumtype

import (
	"encoding/json"
	"errors"
	"strconv"
)

type Variant byte

const (
	Invalid Variant = iota
	Boolean
	Byte
	UnsignedInteger16
	UnsignedInteger32
	UnsignedInteger64
	Integer8
	Integer16
	Integer32
	Integer
	String
)

func (it Variant) Value() byte {
	return byte(it)
}

func (it Variant) IsBoolean() bool {
	return it == Boolean
}

func (it Variant) IsByte() bool {
	return it == Byte
}

func (it Variant) IsUnsignedInteger16() bool {
	return it == UnsignedInteger16
}

func (it Variant) IsUnsignedInteger32() bool {
	return it == UnsignedInteger32
}

func (it Variant) IsUnsignedInteger64() bool {
	return it == UnsignedInteger64
}

func (it Variant) IsInteger8() bool {
	return it == Integer8
}

func (it Variant) IsInteger16() bool {
	return it == Integer16
}

func (it Variant) IsInteger32() bool {
	return it == Integer32
}

func (it Variant) IsInteger() bool {
	return it == Integer
}

func (it Variant) IsString() bool {
	return it == String
}

func (it Variant) Name() string {
	return rangesMap[it]
}

func (it Variant) String() string {
	return rangesMap[it]
}

func (it Variant) NameValue() string {
	return it.Name() + "[" + it.ValueString() + "]"
}

func (it Variant) IsNameEqual(name string) bool {
	return it.Name() == name
}

func (it Variant) IsNameOf(names ...string) bool {
	for _, name := range names {
		if it.IsNameEqual(name) {
			return true
		}
	}

	return false
}

func (it Variant) ToNumberString() string {
	return it.ValueString()
}

func (it Variant) IsValid() bool {
	return it != Invalid
}

func (it Variant) IsInvalid() bool {
	return it == Invalid
}

func (it Variant) ValueByte() byte {
	return byte(it)
}

func (it Variant) ValueInt() int {
	return int(it)
}

func (it Variant) ValueInt8() int8 {
	return int8(it)
}

func (it Variant) ValueInt16() int16 {
	return int16(it)
}

func (it Variant) ValueInt32() int32 {
	return int32(it)
}

func (it Variant) ValueString() string {
	return strconv.Itoa(it.ValueInt())
}

func (it Variant) Format(format string) (compiled string) {
	panic("not supported")
}

func (it Variant) MarshalJSON() ([]byte, error) {
	return json.Marshal(rangesMap[it])
}

func (it *Variant) UnmarshalJSON(data []byte) error {
	var toString string

	if len(data) > 0 {
		toString = string(data)
	}

	if toString == "" || len(toString) <= 2 {
		return errors.New("cannot map to variant or length is below 2 : " + toString)
	}

	unWrapped := toString[1 : len(toString)-1]
	newVariant, hasFound := stringToVariantMap[unWrapped]

	if hasFound {
		*it = Variant(newVariant.ValueByte())

		return nil
	}

	// has error
	return errors.New("not found in map : " + toString)
}
